package reflecthelper

import "reflect"

// List of reflect.Type used in this package
var (
	TypeRuneSlice = reflect.TypeOf([]rune{})
	TypeByteSlice = reflect.TypeOf([]byte{})
)

// IsTypeElemable checks if the type of the reflect.Value can call Elem
func IsTypeElemable(val reflect.Value) bool {
	switch GetKind(val) {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

// GetElemType returns the elem type of a val of reflect.Value
func GetElemType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	switch GetKind(val) {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		typ = val.Type().Elem()
	}
	return
}
