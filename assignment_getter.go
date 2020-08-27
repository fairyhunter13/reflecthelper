package reflecthelper

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	// DefaultFloatPrecision specifies the default precision used in this package.
	// This is the default maximum precision.
	DefaultFloatPrecision = -1
	// DefaultBitSize is the default bit size used for the conversion in this package.
	DefaultBitSize = 64
	// DefaultBaseSystem is the default base system used for decimal in this package.
	DefaultBaseSystem = 10
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

func getDefaultFloatStr(floatVal float64) (result string) {
	result = strconv.FormatFloat(floatVal, 'g', DefaultFloatPrecision, DefaultBitSize)
	return
}

func getDefaultString(val reflect.Value) (result string) {
	result = fmt.Sprintf("%v", val.Interface())
	return
}
