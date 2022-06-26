package kameria

import (
	"fmt"
	"testing"
)

func TestMathCmp(t *testing.T) {
	fmt.Println("[kameria] Math Cmp Test......")

	fmt.Println("kameria.Limit4Int(1, 12, 6)...")
	limitInt := Limit4Int(1, 12, 6)
	if limitInt != 6 {
		t.Fail()
	} else {
		fmt.Println(limitInt)
	}

	fmt.Println("kameria.Limit4Int(1, -12, 6)...")
	limitInt = Limit4Int(1, -12, 6)
	if limitInt != 1 {
		t.Fail()
	} else {
		fmt.Println(limitInt)
	}
}
