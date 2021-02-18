package reflecthelper

import "reflect"

// AssignReflect assigns the val of the reflect.Value to the assigner.
// This function asserts that the assigner Kind is same as the val Kind.
func AssignReflect(assigner reflect.Value, val reflect.Value) (err error) {
	err = assignReflect(assigner, val, NewDefaultOption())
	return
}
