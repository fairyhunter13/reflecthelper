package reflecthelper

import (
	"reflect"
	"testing"
)

type structTest struct{}

func (st structTest) Hello() {}
func TestIsReflectZero(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Get true for nil implementation of interface",
			args: args{
				val: reflect.ValueOf((*structTest)(nil)),
			},
			wantResult: true,
		},
		{
			name: "Invalid reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantResult: false,
		},
		{
			name: "Valid reflect value",
			args: args{
				val: reflect.ValueOf(0),
			},
			wantResult: false,
		},
		{
			name: "Valid reflect value non empty",
			args: args{
				val: reflect.ValueOf(5),
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := IsReflectZero(tt.args.val); gotResult != tt.wantResult {
				t.Errorf("IsReflectZero() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestIsInterfaceReflectZero(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "Nil value",
			args: args{
				val: nil,
			},
			wantResult: true,
		},
		{
			name: "Empty value",
			args: args{
				val: 0,
			},
			wantResult: false,
		},
		{
			name: "Non empty value",
			args: args{
				val: 5,
			},
			wantResult: false,
		},
		{
			name: "Nil implementation of struct",
			args: args{
				val: (*structTest)(nil),
			},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := IsInterfaceReflectZero(tt.args.val); gotResult != tt.wantResult {
				t.Errorf("IsInterfaceReflectZero() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSetReflectZero(t *testing.T) {
	var test *int
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
		},
		{
			name: "nil pointer of pointer integer value",
			args: args{
				val: reflect.ValueOf(&test),
			},
		},
		{
			name: "elem of nil pointer integer value",
			args: args{
				val: reflect.ValueOf(&test).Elem(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetReflectZero(tt.args.val)
		})
	}
}
