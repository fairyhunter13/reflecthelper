package reflecthelper

import (
	"reflect"
)

// IterateStructFunc is function to iterate each field of structInput.
type IterateStructFunc func(structInput reflect.Value, field reflect.Value)

// Value is a custom struct for representing reflect.Value.
type Value struct {
	reflect.Value
	kind reflect.Kind
}

func (s *Value) isStruct() bool {
	return IsKindStruct(s.kind)
}

// IterateStruct iterates all the field in this struct.
func (s *Value) IterateStruct(fns ...IterateStructFunc) *Value {
	if !s.isStruct() {
		return s
	}

	// TODO: Add logic in here.
	return s
}

// IterateStructPanic iterates all the field in this struct by returning err when the iteration function panics.
func (s *Value) IterateStructPanic(fns ...IterateStructFunc) (err error) {
	defer RecoverFn(&err)
	if !s.isStruct() {
		return
	}

	// TODO: Add logic in here.
	return
}

// CastStruct casts the val of reflect.Value to the ReflectStruct.
func CastStruct(val reflect.Value) (res Value) {
	val = GetChildElem(val)

	kind := GetKind(val)
	if IsKindStruct(kind) {
		res = Value{
			Value: val,
			kind:  kind,
		}
	}
	return
}
