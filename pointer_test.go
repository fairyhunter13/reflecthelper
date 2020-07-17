package reflecthelper

import (
	"reflect"
	"testing"
)

func TestGetInitElem(t *testing.T) {
	var strPtr1 *string
	var strPtr2 **string
	type test struct {
		Hello string
	}
	type testPointer struct {
		Hello *string
	}
	var check test
	var checkPointer testPointer
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Value
	}{
		{
			name: "struct pointer nil field type",
			args: args{
				v: reflect.ValueOf(&checkPointer.Hello),
			},
			wantRes: reflect.ValueOf(checkPointer.Hello),
		},
		{
			name: "struct field pointer type",
			args: args{
				v: reflect.ValueOf(checkPointer.Hello),
			},
			wantRes: reflect.ValueOf(checkPointer.Hello),
		},
		{
			name: "struct field type",
			args: args{
				v: reflect.ValueOf(check.Hello),
			},
			wantRes: reflect.ValueOf(check.Hello),
		},
		{
			name: "Pointer second level set",
			args: args{
				v: reflect.ValueOf(&strPtr2),
			},
			wantRes: reflect.ValueOf(strPtr2),
		},
		{
			name: "Pointer second level",
			args: args{
				v: reflect.ValueOf(strPtr2),
			},
			wantRes: reflect.ValueOf(strPtr2),
		},
		{
			name: "Pointer one level set",
			args: args{
				v: reflect.ValueOf(&strPtr1),
			},
			wantRes: reflect.ValueOf(strPtr1),
		},
		{
			name: "Pointer one level",
			args: args{
				v: reflect.ValueOf(strPtr1),
			},
			wantRes: reflect.ValueOf(strPtr1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetInitElem(tt.args.v); !(gotRes.Type() == tt.wantRes.Type()) {
				t.Errorf("GetInitElem() = %v, want %v", gotRes.Type(), tt.wantRes.Type())
			}
		})
	}
}

func TestGetInitChildElem(t *testing.T) {
	var strPtr1 *string
	var strPtr2 **string
	var intPtr2 **int
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Value
	}{
		{
			name: "Get original type if can't set for 2nd level pointer",
			args: args{
				v: reflect.ValueOf(strPtr2),
			},
			wantRes: reflect.ValueOf(strPtr2),
		},
		{
			name: "Get original type if can't set",
			args: args{
				v: reflect.ValueOf(strPtr1),
			},
			wantRes: reflect.ValueOf(strPtr1),
		},
		{
			name: "Get simple child elem of one pointer level",
			args: args{
				v: reflect.ValueOf(&strPtr1),
			},
			wantRes: reflect.ValueOf(""),
		},
		{
			name: "Get simple child elem of two pointer levels",
			args: args{
				v: reflect.ValueOf(&strPtr2),
			},
			wantRes: reflect.ValueOf(""),
		},
		{
			name: "Get simple child elem of two pointer levels of int type",
			args: args{
				v: reflect.ValueOf(&intPtr2),
			},
			wantRes: reflect.ValueOf(int(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetInitChildElem(tt.args.v); !(gotRes.Type() == tt.wantRes.Type()) {
				t.Errorf("GetInitChildElem() = %v, want %v", gotRes.Type(), tt.wantRes.Type())
			}
		})
	}
}

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
