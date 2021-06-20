package reflecthelper

import (
	"net"
	"net/url"
	"reflect"
	"time"

	"github.com/fairyhunter13/task/v2"
	"github.com/mitchellh/mapstructure"
)

func assignReflect(assigner reflect.Value, val reflect.Value, opt *Option) (err error) {
	var oriAssigner reflect.Value
	if assigner.CanSet() {
		oriAssigner = assigner
	} else {
		oriAssigner = GetChildElem(assigner)
	}

	err = checkAssigner(oriAssigner)
	if err != nil {
		return
	}

	val = GetChildElem(val)
	err = checkExtractValid(val, opt.resetCheck())
	if err != nil {
		return
	}

	cloneAssigner := InitNew(oriAssigner)
	assigner = GetInitChildElem(cloneAssigner)
	err = tryAssign(assigner, val, opt)
	if err != nil {
		return
	}

	oriAssigner.Set(cloneAssigner)
	return
}

func tryAssign(assigner reflect.Value, val reflect.Value, opt *Option) (err error) {
	defer recoverFnOpt(&err, opt)

	var (
		assignerKind = GetKind(assigner)
		inSwitch     bool
	)
	switch assignerKind {
	case reflect.Bool:
		inSwitch = true
		var result bool
		result, err = extractBool(val, opt)
		if err != nil {
			return
		}
		assigner.SetBool(result)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		inSwitch = true
		switch GetType(assigner) {
		case TypeDuration:
			var resultDur time.Duration
			resultDur, err = extractDuration(val, opt)
			if err != nil {
				return
			}
			assigner.Set(reflect.ValueOf(resultDur))
		default:
			var result int64
			result, err = extractInt(val, opt)
			if err != nil {
				return
			}
			if assigner.OverflowInt(result) {
				err = getErrOverflow(assigner)
				return
			}
			assigner.SetInt(result)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		inSwitch = true
		var result uint64
		result, err = extractUint(val, opt)
		if err != nil {
			return
		}
		if assigner.OverflowUint(result) {
			err = getErrOverflow(assigner)
			return
		}
		assigner.SetUint(result)
	case reflect.Float32, reflect.Float64:
		inSwitch = true
		var result float64
		result, err = extractFloat(val, opt)
		if err != nil {
			return
		}
		if assigner.OverflowFloat(result) {
			err = getErrOverflow(assigner)
			return
		}
		assigner.SetFloat(result)
	case reflect.Complex64, reflect.Complex128:
		inSwitch = true
		var result complex128
		result, err = extractComplex(val, opt)
		if err != nil {
			return
		}
		if assigner.OverflowComplex(result) {
			err = getErrOverflow(assigner)
			return
		}
		assigner.SetComplex(result)
	case reflect.String:
		inSwitch = true
		var result string
		result, err = extractString(val, opt)
		if err != nil {
			return
		}
		assigner.SetString(result)
	case reflect.Array, reflect.Slice:
		inSwitch = true
		err = assignDefault(assigner, val)
		if err == nil {
			break
		}
		err = nil
		switch GetType(assigner) {
		case TypeIP:
			var ipVal net.IP
			ipVal, err = extractIP(val, opt)
			if err != nil {
				return
			}
			assigner.Set(reflect.ValueOf(ipVal))
		default:
			isSlice := assignerKind == reflect.Slice
			valKind := GetKind(val)
			switch valKind {
			case reflect.Array, reflect.Slice:
				err = getErrOverflowedLength(assigner, val, isSlice)
				if err != nil {
					return
				}

				err = iterateAndAssign(assigner, val, isSlice, opt)
			case reflect.String:
				err = iterateAndAssignString(assigner, val, isSlice, opt)
			default:
				err = getErrUnimplementedAssign(assigner, val)
			}
		}
	case reflect.Map:
		inSwitch = true
		err = assignDefault(assigner, val)
		if err == nil {
			break
		}
		err = nil
		valKind := GetKind(val)
		switch valKind {
		case reflect.Map, reflect.Struct:
			err = assignMap(assigner, val, opt)
		default:
			err = getErrUnimplementedAssign(assigner, val)
		}
	case reflect.Struct:
		inSwitch = true
		err = assignDefault(assigner, val)
		if err == nil {
			break
		}
		err = nil
		switch GetType(assigner) {
		case TypeTime:
			var timeRes time.Time
			timeRes, err = extractTime(val, opt)
			if err != nil {
				return
			}
			assigner.Set(reflect.ValueOf(timeRes))
		case TypeURL:
			var urlRes *url.URL
			urlRes, err = extractURL(val, opt)
			if err != nil {
				return
			}
			assigner.Set(GetElem(reflect.ValueOf(urlRes)))
		default:
			valKind := GetKind(val)
			switch valKind {
			case reflect.Map, reflect.Struct:
				err = assignMap(assigner, val, opt)
			default:
				err = getErrUnimplementedAssign(assigner, val)
			}
		}
	}
	if inSwitch {
		return
	}
	err = assignDefault(assigner, val)
	return
}

func assignMap(assigner reflect.Value, val reflect.Value, opt *Option) (err error) {
	err = getErrCanAddrInterface(assigner)
	if err != nil {
		return
	}

	opt.DecoderConfig.Result = assigner.Addr().Interface()
	decoder, err := mapstructure.NewDecoder(opt.DecoderConfig)
	if err != nil {
		return
	}

	err = decoder.Decode(val.Interface())
	return
}

func assignDefault(assigner reflect.Value, val reflect.Value) (err error) {
	err = getErrUnassignable(assigner, val)
	if err != nil {
		return
	}

	assigner.Set(val)
	return
}

func iterateAndAssign(assigner reflect.Value, val reflect.Value, isSlice bool, opt *Option) (err error) {
	tm := task.NewErrorManager(task.WithBufferSize(val.Len()))
	var emptyList reflect.Value
	if isSlice {
		emptyList = reflect.MakeSlice(GetType(assigner), val.Len(), val.Len())
	} else {
		typeArr := reflect.ArrayOf(assigner.Len(), GetElemType(assigner))
		emptyList = reflect.New(typeArr).Elem()
	}

	for index := 0; index < val.Len(); index++ {
		opt := opt.Clone()
		index := index
		tm.Run(func() (err error) {
			err = assignReflect(emptyList.Index(index), val.Index(index), opt)
			return
		})
	}
	err = tm.Error()
	if err != nil {
		return
	}

	assigner.Set(emptyList)
	return
}

func iterateAndAssignString(assigner reflect.Value, val reflect.Value, isSlice bool, opt *Option) (err error) {
	var listVal reflect.Value
	switch GetElemKind(assigner) {
	case reflect.Uint8:
		listVal = reflect.ValueOf([]byte(val.String()))
	case reflect.Int32:
		listVal = reflect.ValueOf([]rune(val.String()))
	default:
		err = getErrUnimplementedAssign(assigner, val)
		return
	}
	err = getErrOverflowedLength(assigner, listVal, isSlice)
	if err != nil {
		return
	}

	err = iterateAndAssign(assigner, listVal, isSlice, opt)
	return
}
