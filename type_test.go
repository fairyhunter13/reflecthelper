package reflecthelper

import (
	"reflect"
	"testing"
)

func TestGetElemType(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantTyp reflect.Type
	}{
		{
			name: "invalid type",
			args: args{
				val: reflect.ValueOf(nil),
			},
		},
		{
			name: "invalid elem type",
			args: args{
				val: reflect.ValueOf("hello"),
			},
		},
		{
			name: "array of uint8",
			args: args{
				val: reflect.ValueOf([]uint8("hello")),
			},
			wantTyp: reflect.TypeOf(uint8(1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTyp := GetElemType(tt.args.val); !reflect.DeepEqual(gotTyp, tt.wantTyp) {
				t.Errorf("GetElemType() = %v, want %v", gotTyp, tt.wantTyp)
			}
		})
	}
}
