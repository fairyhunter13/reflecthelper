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
	if IsKindTypeElemable(tempRes) {
		elemType := val.Type().Elem()
		tempRes = elemType.Kind()
		for IsKindTypeElemable(tempRes) {
			elemType = elemType.Elem()
			tempRes = elemType.Kind()
		}
	} else if IsKindValueElemable(tempRes) {
		childVal := val.Elem()
		tempRes = childVal.Kind()
		for IsKindValueElemable(tempRes) {
			childVal = childVal.Elem()
			tempRes = childVal.Kind()
		}
	} else {
		return
	}
	res = tempRes
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
	} else {
		return
	}
	res = tempRes
	return
}

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
