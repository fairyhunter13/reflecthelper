package reflecthelper

import "reflect"

// GetInitElem gets the element of a pointer value.
// It initialize the element of a pointer value if it is nil.
func GetInitElem(val reflect.Value) (res reflect.Value) {
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

// GetInitChildElem gets the child elem if it is a pointer with an element of pointer.
// It also initializes the child elem if it is CanSet and IsNil.
func GetInitChildElem(val reflect.Value) (res reflect.Value) {
	res = val
	kind := GetKind(res)
	var tempRes reflect.Value
	for kind == reflect.Ptr {
		tempRes = GetInitElem(res)
		if res == tempRes {
			return
		}
		res = tempRes
		kind = GetKind(res)
	}
	return
}

// GetElem gets the elem of the val without initialize the val.
// GetElem is similar to GetInitElem but without initialization.
func GetElem(val reflect.Value) (res reflect.Value) {
	res = val
	if GetKind(res) != reflect.Ptr {
		return
	}
	res = reflect.Indirect(res)
	return
}

// GetChildElem is similar with GetInitChildElem but without initialize the child elem.
func GetChildElem(val reflect.Value) (res reflect.Value) {
	res = val
	kind := GetKind(res)
	var tempRes reflect.Value
	for kind == reflect.Ptr {
		tempRes = GetElem(val)
		if res == tempRes {
			return
		}
		res = tempRes
		kind = GetKind(res)
	}
	return
}
