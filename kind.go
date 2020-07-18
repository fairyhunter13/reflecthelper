package reflecthelper

import "reflect"

// GetKind gets the kind of the reflect.Value.
func GetKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}
	res = val.Type().Kind()
	return
}

// GetElemKind gets the elem kind from ptr type.
func GetElemKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}

	res = GetKind(val)
	switch res {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		res = val.Type().Elem().Kind()
	}
	return
}
