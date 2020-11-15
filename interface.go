package reflecthelper

import "reflect"

// UnwrapInterfaceValue unwraps the elem of val reflect.Value with the kind reflect.Interface.
// If the val of reflect.Value contains multi level interface,
// then it unwraps until the child of val reflect.Value doesn't have the kind of reflect.Interface.
func UnwrapInterfaceValue(val reflect.Value) (res reflect.Value) {
	res = val
	if !res.IsValid() {
		return
	}

	kind := GetKind(res)
	for kind == reflect.Interface {
		res = res.Elem()
		kind = GetKind(res)
	}
	return
}
