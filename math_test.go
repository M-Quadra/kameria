package kameria

import (
	"fmt"
	"testing"
)

func TestMathCmp(t *testing.T) {
	fmt.Println("[kameria] Math Cmp Test......")

	fmt.Println("kameria.Min4Int(1, 5)...")
	minInt := Min4Int(1, 5)
	if minInt != 1 {
		t.Fail()
	} else {
		fmt.Println(minInt)
	}

	fmt.Println("kameria.Max4Int(1, 3)...")
	maxInt := Max4Int(1, 3)
	if maxInt != 3 {
		t.Fail()
	} else {
		fmt.Println(maxInt)
	}

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

	fmt.Println("kameria.Min4String(\"AMD\", \"Intel\")...")
	minStr := Min4String("AMD", "Intel")
	if minStr != "AMD" {
		t.Fail()
	} else {
		fmt.Println(minStr)
	}
}
