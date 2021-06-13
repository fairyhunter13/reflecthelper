package reflecthelper

import (
	"errors"
	"fmt"
	"reflect"
)

// List of all errors for reflecthelper.
var (
	ErrAssignerCantSet = errors.New("assigner doesn't have the ability to set the value")
)

func getErrOverflow(val reflect.Value) (err error) {
	err = fmt.Errorf(
		"assigner encounters overflow %s, underlying val: %s",
		GetKind(val),
		val.String(),
	)
	return
}

func getErrUnassignable(assigner reflect.Value, val reflect.Value) (err error) {
	assignerType := assigner.Type()
	valType := val.Type()
	if !valType.AssignableTo(assignerType) {
		err = fmt.Errorf(
			"error unassignable for kind: %s with val of reflect.Value, kind: %s type: %s val: %s",
			GetKind(assigner),
			GetKind(val),
			val.Type(),
			val,
		)
	}
	return
}

func getErrOverflowedLength(assigner reflect.Value, val reflect.Value, isSlice bool) (err error) {
	if isSlice {
		return
	}

	if assigner.Len() < val.Len() {
		err = fmt.Errorf(
			"error length of assigner is smaller than the length of val, assigner type:%s length:%d, val type:%s length:%d",
			assigner.Type(),
			assigner.Len(),
			val.Type(),
			val.Len(),
		)
	}
	return
}

func getErrCanInterface(val reflect.Value) (err error) {
	if !val.CanInterface() {
		err = fmt.Errorf(
			"can't get the interface from the val of reflect.Value, kind: %s type: %s val: %s",
			GetKind(val),
			val.Type(),
			val,
		)
	}
	return
}

func getErrIsValid(val reflect.Value) (err error) {
	if !val.IsValid() {
		err = fmt.Errorf(
			"the val of reflect.Value is invalid, underlying val: %s",
			val.String(),
		)
	}
	return
}

func getErrUnimplementedExtract(val reflect.Value) (err error) {
	err = fmt.Errorf(
		"error unimplemented extraction for val of reflect.Value,kind: %s type: %s val: %s",
		GetKind(val),
		val.Type(),
		val,
	)
	return
}

func getErrUnimplementedAssign(assigner reflect.Value, val reflect.Value) (err error) {
	err = fmt.Errorf(
		"error unimplemented assignment for kind: %s with val of reflect.Value, kind: %s type: %s val: %s",
		GetKind(assigner),
		GetKind(val),
		val.Type(),
		val,
	)
	return
}
