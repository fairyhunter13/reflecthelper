package reflecthelper

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	// DefaultFloatPrecision specifies the default precision used in this package.
	// This is the default maximum precision.
	DefaultFloatPrecision = -1
	// DefaultBitSize is the default bit size used for the conversion in this package.
	DefaultBitSize = 64
	// DefaultBaseSystem is the default base system used for decimal in this package.
	DefaultBaseSystem = 10
)

func getBool(val reflect.Value) (result bool, err error) {
	err = checkValid(val)
	if err != nil {
		return
	}
	switch GetKind(val) {
	case reflect.Bool:
		result = val.Bool()
	case reflect.Ptr:
		result, err = getBool(reflect.Indirect(val))
	default:
		str := getDefaultString(val)
		result, err = strconv.ParseBool(str)
	}
	return
}

func getInt(val reflect.Value) (result int64, err error) {
	err = checkValid(val)
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
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uintStr := strconv.FormatUint(val.Uint(), DefaultBaseSystem)
		result, err = strconv.ParseInt(uintStr, DefaultBaseSystem, DefaultBitSize)
	}
	return
}

func getString(val reflect.Value) (result string, err error) {
	err = checkValid(val)
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
		result, err = getString(reflect.Indirect(val))
	default:
		result = getDefaultString(val)
	}
	return
}

func getDefaultFloatStr(floatVal float64) (result string) {
	result = strconv.FormatFloat(floatVal, 'f', DefaultFloatPrecision, DefaultBitSize)
	return
}

func getDefaultString(val reflect.Value) (result string) {
	result = fmt.Sprintf("%v", val.Interface())
	return
}

func getErrOverflow(val reflect.Value) (err error) {
	err = fmt.Errorf(
		"Assigner encounters overflow %s, underlying val: %s",
		GetKind(val),
		val.String(),
	)
	return
}
