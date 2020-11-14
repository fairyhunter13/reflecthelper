package reflecthelper

import "reflect"

// IsReflectZero check if the val of reflect.Value is a reflect.Zero of it's type.
// This operation is solely based on reflect.Zero, not to check if the actual underlying value is zero.
// To check if the value of reflect.Value is zero, use IsZero or IsValueZero.
func IsReflectZero(val reflect.Value) (result bool) {
	if !val.IsValid() {
		return
	}

	zero := reflect.Zero(val.Type())
	result = val == zero
	return
}

// IsInterfaceReflectZero checks whether the interface is nil.
// It also checks if the reflect.Value of val is reflect.Zero of it's type.
func IsInterfaceReflectZero(val interface{}) (result bool) {
	if val == nil {
		result = true
		return
	}

	return IsReflectZero(reflect.ValueOf(val))
}

// SetReflectZero sets the val to the reflect.Zero of its type.
func SetReflectZero(val reflect.Value) {
	if !val.IsValid() {
		return
	}

	if val.CanSet() {
		val.Set(reflect.Zero(val.Type()))
	}
}
