package reflecthelper

import (
	"net/url"
	"reflect"
	"time"
)

func getValFromInterface(input interface{}) (val reflect.Value) {
	switch input := input.(type) {
	case reflect.Value:
		val = input
	default:
		val = reflect.ValueOf(input)
	}
	return
}

// GetBool accepts input as interface{}.
// GetBool is ExtractBool without error.
func GetBool(input interface{}, fnOpts ...FuncOption) (res bool) {
	res, _ = ExtractBool(getValFromInterface(input), fnOpts...)
	return
}

// ExtractBool extract the underlying bool value from the val of reflect.Value.
func ExtractBool(val reflect.Value, fnOpts ...FuncOption) (res bool, err error) {
	opt := NewOption().Assign(fnOpts...)
	res, err = extractBool(val, opt)
	return
}

// GetInt accepts input as interface{}.
// GetInt is ExtractInt without error.
func GetInt(input interface{}, fnOpts ...FuncOption) (result int64) {
	result, _ = ExtractInt(getValFromInterface(input), fnOpts...)
	return
}

// ExtractInt gets the underlying int value from val of reflect.Value.
func ExtractInt(val reflect.Value, fnOpts ...FuncOption) (result int64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractInt(val, opt)
	return
}

// GetUint accepts input as interface{}.
// GetUint is ExtractUint without error.
func GetUint(input interface{}, fnOpts ...FuncOption) (result uint64) {
	result, _ = ExtractUint(getValFromInterface(input), fnOpts...)
	return
}

// ExtractUint extracts the underlying uint value from val of reflect.Value.
func ExtractUint(val reflect.Value, fnOpts ...FuncOption) (result uint64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractUint(val, opt)
	return
}

// GetFloat accepts input as interface{}.
// GetFloat is ExtractFloat without error.
func GetFloat(input interface{}, fnOpts ...FuncOption) (result float64) {
	result, _ = ExtractFloat(getValFromInterface(input), fnOpts...)
	return
}

// ExtractFloat extracts the underlying float value from val of reflect.Value.
func ExtractFloat(val reflect.Value, fnOpts ...FuncOption) (result float64, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractFloat(val, opt)
	return
}

// GetComplex accepts input as interface{}.
// GetComplex is ExtractComplex without error.
func GetComplex(input interface{}, fnOpts ...FuncOption) (result complex128) {
	result, _ = ExtractComplex(getValFromInterface(input), fnOpts...)
	return
}

// ExtractComplex gets the underlying complex value from val of reflect.Value.
func ExtractComplex(val reflect.Value, fnOpts ...FuncOption) (result complex128, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractComplex(val, opt)
	return
}

// GetString accepts input as interface{}.
// GetString is ExtractString without error.
func GetString(input interface{}, fnOpts ...FuncOption) (result string) {
	result, _ = ExtractString(getValFromInterface(input), fnOpts...)
	return
}

// ExtractString gets the underlying string value from val of reflect.Value.
func ExtractString(val reflect.Value, fnOpts ...FuncOption) (result string, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractString(val, opt)
	return
}

// GetTime accepts input as interface{}.
// GetTime is ExtractTime without error.
func GetTime(input interface{}, fnOpts ...FuncOption) (result time.Time) {
	result, _ = ExtractTime(getValFromInterface(input), fnOpts...)
	return
}

// ExtractTime extracts the time from val of reflect.Value.
func ExtractTime(val reflect.Value, fnOpts ...FuncOption) (result time.Time, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractTime(val, opt)
	return
}

// GetDuration accepts input as interface{}.
// GetDuration is ExtractDuration without error.
func GetDuration(input interface{}, fnOpts ...FuncOption) (result time.Duration) {
	result, _ = ExtractDuration(getValFromInterface(input), fnOpts...)
	return
}

// ExtractDuration extracts time.Duration from val of reflect.Value.
func ExtractDuration(val reflect.Value, fnOpts ...FuncOption) (result time.Duration, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractDuration(val, opt)
	return
}

// GetURL accepts input as interface{}.
// GetURL is ExtractURL without error.
func GetURL(input interface{}, fnOpts ...FuncOption) (result *url.URL) {
	result, _ = ExtractURL(getValFromInterface(input), fnOpts...)
	return
}

// ExtractURL extracts *url.URL from val of reflect.Value.
func ExtractURL(val reflect.Value, fnOpts ...FuncOption) (result *url.URL, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = extractURL(val, opt)
	return
}

// TryGet accepts input as interface{}.
// TryGet is TryExtract without error.
func TryGet(input interface{}, fnOpts ...FuncOption) (result interface{}) {
	result, _ = TryExtract(getValFromInterface(input), fnOpts...)
	return
}

// TryExtract tries to extract the real value from the val of reflect.Value.
func TryExtract(val reflect.Value, fnOpts ...FuncOption) (result interface{}, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = tryExtract(val, opt)
	return
}
