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

	tempRes := GetKind(val)
	if IsKindTypeElemable(tempRes) {
		res = val.Type().Elem().Kind()
	} else if IsKindValueElemable(tempRes) {
		res = val.Elem().Kind()
	}
	return
}

// GetChildElemKind returns the child elems' (root child) kind of the val of reflect.Value.
func GetChildElemKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}

	tempRes := GetKind(val)
	if !IsKindTypeElemable(tempRes) || !IsKindValueElemable(tempRes) {
		return
	}

	// TODO: Change this to the reflect.Interface
	tempRes = getKindTypeElemable(val)
	res = tempRes
	return
}

func getKindTypeElemable(val reflect.Value) (res reflect.Kind) {
	elemType := val.Type().Elem()
	res = elemType.Kind()
	for IsKindTypeElemable(res) {
		elemType = elemType.Elem()
		res = elemType.Kind()
	}
	return
}

// GetChildElemPtrKind gets the child elements' (root child) ptr kind of the val of reflect.Value.
func GetChildElemPtrKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}

	tempRes := GetKind(val)
	if tempRes == reflect.Ptr {
		valType := val.Type()
		for tempRes == reflect.Ptr {
			valType = valType.Elem()
			tempRes = valType.Kind()
		}
		res = tempRes
	}
	return
}

// TODO: Add GetChildElemPtrAndInterfaceKind

// IsKindValueElemable checks the kind of reflect.Value that can call Elem method.
func IsKindValueElemable(kind reflect.Kind) bool {
	return kind == reflect.Ptr || kind == reflect.Interface
}

// IsKindTypeElemable checks the kind of reflect.Type that can call Elem method.
func IsKindTypeElemable(kind reflect.Kind) bool {
	return kind == reflect.Array ||
		kind == reflect.Chan ||
		kind == reflect.Map ||
		kind == reflect.Ptr ||
		kind == reflect.Slice
}
