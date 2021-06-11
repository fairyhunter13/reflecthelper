package reflecthelper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Hello string
}

func TestCastStruct(t *testing.T) {
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
			name: "invalid slice value",
			args: args{
				val: reflect.ValueOf([]int{1, 2, 3}),
			},
			wantKind: reflect.Invalid,
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
		val.IterateStruct(func(val reflect.Value, field reflect.Value) {
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
		// TODO: Add test in here
	})
}
