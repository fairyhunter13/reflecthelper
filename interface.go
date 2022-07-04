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

// IsNil checks whether the input val is nil for any type.
func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}

	return IsValueNil(getValFromInterface(val))
}

// IsPtr checks whether the input interface{} is a reflect.Ptr or not.
// The pointer in golang can be represented by reflect.Ptr.
func IsPtr(in interface{}) bool {
	return IsKindPtr(GetKind(getValFromInterface(in)))
}

// GetTypeInterface is a wrapper for GetType by passing an interface instead of reflect.Value.
func GetTypeInterface(in interface{}) reflect.Type {
	return GetType(getValFromInterface(in))
}

// GetKindInterface is a wrapper for GetKind by passing an interface instead of reflect.Value.
func GetKindInterface(in interface{}) reflect.Kind {
	return GetKind(getValFromInterface(in))
}
