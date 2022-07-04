package reflecthelper

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Hello string
}

func TestCast(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name     string
		args     args
		wantKind reflect.Kind
	}{
		{
			name: "invalid nil value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantKind: reflect.Invalid,
		},
		{
			name: "valid slice value",
			args: args{
				val: reflect.ValueOf([]int{1, 2, 3}),
			},
			wantKind: reflect.Slice,
		},
		{
			name: "valid struct value",
			args: args{
				val: reflect.ValueOf(test{"Hi!"}),
			},
			wantKind: reflect.Struct,
		},
		{
			name: "valid ptr struct value",
			args: args{
				val: reflect.ValueOf(test{"Hi!"}),
			},
			wantKind: reflect.Struct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := Cast(tt.args.val)
			assert.Equal(t, tt.wantKind, GetKind(gotRes.Value))
		})
	}
}

func TestValue_IterateStruct(t *testing.T) {
	t.Run("no casting", func(t *testing.T) {
		var hello int
		val := Value{Value: reflect.ValueOf(hello)}
		assert.Nil(t, val.IterateStruct().Error())
	})
	t.Run("kind is not struct", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello), WithDecoderConfig(nil))
		assert.Nil(t, val.IterateStruct().Error())
	})
	t.Run("iterate example function", func(t *testing.T) {
		type test struct {
			Hello string
			Test  int
			One   uint64
		}

		val := Cast(reflect.ValueOf(test{"Hi!", 0, 10}))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			fmt.Println(field.Type(), field.Interface())
			return
		})
		assert.Nil(t, val.Error())

		// Concurrent Mode
		val = Cast(reflect.ValueOf(test{"Hi!", 0, 10}), WithConcurrency(true))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			fmt.Println(field.Type(), field.Interface())
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("error in the iteration", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(test{"Hi!"}))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			return errors.New("random error")
		})
		assert.NotNil(t, val.Error())
	})
	t.Run("error in the iteration ignored", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(test{"Hi!"}), WithIgnoreError(true))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			return errors.New("random error")
		})
		assert.Nil(t, val.Error())
	})
	t.Run("panic with recoverer - error type", func(t *testing.T) {
		type test struct {
			Hello string
		}
		errTest := errors.New("test")

		val := Cast(reflect.ValueOf(test{"Hi!"}), WithPanicRecoverer(true))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			panic(errTest)
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, errTest, val.Error())
	})
	t.Run("panic with recoverer - any type", func(t *testing.T) {
		type test struct {
			Hello string
		}
		val := Cast(reflect.ValueOf(test{"Hi!"}), WithPanicRecoverer(true))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			panic("hello")
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "hello", val.Error().Error())
	})
}

