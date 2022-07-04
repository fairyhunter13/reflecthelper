package reflecthelper

import (
	"net"
	"net/url"
	"reflect"
	"time"
)

// List of reflect.Type used in this package
var (
	TypeRuneSlice   = reflect.TypeOf([]rune{})
	TypeByteSlice   = reflect.TypeOf([]byte{})
	TypeTimePtr     = reflect.TypeOf(new(time.Time))
	TypeTime        = reflect.TypeOf(time.Time{})
	TypeDurationPtr = reflect.TypeOf(new(time.Duration))
	TypeDuration    = reflect.TypeOf(time.Duration(0))
	TypeURLPtr      = reflect.TypeOf(new(url.URL))
	TypeURL         = reflect.TypeOf(url.URL{})
	TypeIPPtr       = reflect.TypeOf(new(net.IP))
	TypeIP          = reflect.TypeOf(net.IP{})
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

// GetTypeElem returns the elem type of a val of reflect.Value.
func GetTypeElem(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	typ = val.Type()
	if IsTypeValueElemable(val) {
		typ = typ.Elem()
	}
	return
}

// GetTypeElemType returns the elem type of the input of reflect.Type.
func GetTypeElemType(input reflect.Type) (typ reflect.Type) {
	typ = input
	if typ == nil {
		return
	}

	if IsTypeElemable(typ) {
		typ = typ.Elem()
	}
	return
}

// GetTypeChildElem returns the child elems' (root child) type of the val of reflect.Value.
func GetTypeChildElem(val reflect.Value) (typ reflect.Type) {
	if !val.IsValid() {
		return
	}

	typ = val.Type()
	for IsTypeElemable(typ) {
		typ = typ.Elem()
	}
	return
}

// GetTypeChildElemType returns the child elems' (root child) type of the input of reflect.Type.
func GetTypeChildElemType(input reflect.Type) (typ reflect.Type) {
	typ = input
	if typ == nil {
		return
	}

	for IsTypeElemable(typ) {
		typ = typ.Elem()
	}
	return
}

// GetTypeChildElemPtr returns the child elems' (root child) ptr type of the val of reflect.Value.
func GetTypeChildElemPtr(val reflect.Value) (typ reflect.Type) {
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

// GetTypeChildElemPtrType returns the child elems' (root child) ptr type of the input of reflect.Type.
func GetTypeChildElemPtrType(input reflect.Type) (typ reflect.Type) {
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

// GetTypeChildElemValue returns the child elem's (root child) type of the val reflect.Value and it only works on ptr kind.
func GetTypeChildElemValue(val reflect.Value) (typ reflect.Type) {
	typ = GetTypeChildElemPtr(UnwrapInterfaceValue(val))
	return
}

// IsTypeValueDuration checks whether the type of val reflect.Value is time.Duration or *time.Duration.
func IsTypeValueDuration(val reflect.Value) bool {
	typeVal := GetType(val)
	return typeVal == TypeDuration || typeVal == TypeDurationPtr
}

// IsTypeValueTime checks whether the type of val reflect.Value is time.Time or *time.Time.
func IsTypeValueTime(val reflect.Value) bool {
	typeVal := GetType(val)
	return typeVal == TypeTime || typeVal == TypeTimePtr
}

// IsTypeValueURL checks whether the type of val reflect.Value is url.URL or *url.URL.
func IsTypeValueURL(val reflect.Value) bool {
	typeVal := GetType(val)
	return typeVal == TypeURL || typeVal == TypeURLPtr
}

// IsTypeValueIP checks whether the type of val reflect.Value is net.IP or *net.IP.
func IsTypeValueIP(val reflect.Value) bool {
	typeVal := GetType(val)
	return typeVal == TypeIP || typeVal == TypeIPPtr
}
