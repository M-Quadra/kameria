package kameria

import (
	"sync"
	"time"
)

// OperationQueue task queue
//  recommend que := &OperationQueue{}
//  if que.Close(true);que=&OperationQueue{};que.Close()
type OperationQueue struct {
	rw sync.RWMutex
	wg sync.WaitGroup

	funcChan  chan func()
	countChan chan int

	runningCount                int
	maxConcurrentOperationCount int

	wannaClose bool
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
	funcChan := make(chan func())
	countChan := make(chan int)
	slf.funcChan = funcChan
	slf.countChan = countChan

	go func() {
		for cnt := range countChan {
			slf.rw.Lock()
			slf.runningCount -= cnt
			if slf.runningCount >= slf.maxConcurrentOperationCount {
				slf.rw.Unlock()
				continue
			}

			slf.runningCount++
			slf.rw.Unlock()

			fc := <-funcChan
			go func() {
				fc()

				slf.rw.RLock()
				if !slf.wannaClose {
					slf.wg.Done()
				}
				slf.rw.RUnlock()

				countChan <- 1
			}()
		}
	}()
}

// AddOperation add task
func (slf *OperationQueue) AddOperation(fc func()) {
	slf.rw.RLock()
	if slf.wannaClose || slf.maxConcurrentOperationCount <= 0 {
		slf.rw.RUnlock()
		fc()
		return
	}
	slf.rw.RUnlock()

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

// Close close goroutine
//  async control when async is empty
//  default async=false
func (slf *OperationQueue) Close(async ...bool) {
	slf.wg.Wait()

	slf.rw.Lock()
	slf.wannaClose = true
	slf.maxConcurrentOperationCount = 0
	slf.rw.Unlock()

	isAsync := len(async) > 0 && async[0]
	if isAsync {
		go slf.closeGoroutine()
	} else {
		slf.closeGoroutine()
	}
}

func (slf *OperationQueue) closeGoroutine() {
	// when que.Close();que = OperationQueue{};que.Close();
	time.Sleep(time.Duration(100) * time.Millisecond)

	for {
		slf.rw.RLock()
		cnt := slf.runningCount
		slf.rw.RUnlock()
		if cnt <= 0 {
			break
		}

		slf.funcChan <- func() {}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}

	if slf.funcChan != nil {
		close(slf.funcChan)
		slf.funcChan = nil
	}
	if slf.countChan != nil {
		close(slf.countChan)
		slf.countChan = nil
	}
}
