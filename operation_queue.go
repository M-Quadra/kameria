package kameria

import (
	"sync"
)

// OperationQueue task queue
type OperationQueue struct {
	rw sync.RWMutex
	wg sync.WaitGroup

	runningCount                int
	maxConcurrentOperationCount int
	taskAry                     []func()
}

// Wait blocks until the WaitGroup counter is zero.
func (slf *OperationQueue) Wait() {
	slf.wg.Wait()
}

// MaxConcurrentOperationCount cnt >= 1
// get when cnt is empty
// set when cnt has value
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

// AddOperation add task
func (slf *OperationQueue) AddOperation(fc func()) {
	slf.rw.RLock()
	nowNum := slf.runningCount
	limNum := slf.maxConcurrentOperationCount
	slf.rw.RUnlock()

	if nowNum < limNum {
		slf.rw.Lock()
		slf.runningCount++
		slf.rw.Unlock()
		slf.wg.Add(1)
		go func() {
			defer slf.wg.Done()

			fc()
			slf.didFinishOne()
		}()
		return
	}

	slf.rw.Lock()
	slf.taskAry = append(slf.taskAry, fc)
	slf.rw.Unlock()
}

func (slf *OperationQueue) didFinishOne() {
	slf.rw.Lock()

	slf.runningCount--

	nowNum := slf.runningCount
	limNum := slf.maxConcurrentOperationCount
	lstNum := len(slf.taskAry)
	if lstNum <= 0 || nowNum >= limNum {
		slf.rw.Unlock()
		return
	}

	fc := slf.taskAry[0]
	if fc == nil {
		slf.rw.Unlock()
		return
	}

	slf.taskAry = slf.taskAry[1:]
	slf.rw.Unlock()
	slf.AddOperation(fc)
}
