package reflecthelper

import (
	"reflect"

	"github.com/fairyhunter13/task/v2"
)

type (

	// IterStructFn is a function type to iterate each field of structInput and returning an error if needed.
	IterStructFn func(structInput reflect.Value, field reflect.Value) error

	// IterArraySliceFn is a function type to iterate each field of array or slice and returning an error if needed.
	IterArraySliceFn func(arrSliceInput reflect.Value, index int, field reflect.Value) error

	// IterMapFn is a function type to iterate each key and element of map and returning an error if needed.
	IterMapFn func(mapInput reflect.Value, key reflect.Value, value reflect.Value) error

	// IterChanFn is a function type to iterate each value received by a channel and returning an error if needed.
	IterChanFn func(chanInput reflect.Value, recv reflect.Value) error
)

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

func (s *Value) isMap() bool {
	return IsKindMap(s.kind)
}

func (s *Value) isChan() bool {
	// CanSet is needed to check whether the reflect Value can fulfill the requirement of mustBeExported().
	return IsKindChan(s.kind) && s.CanSet()
}

func (s *Value) iterateEachStruct(fns []IterStructFn, index int) (err error) {
	for _, fn := range fns {
		if fn == nil {
			continue
		}
		err = fn(s.Value, s.Field(index))
		if s.opt.IgnoreError {
			err = nil
			continue
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *Value) iterateStruct(fns []IterStructFn) (err error) {
	numField := s.NumField()
	tm := task.NewErrorManager(task.WithBufferSize(numField))
	for index := 0; index < numField; index++ {
		index := index
		if s.opt.ConcurrentMode {
			tm.Run(func() (err error) {
				err = s.iterateEachStruct(fns, index)
				return
			})
			continue
		}
		err = s.iterateEachStruct(fns, index)
		if err != nil {
			return
		}
	}
	err = tm.Error()
	return
}

func (s *Value) iterateEachArraySlice(fns []IterArraySliceFn, index int) (err error) {
	for _, fn := range fns {
		if fn == nil {
			continue
		}
		err = fn(s.Value, index, s.Index(index))
		if s.opt.IgnoreError {
			err = nil
			continue
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *Value) iterateArraySlice(fns []IterArraySliceFn) (err error) {
	lenList := s.Len()
	tm := task.NewErrorManager(task.WithBufferSize(lenList))
	for index := 0; index < lenList; index++ {
		index := index
		if s.opt.ConcurrentMode {
			tm.Run(func() (err error) {
				err = s.iterateEachArraySlice(fns, index)
				return
			})
			continue
		}
		err = s.iterateEachArraySlice(fns, index)
		if err != nil {
			return
		}
	}
	err = tm.Error()
	return
}

func (s *Value) iterateMap(fns []IterMapFn) (err error) {
	// TODO: Add concurrent mode (with toggle)?
	iter := s.MapRange()
	for iter.Next() {
		for _, fn := range fns {
			if fn == nil {
				continue
			}
			err = fn(s.Value, iter.Key(), iter.Value())
			if s.opt.IgnoreError {
				err = nil
				continue
			}
			if err != nil {
				return
			}
		}
	}
	return
}

func (s *Value) iterateChan(fns []IterChanFn) (err error) {
	// TODO: Add concurrent mode (with toggle)?
	for {
		var (
			val reflect.Value
			ok  bool
		)
		if s.opt.BlockChannelIteration {
			val, ok = s.Recv()
		} else {
			val, ok = s.TryRecv()
		}
		if !ok {
			break
		}
		for _, fn := range fns {
			if fn == nil {
				continue
			}
			err = fn(s.Value, val)
			if s.opt.IgnoreError {
				err = nil
				continue
			}
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
	if s.opt == nil {
		s.opt = NewDefaultOption()
	}
	s.opt.Assign(fnOpts...)
	return s.init()
}

// IterateStruct iterates the struct field using the IterStructFn.
func (s *Value) IterateStruct(fns ...IterStructFn) *Value {
	if !s.init().isStruct() {
		return s
	}

	defer recoverFnOpt(&s.err, s.opt)
	s.err = s.iterateStruct(fns)
	return s
}

// IterateArraySlice iterates the element of slice or array using the IterArraySliceFn.
func (s *Value) IterateArraySlice(fns ...IterArraySliceFn) *Value {
	if !s.init().isArrayOrSlice() {
		return s
	}

	defer recoverFnOpt(&s.err, s.opt)
	s.err = s.iterateArraySlice(fns)
	return s
}

// IterateMap iterates the element of map using the IterMapFn.
func (s *Value) IterateMap(fns ...IterMapFn) *Value {
	if !s.init().isMap() {
		return s
	}

	defer recoverFnOpt(&s.err, s.opt)
	s.err = s.iterateMap(fns)
	return s
}

// IterateChan iterates the received elements using IterChanFn.
func (s *Value) IterateChan(fns ...IterChanFn) *Value {
	if !s.init().isChan() {
		return s
	}

	defer recoverFnOpt(&s.err, s.opt)
	s.err = s.iterateChan(fns)
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
