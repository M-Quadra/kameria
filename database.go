package kameria

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// StructTagMap map[tag key] value address
func StructTagMap(tag string, model interface{}) (map[string]interface{}, error) {
	rv := reflect.ValueOf(model)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("need pointer")
	}

	addrStr := fmt.Sprintf("%p", model)
	addrStr = strings.Replace(addrStr, "0x", "", 1)
	ptrUint64, err := strconv.ParseUint(addrStr, 16, 64)
	if err != nil {
		return nil, err
	}

	firstAddr := uintptr(ptrUint64)
	elem := rv.Elem()

	tagMap := map[string]interface{}{}

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		key, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}

		offset := field.Offset
		pointer := unsafe.Pointer(firstAddr + offset)

		// "tagMap[key] = interface{}(pointer)" will return type error when sql.Scan
		// need exact type
		value := elem.FieldByName(field.Name).Interface()
		switch value.(type) {
		case string:
			tagMap[key] = (*string)(pointer)

		case int:
			tagMap[key] = (*int)(pointer)
		case int8:
			tagMap[key] = (*int8)(pointer)
		case int16:
			tagMap[key] = (*int16)(pointer)
		case int32:
			tagMap[key] = (*int32)(pointer)
		case int64:
			tagMap[key] = (*int64)(pointer)

		case float32:
			tagMap[key] = (*float32)(pointer)
		case float64:
			tagMap[key] = (*float64)(pointer)
		default:
			return nil, fmt.Errorf(
				"unsupport %s type name: %s kind: %s",
				key,
				field.Type.Name(),
				field.Type.Kind().String(),
			)
		}

		// set value
		// elem.FieldByName(field.Name).Set(reflect.ValueOf(key))
	}

	return tagMap, nil
}
