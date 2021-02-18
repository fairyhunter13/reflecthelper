package reflecthelper

import (
	"reflect"
	"time"
)

// TODO: Add FuncOption in here

// ExtractBool extract the underlying bool value from the val of reflect.Value.
func ExtractBool(val reflect.Value) (res bool, err error) {
	res, err = extractBool(val, NewDefaultOption())
	return
}

// ExtractInt gets the underlying int value from val of reflect.Value.
func ExtractInt(val reflect.Value) (result int64, err error) {
	result, err = extractInt(val, NewDefaultOption())
	return
}

// ExtractUint extracts the underlying uint value from val of reflect.Value.
func ExtractUint(val reflect.Value) (result uint64, err error) {
	result, err = extractUint(val, NewDefaultOption())
	return
}

// ExtractFloat extracts the underlying float value from val of reflect.Value.
func ExtractFloat(val reflect.Value) (result float64, err error) {
	result, err = extractFloat(val, NewDefaultOption())
	return
}

// ExtractComplex gets the underlying complex value from val of reflect.Value.
func ExtractComplex(val reflect.Value) (result complex128, err error) {
	result, err = extractComplex(val, NewDefaultOption())
	return
}

// ExtractString gets the underlying string value from val of reflect.Value.
func ExtractString(val reflect.Value) (result string, err error) {
	result, err = extractString(val, NewDefaultOption())
	return
}

// ExtractTime extracts the time from val of reflect.Value.
func ExtractTime(val reflect.Value) (result *time.Time, err error) {
	result, err = extractTime(val, NewDefaultOption())
	return
}

// TryExtract tries to extract the real value from the val of reflect.Value.
func TryExtract(val reflect.Value) (result interface{}, err error) {
	result, err = tryExtract(val, NewDefaultOption())
	return
}
