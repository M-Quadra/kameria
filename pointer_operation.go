package kameria

import (
	"fmt"
	"strconv"
	"strings"
)

func uintptrAddress(v interface{}) (uintptr, error) {
	addrStr := fmt.Sprintf("%p", v)
	addrStr = strings.Replace(addrStr, "0x", "", 1)
	ptrUint64, err := strconv.ParseUint(addrStr, 16, 64)
	if err != nil {
		return 0, err
	}

	return uintptr(ptrUint64), nil
}
