package kameria

import (
	"sync"
	"sync/atomic"
)

// OnceMutex sync.Once + Reset
type OnceMutex struct {
	done uint32
	m    sync.Mutex
}

// Do once.Do
func (slf *OnceMutex) Do(f func()) {
	if atomic.LoadUint32(&slf.done) == 0 {
		slf.doSlow(f)
	}
}

func (slf *OnceMutex) doSlow(f func()) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.done == 0 {
		defer atomic.StoreUint32(&slf.done, 1)
		f()
	}
}

// Reset reset once
func (slf *OnceMutex) Reset(sync ...bool) {
	if len(sync) > 0 && sync[0] == true {
		slf.m.Lock()
		defer slf.m.Unlock()
		atomic.StoreUint32(&slf.done, 0)
		return
	}

	go func() { // avoid deadlock
		slf.m.Lock()
		defer slf.m.Unlock()
		atomic.StoreUint32(&slf.done, 0)
	}()
}
