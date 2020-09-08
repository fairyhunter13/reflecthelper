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
	// DefaultComplexBitSize is the default bit size for the complex128 type.
	DefaultComplexBitSize = 128
	// DefaultBaseSystem is the default base system used for decimal in this package.
	DefaultBaseSystem = 10
)

func getDefaultFloatStr(floatVal float64) (result string) {
	result = strconv.FormatFloat(floatVal, 'g', DefaultFloatPrecision, DefaultBitSize)
	return
}

func getDefaultString(val reflect.Value) (result string) {
	result = fmt.Sprintf("%v", val.Interface())
	return
}
