package reflecthelper

import (
	"reflect"
	"time"
)

// List of reflect.Type used in this package
var (
	TypeRuneSlice = reflect.TypeOf([]rune{})
	TypeByteSlice = reflect.TypeOf([]byte{})
	TypeTimePtr   = reflect.TypeOf(&time.Time{})
	TypeTime      = reflect.TypeOf(time.Time{})
)

// IsTypeValueElemable checks if the type of the reflect.Value can call Elem.
func IsTypeValueElemable(val reflect.Value) bool {
	return IsKindTypeElemable(GetKind(val))
}

// IsTypeElemable checks wether the typ of reflect.Type can call Elem method.
func IsTypeElemable(typ reflect.Type) (res bool) {
	if typ == nil {
		return
	}

	res = IsKindTypeElemable(typ.Kind())
	return
}

// GetType is a wrapper for val.Type() to safely extract type if it is valid.
func GetType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	typ = val.Type()
	return
}

// GetElemType returns the elem type of a val of reflect.Value.
func GetElemType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	typ = val.Type()
	if IsTypeValueElemable(val) {
		typ = typ.Elem()
	}
	return
}

// GetElemTypeOfType returns the elem type of the input of reflect.Type.
func GetElemTypeOfType(input reflect.Type) (typ reflect.Type) {
	typ = input
	if typ == nil {
		return
	}

	if IsTypeElemable(typ) {
		typ = typ.Elem()
	}
	return
}

// GetChildElemType returns the child elems' (root child) type of the val of reflect.Value.
func GetChildElemType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	typ = val.Type()
	for IsTypeElemable(typ) {
		typ = typ.Elem()
	}
	return
}

// GetChildElemTypeOfType returns the child elems' (root child) type of the input of reflect.Type.
func GetChildElemTypeOfType(input reflect.Type) (typ reflect.Type) {
	typ = input
	if typ == nil {
		return
	}

	for IsTypeElemable(typ) {
		typ = typ.Elem()
	}
	return
}

// GetChildElemPtrType returns the child elems' (root child) ptr type of the val of reflect.Value.
func GetChildElemPtrType(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	typ = val.Type()
	res := typ.Kind()
	for res == reflect.Ptr {
		typ = typ.Elem()
		res = typ.Kind()
	}
	return
}

// GetChildElemPtrTypeOfType returns the child elems' (root child) ptr type of the input of reflect.Type.
func GetChildElemPtrTypeOfType(input reflect.Type) (typ reflect.Type) {
	typ = input
	if typ == nil {
		return
	}

	res := typ.Kind()
	for res == reflect.Ptr {
		typ = typ.Elem()
		res = typ.Kind()
	}
	return
}

// GetChildElemValueType returns the child elem's (root child) type of the val reflect.Value and it only works on ptr kind.
func GetChildElemValueType(val reflect.Value) (typ reflect.Type) {
	typ = GetChildElemPtrType(UnwrapInterfaceValue(val))
	return
}
