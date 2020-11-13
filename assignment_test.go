package reflecthelper

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignReflect(t *testing.T) {
	type args struct {
		assigner func() reflect.Value
		val      func() reflect.Value
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		wantAssigner reflect.Value
	}{
		{
			name: "invalid assigner and val",
			args: args{
				assigner: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantErr: true,
		},
		{
			name: "assigner can't set",
			args: args{
				assigner: func() reflect.Value {
					return reflect.ValueOf(5)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantErr: true,
		},
		{
			name: "invalid val",
			args: args{
				assigner: func() reflect.Value {
					hello := 5
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantErr: true,
		},
		{
			name: "invalid bool",
			args: args{
				assigner: func() reflect.Value {
					hello := true
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantErr: true,
		},
		{
			name: "valid bool",
			args: args{
				assigner: func() reflect.Value {
					hello := true
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("false")
				},
			},
			wantErr:      false,
			wantAssigner: reflect.ValueOf(false),
		},
		{
			name: "invalid int",
			args: args{
				assigner: func() reflect.Value {
					hello := int(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantErr: true,
		},
		{
			name: "overflow int8",
			args: args{
				assigner: func() reflect.Value {
					hello := int8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(math.MaxInt8 + 1)
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assigner := tt.args.assigner()
			if err := AssignReflect(assigner, tt.args.val()); (err != nil) != tt.wantErr {
				t.Errorf("AssignReflect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantAssigner.IsValid() {
				assert.EqualValues(t, tt.wantAssigner.Interface(), GetChildElem(assigner).Interface())
			}
		})
	}
}
