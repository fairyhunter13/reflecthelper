package reflecthelper

import "github.com/reiver/go-cast"

func castString(val interface{}) (result string, err error) {
	err = getErrExtractNil(val)
	if err != nil {
		return
	}
	result, err = cast.String(val)
	return
}

func castComplex128(val interface{}) (result complex128, err error) {
	err = getErrExtractNil(val)
	if err != nil {
		return
	}
	result, err = cast.Complex128(val)
	return
}

func castFloat64(val interface{}) (result float64, err error) {
	err = getErrExtractNil(val)
	if err != nil {
		return
	}
	result, err = cast.Float64(val)
	return
}

func castUint64(val interface{}) (result uint64, err error) {
	err = getErrExtractNil(val)
	if err != nil {
		return
	}
	result, err = cast.Uint64(val)
	return
}

func castInt64(val interface{}) (result int64, err error) {
	err = getErrExtractNil(val)
	if err != nil {
		return
	}
	result, err = cast.Int64(val)
	return
}

func castBool(val interface{}) (result bool, err error) {
	err = getErrExtractNil(val)
	if err != nil {
		return
	}
	result, err = cast.Bool(val)
	return
}
