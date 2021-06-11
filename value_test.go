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
				val: reflect.ValueOf(&test{"Hi!"}),
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
	t.Run("kind is not struct", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello))
		val.IterateStruct()
	})
	t.Run("iterate example function", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(&test{"Hi!"}))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) {
			fmt.Println(val.String())
			fmt.Println(field.String())
		})
	})
}

func TestValue_IterateStructPanic(t *testing.T) {
	t.Run("kind is not struct", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello))
		val.IterateStructPanic()
	})
	t.Run("panic happens in the iteration", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(&test{"Hi!"}))
		val.IterateStructPanic(nil, func(val reflect.Value, field reflect.Value) {
			panic("random panic")
		})
		assert.NotNil(t, val.Error())
	})
}

func TestValue_IterateStructError(t *testing.T) {
	t.Run("kind is not struct", func(t *testing.T) {
		var hello int
		val := Cast(reflect.ValueOf(hello))
		val.IterateStructError()
	})
	t.Run("error in the second iteration", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(&test{"Hi!"}))
		val.IterateStructError(nil, func(val reflect.Value, field reflect.Value) (err error) {
			err = errors.New("random error")
			return
		})
		assert.NotNil(t, val.Error())
	})
	t.Run("success in the second iteration", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(&test{"Hi!"}))
		val.IterateStructError(nil, func(val reflect.Value, field reflect.Value) (err error) {
			fmt.Println("Success!!!")
			return
		})
		assert.Nil(t, val.Error())
	})
}
