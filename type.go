package reflecthelper

import "reflect"

// List of reflect.Type used in this package
var (
	TypeRuneSlice = reflect.TypeOf([]rune{})
	TypeByteSlice = reflect.TypeOf([]byte{})
)

// IsTypeValueElemable checks if the type of the reflect.Value can call Elem
func IsTypeValueElemable(val reflect.Value) bool {
	return IsKindTypeElemable(GetKind(val))
}

// IsTypeElemable checks wether the typ of reflect.Type can call Elem method.
func IsTypeElemable(typ reflect.Type) (res bool) {
	if typ == nil {
		return
	}

	return
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

// GetChildElemType returns the child of elem type of the val of reflect.Value.
func GetChildElemType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	// TODO: Add logic in here
	return
}
