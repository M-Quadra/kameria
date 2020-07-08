package kameria

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestOperationQueue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	rw := sync.RWMutex{}
	num := 0
	lim := 10000

	que := OperationQueue{}
	que.MaxConcurrentOperationCount(900)
	for i := 0; i < lim; i++ {
		que.AddOperation(func() {
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Nanosecond)
			rw.Lock()
			defer rw.Unlock()

			num++
		})

		if rand.Intn(2) == 0 {
			go que.MaxConcurrentOperationCount(rand.Intn(900))
		}
		if rand.Intn(2) == 0 {
			go que.Wait()
		}
	}
	que.Wait()

	if num != lim {
		t.Fail()
	}
}

func TestOperationQueueLoop(t *testing.T) {
	stGoroutineNum := runtime.NumGoroutine()
	for i := 0; i < 10; i++ {
		TestOperationQueue(t)
	}

	time.Sleep(time.Second)
	if runtime.NumGoroutine() != stGoroutineNum {
		t.Fail()
	}
}
