package kameria

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	fmt.Println("[kameria] Queue4String Test...")

	iptAry := []string{"AMD", "YES"}
	optAry := []string{}

	que := Queue4String{}
	for _, v := range iptAry {
		que.Push(v)
	}

	for !que.Empty() {
		v, ok := que.Pop()
		if !ok {
			t.Fail()
		}

		optAry = append(optAry, v)
		fmt.Println(v)
	}

	for i := range optAry {
		if iptAry[i] != optAry[i] {
			t.Fail()
		}
	}
}
