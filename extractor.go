package reflecthelper

import (
	"reflect"
	"strconv"

	"github.com/reiver/go-cast"
)

func checkExtractValid(val reflect.Value) (err error) {
	err = getErrIsValid(val)
	if err != nil {
		return
	}
	err = getErrCanInterface(val)
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
	case reflect.Ptr, reflect.Interface:
		result, err = ExtractBool(val.Elem())
	default:
		result, err = cast.Bool(val.Interface())
		if err == nil {
			return
		}
		str := getDefaultString(val)
		result, err = strconv.ParseBool(str)
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
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		result = int64(val.Uint())
	case reflect.Ptr, reflect.Interface:
		result, err = ExtractInt(val.Elem())
	default:
		result, err = cast.Int64(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = ExtractString(val)
		if err != nil {
			return
		}
		result, err = strconv.ParseInt(str, DefaultBaseSystem, DefaultBitSize)
	}
	return
}

// ExtractUint extracts the underlying uint value from val of reflect.Value.
func ExtractUint(val reflect.Value) (result uint64, err error) {
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
		valInt := val.Int()
		if valInt < 0 {
			err = getErrOverflow(val)
			return
		}
		result = uint64(valInt)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = val.Uint()
	case reflect.Ptr, reflect.Interface:
		result, err = ExtractUint(val.Elem())
	default:
		result, err = cast.Uint64(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = ExtractString(val)
		if err != nil {
			return
		}
		result, err = strconv.ParseUint(str, DefaultBaseSystem, DefaultBitSize)
	}
	return
}

// ExtractFloat extracts the underlying float value from val of reflect.Value.
func ExtractFloat(val reflect.Value) (result float64, err error) {
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
	case reflect.Int8, reflect.Int16, reflect.Int32:
		result = float64(val.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		result = float64(val.Uint())
	case reflect.Float32, reflect.Float64:
		result = val.Float()
	case reflect.Ptr, reflect.Interface:
		result, err = ExtractFloat(val.Elem())
	default:
		result, err = cast.Float64(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = ExtractString(val)
		if err != nil {
			return
		}
		result, err = strconv.ParseFloat(str, DefaultBitSize)
	}
	return
}

// ExtractComplex gets the underlying complex value from val of reflect.Value.
func ExtractComplex(val reflect.Value) (result complex128, err error) {
	err = checkExtractValid(val)
	if err != nil {
		return
	}
	switch GetKind(val) {
	case reflect.Int8, reflect.Int16:
		result = complex(float64(val.Int()), 0)
	case reflect.Uint8, reflect.Uint16:
		result = complex(float64(val.Uint()), 0)
	case reflect.Float32, reflect.Float64:
		result = complex(float64(val.Float()), 0)
	case reflect.Complex64, reflect.Complex128:
		result = val.Complex()
	default:
		result, err = cast.Complex128(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = ExtractString(val)
		if err != nil {
			return
		}
		result, err = strconv.ParseComplex(str, DefaultComplexBitSize)
	}
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
	case reflect.Ptr, reflect.Interface:
		result, err = ExtractString(val.Elem())
	case reflect.String:
		result = val.String()
	default:
		result, err = cast.String(val.Interface())
	}
	return
}

// TryExtract tries to extract the real value from the val of reflect.Value.
func TryExtract(val reflect.Value) (result interface{}, err error) {
	switch GetKind(val) {
	case reflect.Bool:
		result, err = ExtractBool(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result, err = ExtractInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result, err = ExtractUint(val)
	case reflect.Float32, reflect.Float64:
		result, err = ExtractFloat(val)
	case reflect.Complex64, reflect.Complex128:
		result, err = ExtractComplex(val)
	case reflect.String:
		result, err = ExtractString(val)
	default:
		err = getErrUnimplementedExtract(val)
	}
	return
}
