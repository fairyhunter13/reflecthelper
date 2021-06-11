package reflecthelper

import "reflect"

func (s *Value) isArrayOrSlice() bool {
	return IsKindSlice(s.kind) || IsKindArray(s.kind)
}

// IterArrSliceFn is function type to iterate each field of array or slice.
type IterArrSliceFn func(parent reflect.Value, index int, field reflect.Value)

// IterArrSliceErrFn is function type to iterate each field of array or slice and returning error if needed.
type IterArrSliceErrFn func(parent reflect.Value, index int, field reflect.Value) error

func (s *Value) IterateArrayOrSlice(fns ...IterArrSliceFn) *Value {
	if !s.isArrayOrSlice() {
		return s
	}

	// TODO: Add logic in here
	return s
}
