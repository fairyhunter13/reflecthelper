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

// GetNilElem is similar with GetElem but it doesn't check if reflect.Ptr or reflect.Interface is nil.
func GetNilElem(val reflect.Value) (res reflect.Value) {
	res = val
	if !IsValueElemable(res) {
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

// GetChildNilElem is similar with GetChildElem but it uses GetNilElem function.
func GetChildNilElem(val reflect.Value) (res reflect.Value) {
	res = val
	for IsValueElemable(res) {
		res = GetNilElem(res)
	}
	return
}

// IsValueElemable checks whether the val of reflect.Value could call Elem method.
func IsValueElemable(val reflect.Value) bool {
	return IsKindValueElemable(GetKind(val))
}
