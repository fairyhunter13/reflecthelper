package reflecthelper

import "reflect"

// GetKind gets the kind of the val of reflect.Value.
func GetKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}

	res = val.Type().Kind()
	return
}

// GetElemKind gets the elem kind from the val of reflect.Value.
func GetElemKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}

	res = GetKind(val)
	if IsKindTypeElemable(res) {
		res = val.Type().Elem().Kind()
	} else if res == reflect.Interface {
		res = val.Elem().Kind()
	}
	return
}

// GetChildElemTypeKind returns the child elems' (root child) kind of the type of val reflect.Value.
func GetChildElemTypeKind(val reflect.Value) (res reflect.Kind) {
	if !val.IsValid() {
		return
	}

	val = UnwrapInterfaceValue(val)
	res = GetKind(val)
	if !IsKindTypeElemable(res) {
		return
	}

	res = getChildElemTypeKind(val)
	return
}

func getChildElemTypeKind(val reflect.Value) (res reflect.Kind) {
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

	res = GetKind(val)
	valType := val.Type()
	for res == reflect.Ptr {
		valType = valType.Elem()
		res = valType.Kind()
	}
	return
}

// GetChildElemValueKind gets the child elements' (root child) kind of the val reflect.Value and it only works on ptr kind.
func GetChildElemValueKind(val reflect.Value) (res reflect.Kind) {
	res = GetChildElemPtrKind(UnwrapInterfaceValue(val))
	return
}

// IsKindValueElemable checks the kind of reflect.Value that can call Elem method.
func IsKindValueElemable(kind reflect.Kind) bool {
	return kind == reflect.Ptr || kind == reflect.Interface
}

// IsValueElemable checks whether the val of reflect.Value could call Elem method.
func IsValueElemable(val reflect.Value) bool {
	return IsKindValueElemable(GetKind(val))
}

// IsValueElemableParentElem checks whether the res have elemable kind for parent and elem.
func IsValueElemableParentElem(res reflect.Value) bool {
	return IsKindValueElemable(GetKind(res)) && IsKindValueElemable(GetElemKind(res))
}

// IsKindTypeElemable checks the kind of reflect.Type that can call Elem method.
func IsKindTypeElemable(kind reflect.Kind) bool {
	return kind == reflect.Array ||
		kind == reflect.Chan ||
		kind == reflect.Map ||
		kind == reflect.Ptr ||
		kind == reflect.Slice
}

// IsKindBool checks whether the kind is bool or not.
func IsKindBool(kind reflect.Kind) bool {
	return kind == reflect.Bool
}

// IsKindValueBytesSlice checks whether the val of reflect.Value is byte slice.
func IsKindValueBytesSlice(val reflect.Value) bool {
	if !val.IsValid() {
		return false
	}

	if !IsKindSlice(GetKind(val)) {
		return false
	}

	return GetElemKind(val) == reflect.Uint8
}

// IsKindSlice checks whether the kind is slice or not.
func IsKindSlice(kind reflect.Kind) bool {
	return kind == reflect.Slice
}

// IsKindArray checks whether the kind is array or not.
func IsKindArray(kind reflect.Kind) bool {
	return kind == reflect.Array
}

// IsKindComplex checks whether the kind is complex or not.
func IsKindComplex(kind reflect.Kind) bool {
	return kind >= reflect.Complex64 && kind <= reflect.Complex128
}

// IsKindFloat checks whether the kind is float or not.
func IsKindFloat(kind reflect.Kind) bool {
	return kind >= reflect.Float32 && kind <= reflect.Float64
}

// IsKindInt checks whether the kind is int or not.
func IsKindInt(kind reflect.Kind) bool {
	return kind >= reflect.Int && kind <= reflect.Int64
}

// IsKindUnsafePointer checks whether the kind is unsafe ptr or not.
func IsKindUnsafePointer(kind reflect.Kind) bool {
	return kind == reflect.UnsafePointer
}

// IsKindString checks whether the kind is string or not.
func IsKindString(kind reflect.Kind) bool {
	return kind == reflect.String
}

// IsKindUint checks whether the kind is uint or not.
func IsKindUint(kind reflect.Kind) bool {
	return kind >= reflect.Uint && kind <= reflect.Uintptr
}

// IsKindPtr checks whether the input kind is reflect.Ptr.
func IsKindPtr(kind reflect.Kind) bool {
	return kind == reflect.Ptr
}

// IsKindStruct checks whether the input kind is reflect.Struct.
func IsKindStruct(kind reflect.Kind) bool {
	return kind == reflect.Struct
}
