package reflecthelper

import (
	"reflect"
	"time"
)

// GetBool is ExtractBool without error.
func GetBool(val reflect.Value, fnOpts ...FuncOption) (res bool) {
	res, _ = ExtractBool(val, fnOpts...)
	return
}

// ExtractBool extract the underlying bool value from the val of reflect.Value.
func ExtractBool(val reflect.Value, fnOpts ...FuncOption) (res bool, err error) {
	opt := NewOption().Assign(fnOpts...)
	res, err = extractBool(val, opt)
	return
}

// GetInt is ExtractInt without error.
func GetInt(val reflect.Value, fnOpts ...FuncOption) (result int64) {
	result, _ = ExtractInt(val, fnOpts...)
	return
}

// ExtractInt gets the underlying int value from val of reflect.Value.
func ExtractInt(val reflect.Value, fnOpts ...FuncOption) (result int64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractInt(val, opt)
	return
}

// GetUint is ExtractUint without error.
func GetUint(val reflect.Value, fnOpts ...FuncOption) (result uint64) {
	result, _ = ExtractUint(val, fnOpts...)
	return
}

// ExtractUint extracts the underlying uint value from val of reflect.Value.
func ExtractUint(val reflect.Value, fnOpts ...FuncOption) (result uint64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractUint(val, opt)
	return
}

// GetFloat is ExtractFloat without error.
func GetFloat(val reflect.Value, fnOpts ...FuncOption) (result float64) {
	result, _ = ExtractFloat(val, fnOpts...)
	return
}

// ExtractFloat extracts the underlying float value from val of reflect.Value.
func ExtractFloat(val reflect.Value, fnOpts ...FuncOption) (result float64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractFloat(val, opt)
	return
}

// GetComplex is ExtractComplex without error.
func GetComplex(val reflect.Value, fnOpts ...FuncOption) (result complex128) {
	result, _ = ExtractComplex(val, fnOpts...)
	return
}

// ExtractComplex gets the underlying complex value from val of reflect.Value.
func ExtractComplex(val reflect.Value, fnOpts ...FuncOption) (result complex128, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractComplex(val, opt)
	return
}

// GetString is ExtractString without error.
func GetString(val reflect.Value, fnOpts ...FuncOption) (result string) {
	result, _ = ExtractString(val, fnOpts...)
	return
}

// ExtractString gets the underlying string value from val of reflect.Value.
func ExtractString(val reflect.Value, fnOpts ...FuncOption) (result string, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractString(val, opt)
	return
}

// GetTime is ExtractTime without error.
func GetTime(val reflect.Value, fnOpts ...FuncOption) (result time.Time) {
	result, _ = ExtractTime(val, fnOpts...)
	return
}

// ExtractTime extracts the time from val of reflect.Value.
func ExtractTime(val reflect.Value, fnOpts ...FuncOption) (result time.Time, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractTime(val, opt)
	return
}

// TryGet is TryExtract without error.
func TryGet(val reflect.Value, fnOpts ...FuncOption) (result interface{}) {
	result, _ = TryExtract(val, fnOpts...)
	return
}

// TryExtract tries to extract the real value from the val of reflect.Value.
func TryExtract(val reflect.Value, fnOpts ...FuncOption) (result interface{}, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = tryExtract(val, opt)
	return
}
