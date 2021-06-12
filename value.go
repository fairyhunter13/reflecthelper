package reflecthelper

import (
	"reflect"
)

// IterStructErrFn is function type to iterate each field of structInput and returning error if needed.
type IterStructErrFn func(structInput reflect.Value, field reflect.Value) error

// IterArrSliceErrFn is function type to iterate each field of array or slice and returning error if needed.
type IterArrSliceErrFn func(parent reflect.Value, index int, field reflect.Value) error

// Value is a custom struct for representing reflect.Value.
type Value struct {
	reflect.Value
	kind reflect.Kind
	err  error
	opt  *Option
}

func (s *Value) init() *Value {
	if s.opt == nil {
		s.opt = NewDefaultOption()
	}
	return s
}

func (s *Value) isStruct() bool {
	return IsKindStruct(s.kind)
}

func (s *Value) isArrayOrSlice() bool {
	return IsKindSlice(s.kind) || IsKindArray(s.kind)
}

func (s *Value) iterateStructError(fns ...IterStructErrFn) (err error) {
	// TODO: Add concurrent mode (with toggle)?
	for index := 0; index < s.NumField(); index++ {
		for _, fn := range fns {
			if fn == nil {
				continue
			}
			if s.opt.IgnoreError {
				fn(s.Value, s.Field(index))
				continue
			}
			err = fn(s.Value, s.Field(index))
			if err != nil {
				return
			}
		}
	}
	return
}

// Error returns the error contained within the Value.
func (s *Value) Error() error {
	return s.err
}

// Assign assigns the function options to the s.opt.
func (s *Value) Assign(fnOpts ...FuncOption) *Value {
	s.init().opt.Assign(fnOpts...)
	return s
}

// IterateStruct iterates the struct field using the IterStructErrFn.
func (s *Value) IterateStruct(fns ...IterStructErrFn) *Value {
	if !s.init().isStruct() {
		return s
	}

	defer recoverFnOpt(&s.err, s.opt)
	s.err = s.iterateStructError(fns...)
	return s
}

func (s *Value) IterateArrayOrSlice(fns ...IterArrSliceErrFn) *Value {
	if !s.init().isArrayOrSlice() {
		return s
	}

	defer recoverFnOpt(&s.err, s.opt)
	// TODO: Add logic in here
	return s
}

// Cast casts the val of reflect.Value to the Value of this package.
func Cast(val reflect.Value, fnOpts ...FuncOption) (res Value) {
	val = GetChildElem(val)
	res = *(&Value{
		Value: val,
		kind:  GetKind(val),
	}).Assign(fnOpts...)
	return
}
