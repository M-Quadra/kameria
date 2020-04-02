package kameria

import (
	"fmt"
	"sync"
	"testing"
)

func TestOperationQueue(t *testing.T) {
	rw := sync.RWMutex{}
	num := 0
	lim := 10000

	que := OperationQueue{}
	que.MaxConcurrentOperationCount(10)
	fmt.Println("maxConcurrentOperationCount:", que.MaxConcurrentOperationCount())
	for i := 0; i < lim; i++ {
		que.AddOperation(func() {
			rw.Lock()
			defer rw.Unlock()

			num++
		})
	}

	que.Wait()
	fmt.Println("num:", num)
	if num != lim {
		t.Fail()
	}
}
