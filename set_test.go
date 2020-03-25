package kameria

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	fmt.Println("\n[kameria] Unique Test...")

	strIptAry := []string{"app", "google", "m$", "m$", "m$", "m$", "noDieBook"}
	strOptAry := []string{"app", "google", "m$", "noDieBook"}

	intIptAry := []int{9700, 9800, 9900, 7920, 7920, 7920, 7920}
	intOptAry := []int{9700, 9800, 9900, 7920}

	int64IptAry := []int64{3700, 3800, 3900, 3950, 3950, 3950, 3950}
	int64OptAry := []int64{3700, 3800, 3900, 3950}

	for i, v := range Unique4String(strIptAry) {
		if strOptAry[i] != v {
			t.Fail()
		}
	}

	for i, v := range Unique4Int(intIptAry) {
		if intOptAry[i] != v {
			t.Fail()
		}
	}

	for i, v := range Unique4int64(int64IptAry) {
		if int64OptAry[i] != v {
			t.Fail()
		}
	}
}
