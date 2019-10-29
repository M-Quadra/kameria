package kameria

import "fmt"

var lastErrStr string

// HasError error printer
func HasError(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	if errStr != lastErrStr {
		fmt.Println(errStr)
		lastErrStr = errStr
	}
	return true
}
