package reflecthelper

import (
	"reflect"
)

// IterateStructFunc is function to iterate each field of structInput.
type IterateStructFunc func(structInput reflect.Value, field reflect.Value)

// ReflectStruct is a custom struct for representing reflect.Struct.
type ReflectStruct struct {
	reflect.Value
}

// Iterate iterates all the field in this struct.
func (s *ReflectStruct) Iterate(fns ...IterateStructFunc) *ReflectStruct {
	// TODO: Add logic in here.
	return s
}

// IteratePanic iterates all the field in this struct by returning err when the iteration function panics.
func (s *ReflectStruct) IteratePanic(fns ...IterateStructFunc) (err error) {
	defer RecoverFn(&err)

	// TODO: Add logic in here.
	return
}

// CastStruct casts the val of reflect.Value to the ReflectStruct.
func CastStruct(val reflect.Value) (res ReflectStruct, err error) {
	val = GetChildElem(val)
	err = getErrIsValid(val)
	if err != nil {
		return
	}

	switch GetKind(val) {
	case reflect.Struct:
		res = ReflectStruct{val}
	default:
		err = getErrUnimplementedCasting(val, reflect.Struct)
	}
	return
}
