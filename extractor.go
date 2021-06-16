package reflecthelper

import (
	"reflect"
	"strconv"
	"time"
)

func extractBool(val reflect.Value, option *Option) (result bool, err error) {
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	originKind := GetKind(val)
	tempVal := GetElem(val)
	for IsValueElemable(val) && tempVal != val {
		result, err = castBool(val.Interface())
		if err == nil {
			return
		}
		val = tempVal
		tempVal = GetElem(val)
	}
	err = nil

	switch GetKind(val) {
	case reflect.Bool:
		result = val.Bool()
	default:
		if val.CanAddr() && !IsKindPtr(originKind) {
			result, err = castBool(val.Addr().Interface())
			if err == nil {
				return
			}
		}
		result, err = castBool(val.Interface())
		if err == nil {
			return
		}
		str := getDefaultString(val)
		result, err = strconv.ParseBool(str)
	}
	return
}

func extractInt(val reflect.Value, option *Option) (result int64, err error) {
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	originKind := GetKind(val)
	tempVal := GetElem(val)
	for IsValueElemable(val) && tempVal != val {
		result, err = castInt64(val.Interface())
		if err == nil {
			return
		}
		val = tempVal
		tempVal = GetElem(val)
	}
	err = nil

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
	default:
		if val.CanAddr() && !IsKindPtr(originKind) {
			result, err = castInt64(val.Addr().Interface())
			if err == nil {
				return
			}
		}
		result, err = castInt64(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = extractString(val, option)
		if err != nil {
			return
		}
		result, err = strconv.ParseInt(str, option.BaseSystem, option.BitSize)
	}
	return
}

func extractUint(val reflect.Value, option *Option) (result uint64, err error) {
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	originKind := GetKind(val)
	tempVal := GetElem(val)
	for IsValueElemable(val) && tempVal != val {
		result, err = castUint64(val.Interface())
		if err == nil {
			return
		}
		val = tempVal
		tempVal = GetElem(val)
	}
	err = nil

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
	default:
		if val.CanAddr() && !IsKindPtr(originKind) {
			result, err = castUint64(val.Addr().Interface())
			if err == nil {
				return
			}
		}
		result, err = castUint64(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = extractString(val, option)
		if err != nil {
			return
		}
		result, err = strconv.ParseUint(str, option.BaseSystem, option.BitSize)
	}
	return
}

func extractFloat(val reflect.Value, option *Option) (result float64, err error) {
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	originKind := GetKind(val)
	tempVal := GetElem(val)
	for IsValueElemable(val) && tempVal != val {
		result, err = castFloat64(val.Interface())
		if err == nil {
			return
		}
		val = tempVal
		tempVal = GetElem(val)
	}
	err = nil

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
	default:
		if val.CanAddr() && !IsKindPtr(originKind) {
			result, err = castFloat64(val.Addr().Interface())
			if err == nil {
				return
			}
		}
		result, err = castFloat64(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = extractString(val, option)
		if err != nil {
			return
		}
		result, err = strconv.ParseFloat(str, option.BitSize)
	}
	return
}

func extractComplex(val reflect.Value, option *Option) (result complex128, err error) {
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	originKind := GetKind(val)
	tempVal := GetElem(val)
	for IsValueElemable(val) && tempVal != val {
		result, err = castComplex128(val.Interface())
		if err == nil {
			return
		}
		val = tempVal
		tempVal = GetElem(val)
	}
	err = nil

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
		if val.CanAddr() && !IsKindPtr(originKind) {
			result, err = castComplex128(val.Addr().Interface())
			if err == nil {
				return
			}
		}
		result, err = castComplex128(val.Interface())
		if err == nil {
			return
		}
		var str string
		str, err = extractString(val, option)
		if err != nil {
			return
		}
		result, err = strconv.ParseComplex(str, option.ComplexBitSize)
	}
	return
}

func extractString(val reflect.Value, option *Option) (result string, err error) {
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	originKind := GetKind(val)
	tempVal := GetElem(val)
	for IsValueElemable(val) && tempVal != val {
		result, err = castString(val.Interface())
		if err == nil {
			return
		}
		val = tempVal
		tempVal = GetElem(val)
	}
	err = nil

	switch GetKind(val) {
	case reflect.Bool:
		boolVal := val.Bool()
		result = strconv.FormatBool(boolVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal := val.Int()
		result = strconv.FormatInt(intVal, option.BaseSystem)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uintVal := val.Uint()
		result = strconv.FormatUint(uintVal, option.BaseSystem)
	case reflect.Float32, reflect.Float64:
		result = getDefaultFloatStr(val.Float(), option)
	case reflect.String:
		result = val.String()
	default:
		if val.CanAddr() && !IsKindPtr(originKind) {
			result, err = castString(val.Addr().Interface())
			if err == nil {
				return
			}
		}
		result, err = castString(val.Interface())
	}
	return
}

func extractTime(val reflect.Value, option *Option) (result time.Time, err error) {
	val = GetChildElem(val)
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	var timeStr string
	switch GetKind(val) {
	case reflect.String:
		timeStr = val.String()
	case reflect.Struct:
		if GetType(val) == TypeTime {
			timeVal := val.Interface().(time.Time)
			result = timeVal
			return
		}

		fallthrough
	default:
		timeStr, err = extractString(val, option)
	}
	if err != nil {
		return
	}

	result, err = parseTime(timeStr, option)
	return
}

func extractDuration(val reflect.Value, option *Option) (res time.Duration, err error) {
	val = GetChildElem(val)
	err = checkExtractValid(val, option)
	if err != nil {
		return
	}

	switch GetKind(val) {
	case reflect.String:
		res, err = time.ParseDuration(val.String())
	default:
		if GetType(val) == TypeDuration {
			durVal := val.Interface().(time.Duration)
			res = durVal
			return
		}

		var resInt64 int64
		resInt64, err = extractInt(val, option)
		if err != nil {
			return
		}

		res = time.Duration(resInt64)
	}
	return
}

func tryExtract(val reflect.Value, opt *Option) (result interface{}, err error) {
	val = GetChildElem(val)
	err = checkExtractValid(val, opt)
	if err != nil {
		return
	}

	switch GetKind(val) {
	case reflect.Bool:
		result, err = extractBool(val, opt)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if IsTypeValueDuration(val) {
			result, err = extractDuration(val, opt)
			return
		}

		result, err = extractInt(val, opt)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result, err = extractUint(val, opt)
	case reflect.Float32, reflect.Float64:
		result, err = extractFloat(val, opt)
	case reflect.Complex64, reflect.Complex128:
		result, err = extractComplex(val, opt)
	case reflect.String:
		result, err = extractString(val, opt)
	case reflect.Struct, reflect.Ptr:
		if IsTypeValueTime(val) {
			result, err = extractTime(val, opt)
			return
		}

		fallthrough
	default:
		err = getErrUnimplementedExtract(val)
	}
	return
}
