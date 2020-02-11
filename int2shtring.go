package kameria

import "strconv"

// Int2String int to string
func Int2String(i interface{}, base int) string {
	switch i.(type) {
	case int:
		v := int64(i.(int))
		return strconv.FormatInt(v, base)
	case int8:
		v := int64(i.(int8))
		return strconv.FormatInt(v, base)
	case int16:
		v := int64(i.(int16))
		return strconv.FormatInt(v, base)
	case int32:
		v := int64(i.(int32))
		return strconv.FormatInt(v, base)
	case int64:
		return strconv.FormatInt(i.(int64), base)
	}

	return ""
}
