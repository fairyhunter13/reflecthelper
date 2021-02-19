package reflecthelper

import "reflect"

// AssignReflect assigns the val of the reflect.Value to the assigner.
// This function asserts that the assigner Kind is same as the val Kind.
func AssignReflect(assigner reflect.Value, val reflect.Value, funcOpts ...FuncOption) (err error) {
	opt := NewOption().Assign(funcOpts...)
	err = assignReflect(assigner, val, opt)
	return
}
