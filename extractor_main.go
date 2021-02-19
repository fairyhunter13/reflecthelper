package reflecthelper

import (
	"reflect"
	"time"
)

// ExtractBool extract the underlying bool value from the val of reflect.Value.
func ExtractBool(val reflect.Value, fnOpts ...FuncOption) (res bool, err error) {
	opt := NewOption().Assign(fnOpts...)
	res, err = extractBool(val, opt)
	return
}

// ExtractInt gets the underlying int value from val of reflect.Value.
func ExtractInt(val reflect.Value, fnOpts ...FuncOption) (result int64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractInt(val, opt)
	return
}

// ExtractUint extracts the underlying uint value from val of reflect.Value.
func ExtractUint(val reflect.Value, fnOpts ...FuncOption) (result uint64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractUint(val, opt)
	return
}

// ExtractFloat extracts the underlying float value from val of reflect.Value.
func ExtractFloat(val reflect.Value, fnOpts ...FuncOption) (result float64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractFloat(val, opt)
	return
}

// ExtractComplex gets the underlying complex value from val of reflect.Value.
func ExtractComplex(val reflect.Value, fnOpts ...FuncOption) (result complex128, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractComplex(val, opt)
	return
}

// ExtractString gets the underlying string value from val of reflect.Value.
func ExtractString(val reflect.Value, fnOpts ...FuncOption) (result string, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractString(val, opt)
	return
}

// ExtractTime extracts the time from val of reflect.Value.
func ExtractTime(val reflect.Value, fnOpts ...FuncOption) (result *time.Time, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractTime(val, opt)
	return
}

// TryExtract tries to extract the real value from the val of reflect.Value.
func TryExtract(val reflect.Value, fnOpts ...FuncOption) (result interface{}, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = tryExtract(val, opt)
	return
}
