package kameria

import (
	"sync"
	"testing"
	"time"
)

func TestOperationQueue(t *testing.T) {
	rw := sync.RWMutex{}
	num := 0
	lim := 10000

	que := OperationQueue{}
	que.MaxConcurrentOperationCount(900)
	for i := 0; i < lim; i++ {
		que.AddOperation(func() {
			rw.Lock()
			defer rw.Unlock()

			num++
		})
	}
	que.MaxConcurrentOperationCount(100)
	que.Wait()
	que.Close()

	que = OperationQueue{}
	que.MaxConcurrentOperationCount(200)
	for i := 0; i < lim; i++ {
		que.AddOperation(func() {
			rw.Lock()
			defer rw.Unlock()

			num++
		})
	}
	que.Wait()
	que.Close()

	que2 := &OperationQueue{}
	que2.MaxConcurrentOperationCount(100)
	for i := 0; i < lim; i++ {
		que2.AddOperation(func() {
			rw.Lock()
			defer rw.Unlock()

			num++
		})
	}
	que2.MaxConcurrentOperationCount(10)
	que2.Wait()
	que2.Close(true)

	que2 = &OperationQueue{}
	for i := 0; i < lim; i++ {
		que2.AddOperation(func() {
			rw.Lock()
			defer rw.Unlock()

			num++
		})
	}
	que2.Wait()
	que2.Close()

	time.Sleep(time.Duration(1) * time.Second)
	if num != lim*4 {
		t.Fail()
	}
}
