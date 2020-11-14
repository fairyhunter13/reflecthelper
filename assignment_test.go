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
		wantAssigner func() reflect.Value
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
			wantErr:      true,
			wantAssigner: nil,
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
			wantErr:      true,
			wantAssigner: nil,
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
			wantErr:      true,
			wantAssigner: nil,
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
			wantErr:      true,
			wantAssigner: nil,
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
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(false)
			},
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
			wantErr:      true,
			wantAssigner: nil,
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
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "set int8 succeeded",
			args: args{
				assigner: func() reflect.Value {
					hello := int8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(10)
				},
			},
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(10)
			},
		},
		{
			name: "invalid uint value",
			args: args{
				assigner: func() reflect.Value {
					hello := uint(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "overflow uint8 value",
			args: args{
				assigner: func() reflect.Value {
					hello := uint8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(math.MaxUint8 + 1)
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid uint8 value",
			args: args{
				assigner: func() reflect.Value {
					hello := uint8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("10")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(uint8(10))
			},
			wantErr: false,
		},
		{
			name: "invalid float value",
			args: args{
				assigner: func() reflect.Value {
					hello := float32(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "overflow float32 value",
			args: args{
				assigner: func() reflect.Value {
					hello := float32(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(math.MaxFloat64)
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid float32 value",
			args: args{
				assigner: func() reflect.Value {
					hello := float32(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(10)
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(float32(10))
			},
			wantErr: false,
		},
		{
			name: "invalid complex64 value",
			args: args{
				assigner: func() reflect.Value {
					hello := complex64(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "overflow complex64 value",
			args: args{
				assigner: func() reflect.Value {
					hello := complex64(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(complex(math.MaxFloat64, 0))
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid complex64 value",
			args: args{
				assigner: func() reflect.Value {
					hello := complex64(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(complex64(10))
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(complex64(10))
			},
			wantErr: false,
		},
		{
			name: "unimplemented val kind",
			args: args{
				assigner: func() reflect.Value {
					hello := []int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(make(chan int, 0))
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		// TODO: Add more tests for slice and array
		{
			name: "invalid string value",
			args: args{
				assigner: func() reflect.Value {
					hello := "check"
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{0, 1, 2, 3})
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid string value",
			args: args{
				assigner: func() reflect.Value {
					hello := "check"
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf("hello")
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assigner := tt.args.assigner()
			if err := AssignReflect(assigner, tt.args.val()); (err != nil) != tt.wantErr {
				t.Errorf("AssignReflect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantAssigner != nil && tt.wantAssigner().IsValid() {
				assert.EqualValues(t, tt.wantAssigner().Interface(), GetChildElem(assigner).Interface())
			}
		})
	}
}
