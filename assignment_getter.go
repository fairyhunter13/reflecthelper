package reflecthelper

import (
	"fmt"
	"reflect"
	"strconv"
)

func getDefaultFloatStr(floatVal float64, opt *Option) (result string) {
	result = strconv.FormatFloat(floatVal, opt.FloatFormat, opt.FloatPrecision, opt.BitSize)
	return
}

func getDefaultString(val reflect.Value) (result string) {
	result = fmt.Sprintf("%v", val.Interface())
	return
}
