package reflecthelper

import (
	"reflect"
)

// GetInitElem gets the element of a pointer value.
// It initialize the element of a pointer value if it is nil.
func GetInitElem(val reflect.Value) (res reflect.Value) {
	res = val
	if !IsValueElemable(res) {
		return
	}
	if res.IsNil() {
		if !res.CanSet() || !IsTypeValueElemable(res) {
			return
		}
		res.Set(reflect.New(res.Type().Elem()))
	}

	res = res.Elem()
	return
}

// GetInitChildElem gets the child elem (root child) if it is a pointer with an element of pointer.
// It also initializes the child elem if it is CanSet and IsNil.
func GetInitChildElem(val reflect.Value) (res reflect.Value) {
	res = val
	var tempRes reflect.Value
	for IsValueElemable(res) {
		tempRes = GetInitElem(res)
		if res == tempRes {
			return
		}
		res = tempRes
	}
	return
}

// GetInitChildPtrElem is similar with GetInitChildElem but the function stops when the elem is ptr and the elem of that ptr is non ptr.
func GetInitChildPtrElem(val reflect.Value) (res reflect.Value) {
	res = val
	var tempRes reflect.Value
	for IsValueElemableParentElem(res) {
		tempRes = GetInitElem(res)
		if res == tempRes {
			return
		}
		res = tempRes
	}
	return
}
