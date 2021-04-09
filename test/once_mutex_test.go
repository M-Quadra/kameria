package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/M-Quadra/kameria"
)

func TestOnceMutex(t *testing.T) {
	once := kameria.OnceMutex{}
	num := 0

	once.Do(func() {
		num++
	})
	once.Do(func() {
		num++
	})
	if num != 1 {
		t.Fail()
		return
	}

	once.Reset(true)
	once.Do(func() {
		num++
	})
	if num != 2 {
		t.Fail()
		return
	}

	once.Reset(true)
	once.Do(func() {
		defer once.Reset()
		num++
	})
	if num != 3 {
		t.Fail()
		return
	}

	l := 10000
	wg := sync.WaitGroup{}
	wg.Add(l)
	once.Reset(true)
	for i := 0; i < l; i++ {
		if rand.Intn(2) == 0 {
			go func() {
				defer wg.Done()

				time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
				once.Do(func() {
					defer once.Reset()
					num++
				})
			}()
		} else {
			go func() {
				defer wg.Done()

				time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
				once.Reset(true)
			}()
		}
	}
	wg.Wait()

	if num == 3 {
		t.Fail()
	}
}
