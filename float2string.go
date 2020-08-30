package kameria

import "strconv"

// Float2String float to 'f' string
func Float2String(i interface{}) string {
	switch i.(type) {
	case float32:
		v := float64(i.(float32))
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float64:
		v := i.(float64)
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	return ""
}
