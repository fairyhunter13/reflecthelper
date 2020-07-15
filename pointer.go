package reflecthelper

import "reflect"

// GetInitElem gets the element of a pointer value.
// It initialize the element of a pointer value if it is nil.
func GetInitElem(v reflect.Value) (res reflect.Value) {
	kind := GetKind(v)
	if kind != reflect.Ptr {
		res = v
		return
	}
	if v.IsNil() {
		if !v.CanSet() {
			res = v
			return
		}
		v.Set(reflect.New(v.Type().Elem()))
	}

	res = reflect.Indirect(v)
	return
}

// GetChildElem gets the child elem if it is a pointer with an element of pointer.
func GetChildElem(v reflect.Value) (res reflect.Value) {
	res = v
	kind := GetKind(res)
	for kind == reflect.Ptr {
		res = GetInitElem(res)
		kind = GetKind(res)
		if !res.IsValid() || !res.CanSet() {
			return
		}
	}
	return
}
