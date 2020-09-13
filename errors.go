package reflecthelper

import (
	"errors"
	"fmt"
	"reflect"
)

// List of all errors for reflecthelper.
var (
	ErrAssignerCantSet = errors.New("Assigner doesn't have the ability to set the value")
)

func getErrOverflow(val reflect.Value) (err error) {
	err = fmt.Errorf(
		"Assigner encounters overflow %s, underlying val: %s",
		GetKind(val),
		val.String(),
	)
	return
}

func getErrCanInterface(val reflect.Value) (err error) {
	if !val.CanInterface() {
		err = fmt.Errorf(
			"Can't get the interface from the val of reflect.Value, kind: %s type: %s val: %s",
			GetKind(val),
			val.Type(),
			val,
		)
	}
	return
}

func getErrUnimplementedAssign(assigner reflect.Value, val reflect.Value) (err error) {
	err = fmt.Errorf(
		"Error unimplemented assignment for kind: %s with val of reflect.Value, kind: %s type: %s val: %s",
		GetKind(assigner),
		GetKind(val),
		val.Type(),
		val,
	)
	return
}

func getErrOverflowedLength(assigner reflect.Value, val reflect.Value) (err error) {
	err = fmt.Errorf(
		"Error length of assigner is smaller than the length of val, assigner type:%s length:%d, val type:%s length:%d",
		assigner.Type(),
		assigner.Len(),
		val.Type(),
		val.Len(),
	)
	return
}

func getErrUnimplementedExtract(val reflect.Value) (err error) {
	err = fmt.Errorf(
		"Error unimplemented extraction for val of reflect.Value,kind: %s type: %s val: %s",
		GetKind(val),
		val.Type(),
		val,
	)
	return
}
