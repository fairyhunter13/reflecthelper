package reflecthelper

import "reflect"

// GetElem gets the elem of the pointer val without initialize the pointer val.
// GetElem is similar to GetInitElem but without initialization.
func GetElem(val reflect.Value) (res reflect.Value) {
	res = val
	if !IsValueElemable(res) {
		return
	}
	if res.IsNil() {
		return
	}
	res = res.Elem()
	return
}

// GetChildElem is similar with GetInitChildElem but without initialize the child elem.
func GetChildElem(val reflect.Value) (res reflect.Value) {
	res = val
	var tempRes reflect.Value
	for IsValueElemable(res) {
		tempRes = GetElem(res)
		if res == tempRes {
			return
		}
		res = tempRes
	}
	return
}

// IsValueElemable checks whether the val of reflect.Value could call Elem method.
func IsValueElemable(val reflect.Value) bool {
	kind := GetKind(val)
	return kind == reflect.Ptr || kind == reflect.Interface
}
