package reflecthelper

import (
	"math"
	"reflect"
	"time"
)

// Zeroable is a contract to specifies the Zero attribute of a custom type.
type Zeroable interface {
	IsZero() bool
}

var nilTime *time.Time

// IsZero returns false if k is nil or has a zero value
func IsZero(k interface{}) bool {
	if k == nil {
		return true
	}

	switch val := k.(type) {
	case bool:
		return !val
	case int:
		return val == 0
	case int8:
		return val == 0
	case int16:
		return val == 0
	case int32:
		return val == 0
	case int64:
		return val == 0
	case uint:
		return val == 0
	case uint8:
		return val == 0
	case uint16:
		return val == 0
	case uint32:
		return val == 0
	case uint64:
		return val == 0
	case uintptr:
		return val == 0
	case float32:
		return val == 0
	case float64:
		return val == 0
	case complex64:
		return val == 0
	case complex128:
		return val == 0
	case string:
		return val == ""
	case *time.Time:
		return val == nilTime || IsTimeZero(*val)
	case time.Time:
		return IsTimeZero(val)
	case Zeroable:
		if IsNil(val) {
			return true
		}
		return val == nil || val.IsZero()
	case reflect.Value: // for go version less than 1.13 because reflect.Value has no method IsZero
		return IsValueZero(val)
	}

	return IsValueZero(reflect.ValueOf(k))
}

var zeroType = reflect.TypeOf((*Zeroable)(nil)).Elem()

// IsValueZero check the reflect.Value if it is zero based on it's kind.
func IsValueZero(v reflect.Value) bool {
	switch GetKind(v) {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		complexNum := v.Complex()
		return math.Float64bits(real(complexNum)) == 0 && math.Float64bits(imag(complexNum)) == 0
	case reflect.Array, reflect.Slice:
		return IsArrayZero(v)
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return IsValueZero(v.Elem())
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		return IsStructZero(v)
	case reflect.Invalid:
		return true
	}
	return false
}

// IsStructZero checks if the struct is zero.
func IsStructZero(v reflect.Value) bool {
	if !v.IsValid() || !IsKindStruct(GetKind(v)) || v.NumField() == 0 {
		return true
	}

	if v.Type().Implements(zeroType) {
		f := v.MethodByName("IsZero")
		if f.IsValid() {
			res := f.Call(nil)
			return len(res) == 1 && res[0].Bool()
		}
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsValid() {
			continue
		}
		if field.CanInterface() && !IsZero(field.Interface()) {
			return false
		}
	}
	return true
}

// IsArrayZero checks if the array is empty.
func IsArrayZero(v reflect.Value) bool {
	if !v.IsValid() || !IsKindList(GetKind(v)) || v.Len() == 0 {
		return true
	}

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if !elem.IsValid() {
			continue
		}
		if elem.CanInterface() && !IsZero(elem.Interface()) {
			return false
		}
	}

	return true
}

// List of constants for the zero time.
const (
	ZeroTime0 = "0000-00-00 00:00:00"
	ZeroTime1 = "0001-01-01 00:00:00"
)

// IsTimeZero checks if the time is zero.
func IsTimeZero(t time.Time) bool {
	return t.IsZero() || t.Format("2006-01-02 15:04:05") == ZeroTime0 ||
		t.Format("2006-01-02 15:04:05") == ZeroTime1
}

// IsPtrValueZero overrides the default behavior for the reflect.Ptr case in the IsValueZero method.
func IsPtrValueZero(val reflect.Value) bool {
	if val.Kind() != reflect.Ptr {
		return IsValueZero(val)
	}

	return val.IsNil()
}
