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

	ptr, file, line, ok := runtime.Caller(1)
	if ok {
		nowTime := time.Now()
		fmt.Println("error happen:")
		fmt.Println("    ", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix(), nowTime.UnixNano())
		fmt.Printf("     %s:%d +%#x\n", file, line, ptr)
		fmt.Println("    ", err)
	}
	return true
}
