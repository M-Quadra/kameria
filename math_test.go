package kameria

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/M-Quadra/kazaana"
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

func TestSoftmaxFloat64(t *testing.T) {

	ary0 := SoftmaxFloat64([]float64{1, 2, 3, 4, 5})
	ary1 := []float64{
		0.011656230956039607,
		0.03168492079612427,
		0.0861285444362687,
		0.23412165725273662,
		0.6364086465588308,
	}

	data0, err := json.Marshal(ary0)
	if kazaana.HasError(err) {
		t.Fail()
	}

	data1, err := json.Marshal(ary1)
	if kazaana.HasError(err) {
		t.Fail()
	}

	if string(data0) != string(data1) {
		t.Fail()
	}
}

func TestAvgFloat64(t *testing.T) {
	avg := AvgFloat64([]float64{1, 2, 3, 4, 5})
	if avg != 3 {
		t.Fail()
	}
}
