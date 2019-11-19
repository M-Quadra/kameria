package kameria

import (
	"fmt"
	"runtime"
	"time"
)

// HasError error printer
func HasError(err error) bool {
	if err == nil {
		return false
	}

	nowTime := time.Now()
	fmt.Println("error happen:")
	fmt.Println("    ", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix(), nowTime.UnixNano())
	fmt.Println("    ", err)
	for i := 1; i < 5; i++ {
		ptr, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fmt.Printf("     %s:%d +%#x\n", file, line, ptr)
	}
	return true
}
