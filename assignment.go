package reflecthelper

import (
	"reflect"
	"time"
)

func checkAssigner(assigner reflect.Value) (err error) {
	err = getErrIsValid(assigner)
	if err != nil {
		return
	}
	if !assigner.CanSet() {
		err = ErrAssignerCantSet
	}
	return
}

func assignReflect(assigner reflect.Value, val reflect.Value, opt *Option) (err error) {
	assigner = GetChildElem(assigner)
	val = GetChildElem(val)
	err = checkAssigner(assigner)
	if err != nil {
		return
	}
	opt.ResetCheck()
	err = checkExtractValid(val, opt)
	if err != nil {
		return
	}
	err = tryAssign(assigner, val, opt)
	return
}

func tryAssign(assigner reflect.Value, val reflect.Value, opt *Option) (err error) {
	defer RecoverFn(&err)

	assignerKind := GetKind(assigner)
	valKind := GetKind(val)
	switch assignerKind {
	case reflect.Bool:
		var result bool
		result, err = extractBool(val, opt)
		if err != nil {
			return
		}
		assigner.SetBool(result)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
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
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
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
	case reflect.Array, reflect.Slice:
		isSlice := assignerKind == reflect.Slice
		switch valKind {
		case reflect.Array, reflect.Slice:
			if !isSlice {
				err = checkOverLength(assigner, val)
				if err != nil {
					return
				}
			}
			err = iterateAndAssign(assigner, val, isSlice)
		case reflect.String:
			err = iterateAndAssignString(assigner, val, isSlice)
		default:
			err = getErrUnimplementedAssign(assigner, val)
		}
	case reflect.String:
		var result string
		result, err = extractString(val, opt)
		if err != nil {
			return
		}
		assigner.SetString(result)
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Struct:
		assignerType := assigner.Type()
		valType := val.Type()
		switch assignerType {
		case TypeTime:
			var timeStr string
			switch valKind {
			case reflect.String:
				timeStr = val.String()
			default:
				if valType == TypeTime {
					assigner.Set(val)
					return
				}
				timeStr, err = extractString(val, opt)
				if err != nil {
					return
				}
			}

			var timeStruct time.Time
			timeStruct, err = ParseTime(timeStr)
			if err != nil {
				return
			}
			assigner.Set(reflect.ValueOf(timeStruct))
		default:
			if !valType.AssignableTo(assignerType) {
				err = getErrUnassignable(assigner, val)
				return
			}
			assigner.Set(val)
		}
	default:
		err = getErrUnimplementedAssign(assigner, val)
	}
	return
}

func checkOverLength(assigner reflect.Value, val reflect.Value) (err error) {
	if assigner.Len() < val.Len() {
		err = getErrOverflowedLength(assigner, val)
	}
	return
}

func iterateAndAssign(assigner reflect.Value, val reflect.Value, isSlice bool) (err error) {
	if isSlice {
		emptySlice := reflect.MakeSlice(assigner.Type(), 0, val.Len())
		for index := 0; index < val.Len(); index++ {
			elemVal := GetInitChildElem(reflect.New(GetElemType(assigner)).Elem())
			err = AssignReflect(elemVal, val.Index(index))
			if err != nil {
				return
			}
			emptySlice = reflect.Append(emptySlice, elemVal)
		}
		assigner.Set(emptySlice)
	} else {
		typeArr := reflect.ArrayOf(assigner.Len(), GetElemType(assigner))
		emptyArray := reflect.New(typeArr).Elem()
		for index := 0; index < val.Len(); index++ {
			elemVal := GetInitChildElem(emptyArray.Index(index))
			err = AssignReflect(elemVal, val.Index(index))
			if err != nil {
				return
			}
		}
		assigner.Set(emptyArray)
	}
	return
}

func iterateAndAssignString(assigner reflect.Value, val reflect.Value, isSlice bool) (err error) {
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
	if !isSlice {
		err = checkOverLength(assigner, listVal)
		if err != nil {
			return
		}
	}
	err = iterateAndAssign(assigner, listVal, isSlice)
	return
}
