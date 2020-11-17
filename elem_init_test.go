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

func TestGetInitChildPtrElem(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes func() reflect.Value
	}{
		{
			name: "valid string value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf("hello")
			},
		},
		{
			name: "valid int value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(5)
				},
			},
			wantRes: func() reflect.Value {
				return reflect.ValueOf(5)
			},
		},
		{
			name: "valid multi level ptr unsettable int value",
			args: args{
				val: func() reflect.Value {
					var test **int
					return reflect.ValueOf(test)
				},
			},
			wantRes: func() reflect.Value {
				var test **int
				return reflect.ValueOf(test)
			},
		},
		{
			name: "valid ptr int value",
			args: args{
				val: func() reflect.Value {
					var test *int
					return reflect.ValueOf(&test)
				},
			},
			wantRes: func() reflect.Value {
				var test *int
				return reflect.ValueOf(test)
			},
		},
		{
			name: "valid multi level ptr int value",
			args: args{
				val: func() reflect.Value {
					var test ****int
					return reflect.ValueOf(&test)
				},
			},
			wantRes: func() reflect.Value {
				var test *int
				return reflect.ValueOf(test)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := GetInitChildPtrElem(tt.args.val())
			if !(gotRes.Type() == tt.wantRes().Type()) {
				t.Errorf("GetInitChildPtrElem() = %v, want %v", gotRes.Type(), tt.wantRes().Type())
			}
		})
	}
}
