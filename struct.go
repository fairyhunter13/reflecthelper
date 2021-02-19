package reflecthelper

import "reflect"

// ReflectStruct is a custom struct for representing reflect.Struct.
type ReflectStruct struct {
	reflect.Value
}

// CastStruct casts the val of reflect.Value to the ReflectStruct.
func CastStruct(val reflect.Value) (res ReflectStruct, err error) {
	err = getErrIsValid(val)
	if err != nil {
		return
	}

	// TODO: Add logic in here
	return
}
