package kameria

import (
	"sync"
)

// OperationQueue task queue
type OperationQueue struct {
	rw sync.RWMutex
	wg sync.WaitGroup

	funcChan  chan func()
	countChan chan int

	runningCount                int
	maxConcurrentOperationCount int
}

// Wait blocks until the WaitGroup counter is zero.
func (slf *OperationQueue) Wait() {
	slf.wg.Wait()
}

// MaxConcurrentOperationCount cnt >= 1
//  Get when cnt is empty
//  Set when cnt has value
func (slf *OperationQueue) MaxConcurrentOperationCount(cnt ...int) int {
	if len(cnt) <= 0 {
		slf.rw.RLock()
		val := slf.maxConcurrentOperationCount
		slf.rw.RUnlock()
		if val <= 0 {
			val = 1
		}
		return val
	}

	newVal := cnt[0]
	if newVal <= 0 {
		return newVal
	}

	slf.rw.Lock()
	slf.maxConcurrentOperationCount = newVal
	slf.rw.Unlock()
	return newVal
}

func (slf *OperationQueue) init() {
	slf.funcChan = make(chan func())
	slf.countChan = make(chan int)
	go func() {
		for cnt := range slf.countChan {
			slf.rw.Lock()
			slf.runningCount -= cnt
			if slf.runningCount >= slf.maxConcurrentOperationCount {
				slf.rw.Unlock()
				continue
			}

			slf.runningCount++
			slf.rw.Unlock()

			fc := <-slf.funcChan
			go func() {
				fc()
				slf.wg.Done()
				slf.countChan <- 1
			}()
		}
	}()
}

// AddOperation add task
func (slf *OperationQueue) AddOperation(fc func()) {
	slf.wg.Add(1)

	slf.rw.Lock()
	if slf.funcChan == nil {
		slf.init()
	}

	if slf.runningCount >= slf.maxConcurrentOperationCount {
		slf.rw.Unlock()
		slf.funcChan <- fc
		return
	}

	slf.runningCount++
	slf.rw.Unlock()

	go func() {
		fc()
		slf.wg.Done()
		slf.countChan <- 1
	}()
}
