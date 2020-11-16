package reflecthelper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetChildElem(t *testing.T) {
	var test **string
	var k *string
	var y *int
	x := 5
	y = &x
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Value
	}{
		{
			name: "Get child of two pointer level",
			args: args{
				val: reflect.ValueOf(test),
			},
			wantRes: reflect.ValueOf(test),
		},
		{
			name: "Get child of two pointer level with reference",
			args: args{
				val: reflect.ValueOf(&test),
			},
			wantRes: reflect.ValueOf(test),
		},
		{
			name: "Get child of one pointer level",
			args: args{
				val: reflect.ValueOf(k),
			},
			wantRes: reflect.ValueOf(k),
		},
		{
			name: "Get child of one pointer level with reference",
			args: args{
				val: reflect.ValueOf(&k),
			},
			wantRes: reflect.ValueOf(k),
		},
		{
			name: "Get child of initialized one pointer level",
			args: args{
				val: reflect.ValueOf(y),
			},
			wantRes: reflect.ValueOf(int(0)),
		},
		{
			name: "Get child of initialized one pointer level with reference",
			args: args{
				val: reflect.ValueOf(&y),
			},
			wantRes: reflect.ValueOf(int(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetChildElem(tt.args.val); !(gotRes.Type() == tt.wantRes.Type()) {
				t.Errorf("GetChildElem() = %v, want %v", gotRes.Type(), tt.wantRes.Type())
			}
		})
	}
}

func TestGetElem(t *testing.T) {
	var k *string
	var y **string
	var test *int
	x := 5
	test = &x
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Value
	}{
		{
			name: "Get child with one pointer level with reference",
			args: args{
				val: reflect.ValueOf(&k),
			},
			wantRes: reflect.ValueOf(k),
		},
		{
			name: "Get child with one pointer level",
			args: args{
				val: reflect.ValueOf(k),
			},
			wantRes: reflect.ValueOf(k),
		},
		{
			name: "Get child with two pointer level with reference",
			args: args{
				val: reflect.ValueOf(&y),
			},
			wantRes: reflect.ValueOf(y),
		},
		{
			name: "Get child with two pointer level",
			args: args{
				val: reflect.ValueOf(y),
			},
			wantRes: reflect.ValueOf(y),
		},
		{
			name: "Get child initialized with reference",
			args: args{
				val: reflect.ValueOf(&test),
			},
			wantRes: reflect.ValueOf(test),
		},
		{
			name: "Get child initialized",
			args: args{
				val: reflect.ValueOf(test),
			},
			wantRes: reflect.ValueOf(int(0)),
		},
		{
			name: "non pointer passing",
			args: args{
				val: reflect.ValueOf(int(0)),
			},
			wantRes: reflect.ValueOf(int(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetElem(tt.args.val); !(gotRes.Type() == tt.wantRes.Type()) {
				t.Errorf("GetElem() = %v, want %v", gotRes.Type(), tt.wantRes.Type())
			}
		})
	}
}

func TestIsValueElemable(t *testing.T) {
	number := 1
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil value is not elemable",
			args: args{
				val: reflect.ValueOf(nil),
			},
			want: false,
		},
		{
			name: "value is not elemable",
			args: args{
				val: reflect.ValueOf(1),
			},
			want: false,
		},
		{
			name: "value is elemable",
			args: args{
				val: reflect.ValueOf(&number),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValueElemable(tt.args.val); got != tt.want {
				t.Errorf("IsValueElemable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNilElem(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes func() reflect.Value
	}{
		{
			name: "slice not elemable",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf([]int{1, 2, 3})
			},
		},
		{
			name: "valid nil ptr",
			args: args{
				val: func() reflect.Value {
					var hello *string
					return reflect.ValueOf(hello)
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf(nil)
			},
		},
		{
			name: "valid ptr string",
			args: args{
				val: func() reflect.Value {
					test := "hello"
					hello := &test
					return reflect.ValueOf(hello)
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf("hi")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetNilElem(tt.args.val()); !(gotRes.Kind() == tt.wantRes().Kind()) {
				fmt.Println(gotRes)
				t.Errorf("GetNilElem() = %v, want %v", gotRes, tt.wantRes())
			}
		})
	}
}

func TestGetChildNilElem(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes func() reflect.Value
	}{
		{
			name: "slice not elemable",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf([]int{1, 2, 3})
			},
		},
		{
			name: "valid nil ptr",
			args: args{
				val: func() reflect.Value {
					var hello *string
					return reflect.ValueOf(hello)
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf(nil)
			},
		},
		{
			name: "valid multi ptr string",
			args: args{
				val: func() reflect.Value {
					test := "hello"
					hello := &test
					hi := &hello
					return reflect.ValueOf(&hi)
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf("hello")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetChildNilElem(tt.args.val()); !(gotRes.Kind() == tt.wantRes().Kind()) {
				t.Errorf("GetChildNilElem() = %v, want %v", gotRes, tt.wantRes())
			}
		})
	}
}
