package reflecthelper

import (
	"reflect"
	"strconv"
)

func checkExtractValid(val reflect.Value) (err error) {
	err = checkValid(val)
	if err != nil {
		return
	}
	err = getErrCanInterface(val)
	return
}

// ExtractString gets the underlying string value from val of reflect.Value.
func ExtractString(val reflect.Value) (result string, err error) {
	err = checkExtractValid(val)
	if err != nil {
		return
	}
	switch GetKind(val) {
	case reflect.Bool:
		boolVal := val.Bool()
		result = strconv.FormatBool(boolVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal := val.Int()
		result = strconv.FormatInt(intVal, DefaultBaseSystem)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uintVal := val.Uint()
		result = strconv.FormatUint(uintVal, DefaultBaseSystem)
	case reflect.Float32, reflect.Float64:
		result = getDefaultFloatStr(val.Float())
	case reflect.Ptr:
		result, err = ExtractString(reflect.Indirect(val))
	case reflect.String:
		result = val.String()
	default:
		result = getDefaultString(val)
	}
	return
}

// ExtractInt gets the underlying int value from val of reflect.Value.
func ExtractInt(val reflect.Value) (result int64, err error) {
	err = checkExtractValid(val)
	if err != nil {
		return
	}
	switch GetKind(val) {
	case reflect.Bool:
		if val.Bool() {
			result = 1
		} else {
			result = 0
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = val.Int()
	case reflect.Ptr:
		result, err = ExtractInt(reflect.Indirect(val))
	default:
		var str string
		str, err = ExtractString(val)
		if err != nil {
			return
		}
		result, err = strconv.ParseInt(str, DefaultBaseSystem, DefaultBitSize)
	}
	return
}

// ExtractBool extract the underlying bool value from the val of reflect.Value.
func ExtractBool(val reflect.Value) (result bool, err error) {
	err = checkExtractValid(val)
	if err != nil {
		return
	}
	switch GetKind(val) {
	case reflect.Bool:
		result = val.Bool()
	case reflect.Ptr:
		result, err = ExtractBool(reflect.Indirect(val))
	default:
		str := getDefaultString(val)
		result, err = strconv.ParseBool(str)
	}
	return
}
