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
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			fmt.Println(val.String())
			fmt.Println(field.String())
			return
		})
	})
	t.Run("error in the iteration", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(&test{"Hi!"}))
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			return errors.New("random error")
		})
		assert.NotNil(t, val.Error())
	})
	t.Run("error in the iteration ignored", func(t *testing.T) {
		type test struct {
			Hello string
		}

		val := Cast(reflect.ValueOf(&test{"Hi!"}), EnableIgnoreError())
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

		val := Cast(reflect.ValueOf(&test{"Hi!"}), EnablePanicRecoverer())
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
		val := Cast(reflect.ValueOf(&test{"Hi!"}), EnablePanicRecoverer())
		val.IterateStruct(nil, func(val reflect.Value, field reflect.Value) (err error) {
			panic("hello")
		})
		assert.NotNil(t, val.Error())
		assert.Equal(t, "hello", val.Error().Error())
	})
}
