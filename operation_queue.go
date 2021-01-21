package kameria

import (
	"sync"
)

// OperationQueue task queue
type OperationQueue struct {
	rw sync.RWMutex
	wg sync.WaitGroup

	infoChan        chan struct{}
	maxGoroutineNum int
}

// Wait blocks until the WaitGroup counter is zero.
func (slf *OperationQueue) Wait() {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.wg.Wait()
}

// MaxConcurrentOperationCount cnt >= 1
//  Get when cnt is empty
//  Set when cnt has value
func (slf *OperationQueue) MaxConcurrentOperationCount(cnt ...int) int {
	if len(cnt) <= 0 {
		slf.rw.RLock()
		val := slf.maxGoroutineNum
		slf.rw.RUnlock()
		return Max4Int(0, val)
	}

	newVal := Max4Int(0, cnt[0])

	slf.rw.Lock()
	defer slf.rw.Unlock()

	if newVal == slf.maxGoroutineNum {
		return newVal
	}

	slf.maxGoroutineNum = newVal
	slf.wg.Wait()

	if newVal <= 0 {
		slf.infoChan = nil
		return newVal
	}

	slf.infoChan = make(chan struct{}, newVal)
	return newVal
}

// Add add operation
func (slf *OperationQueue) Add(operation func()) {
	slf.rw.RLock()
	defer slf.rw.RUnlock()

	cnt := slf.maxGoroutineNum
	if cnt <= 0 {
		operation()
		return
	}

	slf.wg.Add(1)
	slf.infoChan <- struct{}{}
	go func() {
		defer slf.wg.Done()

		operation()
		<-slf.infoChan
	}()
}
