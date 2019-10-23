package kameria

import (
	"errors"
	"strconv"
	"strings"
)

//IPv4StringToInt "127.0.0.1" -> 2130706433
func IPv4StringToInt(ipStr string) (int64, error) {
	strAry := strings.Split(ipStr, ".")
	if len(strAry) != 4 {
		return 0, errors.New("no IPv4")
	}

	opt := int64(0)
	for _, str := range strAry {
		tmp, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, err
		}

		opt <<= 8
		opt += tmp
	}

	return opt, nil
}

//IPv4IntToString 2130706433 -> "127.0.0.1"
func IPv4IntToString(ipInt int64) (string, error) {
	opt := ""
	for i := 0; i < 4; i++ {
		tmpInt := ipInt & 0xFF
		if tmpInt > 0xFF {
			return "", errors.New("no IPv4")
		}

		ipInt >>= 8
		opt = strconv.FormatInt(tmpInt, 10) + opt
		if i+1 < 4 {
			opt = "." + opt
		}
	}

	return opt, nil
}