func TestValue_iterateArraySlice(t *testing.T) {
	t.Run("kind is not slice or array", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello))
		assert.Nil(t, val.IterateArraySlice().Error())
	})
	t.Run("iterate example function", func(t *testing.T) {
		var hello = []int{1, 2, 3, 4, 5}

		val := Cast(reflect.ValueOf(hello))
		val.IterateArraySlice(nil, func(parent reflect.Value, index int, elem reflect.Value) (err error) {
			fmt.Println("index: ", index, "elem: ", elem.Interface())
			return
		})
		assert.Nil(t, val.Error())

		var helloArray = [5]int{1, 2, 3, 4, 5}

		val = Cast(reflect.ValueOf(helloArray))
		val.IterateArraySlice(nil, func(parent reflect.Value, index int, elem reflect.Value) (err error) {
			fmt.Println("index: ", index, "elem: ", elem.Interface())
			return
		})
		assert.Nil(t, val.Error())

		val = Cast(reflect.ValueOf(helloArray), WithConcurrency(true))
		val.IterateArraySlice(nil, func(parent reflect.Value, index int, elem reflect.Value) (err error) {
			fmt.Println("index: ", index, "elem: ", elem.Interface())
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("error in the iteration", func(t *testing.T) {
		var hello = []int{1, 2, 3, 4, 5}

		val := Cast(reflect.ValueOf(hello))
		val.IterateArraySlice(nil, func(parent reflect.Value, index int, elem reflect.Value) (err error) {
			err = errors.New("any error")
			return
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "any error", val.Error().Error())
	})
	t.Run("error in the iteration ignored", func(t *testing.T) {
		var hello = [5]int{1, 2, 3, 4, 5}
		val := Cast(reflect.ValueOf(hello), WithIgnoreError(true))
		val.IterateArraySlice(nil, func(parent reflect.Value, index int, elem reflect.Value) (err error) {
			err = errors.New("any error")
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("panic with recoverer", func(t *testing.T) {
		var hello = [5]int{1, 2, 3, 4, 5}
		val := Cast(reflect.ValueOf(hello), WithPanicRecoverer(true))
		val.IterateArraySlice(nil, func(parent reflect.Value, index int, elem reflect.Value) (err error) {
			panic(errors.New("panic error"))
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "panic error", val.Error().Error())
	})
}

func TestValue_IterateMap(t *testing.T) {
	t.Run("kind is not map", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello))
		assert.Nil(t, val.IterateMap().Error())
	})
	t.Run("iterate example function", func(t *testing.T) {
		var test = map[string]string{
			"hello": "hi",
			"hi":    "hello",
			"test":  "testing",
			"try":   "trial",
		}

		val := Cast(reflect.ValueOf(test))
		val.IterateMap(nil, func(parent reflect.Value, key reflect.Value, elem reflect.Value) (err error) {
			fmt.Println("key: ", key, "val: ", elem.Interface())
			return
		})
		assert.Nil(t, val.Error())

		val = Cast(reflect.ValueOf(test), WithConcurrency(true))
		val.IterateMap(nil, func(parent reflect.Value, key reflect.Value, elem reflect.Value) (err error) {
			fmt.Println("key: ", key, "val: ", elem.Interface())
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("error in the iteration", func(t *testing.T) {
		var test = map[string]string{
			"hello": "hi",
			"hi":    "hello",
		}

		val := Cast(reflect.ValueOf(test))
		val.IterateMap(nil, func(parent reflect.Value, key reflect.Value, elem reflect.Value) (err error) {
			err = errors.New("random error")
			return
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "random error", val.Error().Error())

		// concurrent mode
		val = Cast(reflect.ValueOf(test), WithConcurrency(true))
		val.IterateMap(nil, func(parent reflect.Value, key reflect.Value, elem reflect.Value) (err error) {
			err = errors.New("random error")
			return
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "random error", val.Error().Error())
	})
	t.Run("error in the iteration ignored", func(t *testing.T) {
		var test = map[string]string{
			"hello": "hi",
			"hi":    "hello",
		}

		val := Cast(reflect.ValueOf(test), WithIgnoreError(true))
		val.IterateMap(nil, func(parent reflect.Value, key reflect.Value, elem reflect.Value) (err error) {
			err = errors.New("random error")
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("panic with recoverer", func(t *testing.T) {
		var test = map[string]string{
			"hello": "hi",
			"hi":    "hello",
		}

		val := Cast(reflect.ValueOf(test), WithPanicRecoverer(true))
		val.IterateMap(nil, func(parent reflect.Value, key reflect.Value, elem reflect.Value) (err error) {
			panic(errors.New("panic error"))
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "panic error", val.Error().Error())
	})
}

func TestValue_IterateChan(t *testing.T) {
	t.Run("kind is not map", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello))
		assert.Nil(t, val.IterateChan().Error())
	})
	t.Run("iterate example function", func(t *testing.T) {
		chanInt := make(chan int, 10)
		chanInt <- 1
		chanInt <- 2
		chanInt <- 3
		close(chanInt)

		// Iterate without blocking
		val := Cast(reflect.ValueOf(&chanInt))
		val.IterateChan(nil, func(chanInput, recv reflect.Value) (err error) {
			fmt.Println(recv.Interface())
			return
		})
		assert.Nil(t, val.Error())

		chanInt = make(chan int, 10)
		chanInt <- 1
		chanInt <- 2
		chanInt <- 3
		close(chanInt)

		// Iterate with blocking
		val = Cast(reflect.ValueOf(&chanInt), WithBlockChannel(true))
		val.IterateChan(nil, func(chanInput, recv reflect.Value) (err error) {
			fmt.Println(recv.Interface())
			return
		})
		assert.Nil(t, val.Error())

		// concurrent mode
		chanInt = make(chan int, 10)
		chanInt <- 1
		chanInt <- 2
		chanInt <- 3
		close(chanInt)

		// Iterate without blocking
		val = Cast(reflect.ValueOf(&chanInt), WithConcurrency(true))
		val.IterateChan(nil, func(chanInput, recv reflect.Value) (err error) {
			fmt.Println(recv.Kind(), recv.Interface())
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("error in the iteration", func(t *testing.T) {
		chanInt := make(chan int, 10)
		chanInt <- 1
		chanInt <- 2
		chanInt <- 3
		close(chanInt)

		val := Cast(reflect.ValueOf(&chanInt))
		val.IterateChan(nil, func(chanInput, recv reflect.Value) (err error) {
			err = errors.New("error in the channel")
			return
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "error in the channel", val.Error().Error())
	})
	t.Run("error in the iteration ignored", func(t *testing.T) {
		chanInt := make(chan int, 10)
		chanInt <- 1
		chanInt <- 2
		chanInt <- 3
		close(chanInt)

		val := Cast(reflect.ValueOf(&chanInt), WithIgnoreError(true))
		val.IterateChan(nil, func(chanInput, recv reflect.Value) (err error) {
			err = errors.New("error in the channel")
			return
		})
		assert.Nil(t, val.Error())
	})
	t.Run("panic with recoverer", func(t *testing.T) {
		chanInt := make(chan int, 10)
		chanInt <- 1
		chanInt <- 2
		chanInt <- 3
		close(chanInt)

		val := Cast(reflect.ValueOf(&chanInt), WithPanicRecoverer(true))
		val.IterateChan(nil, func(chanInput, recv reflect.Value) (err error) {
			panic(errors.New("error in the channel"))
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "error in the channel", val.Error().Error())
	})
}

func TestIsNil(t *testing.T) {
	var (
		interfaceVal boolInt
		checkVal     *checkBool
		checkSlice   []int
		checkFunc    func()
	)
	interfaceVal = checkVal
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil ptr inside interface",
			args: args{
				val: interfaceVal,
			},
			want: true,
		},
		{
			name: "nil interface",
			args: args{
				val: nil,
			},
			want: true,
		},
		{
			name: "nil slice",
			args: args{
				val: checkSlice,
			},
			want: true,
		},
		{
			name: "nil func",
			args: args{
				val: checkFunc,
			},
			want: true,
		},
		{
			name: "valid int",
			args: args{
				val: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.val); got != tt.want {
				t.Errorf("IsValueNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPtr(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args func() args
		want bool
	}{
		{
			name: "a pointer",
			args: func() args {
				var newVar int
				return args{&newVar}
			},
			want: true,
		},
		{
			name: "invalid null pointer",
			args: func() args {
				type test struct{}
				var newVar *test
				return args{&newVar}
			},
			want: true,
		},

		{
			name: "valid value",
			args: func() args {
				var newVar string
				return args{newVar}
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args()
			assert.Equalf(t, tt.want, IsPtr(args.in), "IsPtr(%v)", args.in)
		})
	}
}
