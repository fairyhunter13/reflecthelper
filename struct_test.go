package reflecthelper

import (
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
		wantErr  bool
	}{
		{
			name: "invalid nil value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantKind: reflect.Invalid,
			wantErr:  true,
		},
		{
			name: "invalid slice value",
			args: args{
				val: reflect.ValueOf([]int{1, 2, 3}),
			},
			wantKind: reflect.Invalid,
			wantErr:  true,
		},
		{
			name: "valid struct value",
			args: args{
				val: reflect.ValueOf(test{"Hi!"}),
			},
			wantKind: reflect.Struct,
			wantErr:  false,
		},
		{
			name: "valid ptr struct value",
			args: args{
				val: reflect.ValueOf(&test{"Hi!"}),
			},
			wantKind: reflect.Struct,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := CastStruct(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("CastStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantKind, GetKind(gotRes.Value))
		})
	}
}
