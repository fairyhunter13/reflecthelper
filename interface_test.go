package reflecthelper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnwrapInterfaceValue(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name     string
		args     args
		wantKind reflect.Kind
	}{
		{
			name: "invalid val",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantKind: reflect.Invalid,
		},
		{
			name: "valid interface value",
			args: args{
				val: func() reflect.Value {
					test := []interface{}{5}
					testVal := reflect.ValueOf(test)
					return testVal.Index(0)
				},
			},
			wantKind: reflect.Int,
		},
		{
			name: "valid multi level interface value",
			args: args{
				val: func() reflect.Value {
					test := []interface{}{interface{}(5)}
					testVal := reflect.ValueOf(test)
					return testVal.Index(0)
				},
			},
			wantKind: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := UnwrapInterfaceValue(tt.args.val())
			assert.Equal(t, gotRes.Kind(), tt.wantKind)
		})
	}
}
