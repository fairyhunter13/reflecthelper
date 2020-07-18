package reflecthelper

import "reflect"

// GetIndirect gets the elem of the pointer val without initialize the pointer val.
// GetIndirect is similar to GetInitElem but without initialization.
func GetIndirect(val reflect.Value) (res reflect.Value) {
	res = val
	if GetKind(res) != reflect.Ptr {
		return
	}
	if res.IsNil() {
		return
	}
	res = reflect.Indirect(res)
	return
}

// GetChildIndirect is similar with GetInitChildElem but without initialize the child elem.
func GetChildIndirect(val reflect.Value) (res reflect.Value) {
	res = val
	kind := GetKind(res)
	var tempRes reflect.Value
	for kind == reflect.Ptr {
		tempRes = GetIndirect(res)
		if res == tempRes {
			return
		}
		res = tempRes
		kind = GetKind(res)
	}
	return
}
