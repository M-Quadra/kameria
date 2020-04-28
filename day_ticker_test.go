package kameria

import (
	"fmt"
	"testing"
	"time"
)

func TestDayTicker(t *testing.T) {
	nowTime := time.Now()
	dTicker := NewDayTicker(nowTime.Hour(), nowTime.Minute(), nowTime.Second())
	nextTime := dTicker.Next()
	fmt.Println(nextTime)
}
