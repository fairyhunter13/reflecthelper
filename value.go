package reflecthelper

import (
	"reflect"
)

// IterStructFn is function type to iterate each field of structInput.
type IterStructFn func(structInput reflect.Value, field reflect.Value)

// IterStructErrFn is function type to iterate each field of structInput and returning error if needed.
type IterStructErrFn func(structInput reflect.Value, field reflect.Value) error

// Value is a custom struct for representing reflect.Value.
type Value struct {
	reflect.Value
	kind reflect.Kind
	err  error
}

// Error returns the error contained within the Value.
func (s *Value) Error() error {
	return s.err
}

func (s *Value) isStruct() bool {
	return IsKindStruct(s.kind)
}

func (s *Value) iterateStruct(fns ...IterStructFn) {
	for index := 0; index < s.NumField(); index++ {
		for _, fn := range fns {
			if fn == nil {
				continue
			}
			fn(s.Value, s.Field(index))
		}
	}
}

func (s *Value) iterateStructError(fns ...IterStructErrFn) (err error) {
	for index := 0; index < s.NumField(); index++ {
		for _, fn := range fns {
			if fn == nil {
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

// IterateStruct iterates all the field in this struct.
func (s *Value) IterateStruct(fns ...IterStructFn) *Value {
	if !s.isStruct() {
		return s
	}

	s.iterateStruct(fns...)
	return s
}

// IterateStructPanic iterates all the field in this struct by returning err when the iteration function panics.
func (s *Value) IterateStructPanic(fns ...IterStructFn) *Value {
	if !s.isStruct() {
		return s
	}

	func() {
		defer RecoverFn(&s.err)
		s.iterateStruct(fns...)
	}()
	return s
}

func (s *Value) IterateStructError(fns ...IterStructErrFn) *Value {
	if !s.isStruct() {
		return s
	}

	s.err = s.iterateStructError(fns...)
	return s
}

// Cast casts the val of reflect.Value to the Value of this package.
func Cast(val reflect.Value) (res Value) {
	val = GetChildElem(val)
	res = Value{
		Value: val,
		kind:  GetKind(val),
	}
	return
}
