package reflecthelper

import (
	"reflect"
)

// IterateStructFunc is function to iterate each field of structInput.
type IterateStructFunc func(structInput reflect.Value, field reflect.Value)

// ReflectStruct is a custom struct for representing reflect.Struct.
type ReflectStruct struct {
	reflect.Value
	kind reflect.Kind
}

func (s *ReflectStruct) isStruct() bool {
	return IsKindStruct(s.kind)
}

// Iterate iterates all the field in this struct.
func (s *ReflectStruct) Iterate(fns ...IterateStructFunc) *ReflectStruct {
	if !s.isStruct() {
		return s
	}

	// TODO: Add logic in here.
	return s
}

// IteratePanic iterates all the field in this struct by returning err when the iteration function panics.
func (s *ReflectStruct) IteratePanic(fns ...IterateStructFunc) (err error) {
	defer RecoverFn(&err)
	if !s.isStruct() {
		return
	}

	// TODO: Add logic in here.
	return
}

// CastStruct casts the val of reflect.Value to the ReflectStruct.
func CastStruct(val reflect.Value) (res ReflectStruct) {
	val = GetChildElem(val)

	kind := GetKind(val)
	if IsKindStruct(kind) {
		res = ReflectStruct{
			Value: val,
			kind:  kind,
		}
	}
	return
}
