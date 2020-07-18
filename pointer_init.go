package reflecthelper

import "reflect"

// GetInitIndirect gets the element of a pointer value.
// It initialize the element of a pointer value if it is nil.
func GetInitIndirect(val reflect.Value) (res reflect.Value) {
	res = val
	if GetKind(res) != reflect.Ptr {
		return
	}
	if res.IsNil() {
		if !res.CanSet() {
			return
		}
		res.Set(reflect.New(res.Type().Elem()))
	}

	res = reflect.Indirect(res)
	return
}

// GetInitChildIndirect gets the child elem if it is a pointer with an element of pointer.
// It also initializes the child elem if it is CanSet and IsNil.
func GetInitChildIndirect(val reflect.Value) (res reflect.Value) {
	res = val
	kind := GetKind(res)
	var tempRes reflect.Value
	for kind == reflect.Ptr {
		tempRes = GetInitIndirect(res)
		if res == tempRes {
			return
		}
		res = tempRes
		kind = GetKind(res)
	}
	return
}
