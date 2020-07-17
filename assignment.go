package reflecthelper

import (
	"fmt"
	"reflect"
)

func checkValid(val reflect.Value) (err error) {
	if !val.IsValid() {
		err = fmt.Errorf(
			"The val of reflect.Value is invalid, underlying val: %s",
			val.String(),
		)
	}
	return
}

func checkAssign(assigner reflect.Value, val reflect.Value) (err error) {
	err = checkValid(assigner)
	if err != nil {
		return
	}
	err = checkValid(val)
	if err != nil {
		return
	}
	if !assigner.CanSet() {
		err = ErrAssignerCantSet
		return
	}
	if !val.CanInterface() {
		err = fmt.Errorf(
			"The val can't extract the underlying value as interface, underlying val: %s kind: %s",
			val.String(),
			GetKind(val),
		)
	}
	return
}

// AssignReflect assigns the val of the reflect.Value to the assigner.
// This function asserts that the assigner Kind is same as the val Kind.
func AssignReflect(assigner reflect.Value, val reflect.Value) (err error) {
	err = checkAssign(assigner, val)
	if err != nil {
		return
	}
	return
}

func tryAssign(assigner reflect.Value, val reflect.Value) (err error) {
	switch GetKind(assigner) {
	case reflect.Bool:
		var result bool
		result, err = getBool(val)
		if err != nil {
			return
		}
		assigner.SetBool(result)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var result int64
		result, err = getInt(val)
		if err != nil {
			return
		}
		if assigner.OverflowInt(result) {
			err = getErrOverflow(assigner)
			return
		}
		assigner.SetInt(result)
	case reflect.String:
		var result string
		result, err = getString(val)
		if err != nil {
			return
		}
		assigner.SetString(result)
	}
	return
}
