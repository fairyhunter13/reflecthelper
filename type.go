package reflecthelper

import "reflect"

// List of reflect.Type used in this package
var (
	TypeRuneSlice = reflect.TypeOf([]rune{})
	TypeByteSlice = reflect.TypeOf([]byte{})
)

// GetElemType returns the elem type of a val of reflect.Value
func GetElemType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	kind := GetKind(val)
	switch kind {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		typ = val.Type().Elem()
	}
	return
}
