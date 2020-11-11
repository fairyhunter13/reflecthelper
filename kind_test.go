package reflecthelper

import (
	"reflect"
	"testing"
)

func TestGetKind(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid kind",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "normal kind",
			args: args{
				val: reflect.ValueOf(int(5)),
			},
			wantRes: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetKind(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetKind() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetElemKind(t *testing.T) {
	var valInt *int
	test := 5
	valInt = &test
	testSlice := []interface{}{
		0, valInt, nil,
	}
	valSlice := reflect.ValueOf(testSlice)

	var valNilPtr **int
	type args struct {
		val reflect.Value
	}
	interfaceOfString := interface{}("hello")
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid kind",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "normal kind",
			args: args{
				val: reflect.ValueOf(int(5)),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "normal pointer kind",
			args: args{
				val: reflect.ValueOf(valInt),
			},
			wantRes: reflect.Int,
		},
		{
			name: "nil pointer kind",
			args: args{
				val: reflect.ValueOf(valNilPtr),
			},
			wantRes: reflect.Ptr,
		},
		{
			name: "array of int type",
			args: args{
				val: reflect.ValueOf([3]int{}),
			},
			wantRes: reflect.Int,
		},
		{
			name: "pointer array of int type",
			args: args{
				val: reflect.ValueOf(&[3]int{}),
			},
			wantRes: reflect.Array,
		},
		{
			name: "slice of int type",
			args: args{
				val: reflect.ValueOf([]string{}),
			},
			wantRes: reflect.String,
		},
		{
			name: "pointer slice of int type",
			args: args{
				val: reflect.ValueOf(&[]string{}),
			},
			wantRes: reflect.Slice,
		},
		{
			name: "interface of interface type",
			args: args{
				val: reflect.ValueOf(interface{}(interfaceOfString)),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "elem of slice interface",
			args: args{
				val: valSlice.Index(0),
			},
			wantRes: reflect.Int,
		},
		{
			name: "elem ptr of slice interface",
			args: args{
				val: valSlice.Index(1),
			},
			wantRes: reflect.Int,
		},
		{
			name: "elem nil of slice interface",
			args: args{
				val: valSlice.Index(2),
			},
			wantRes: reflect.Invalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetElemKind(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetElemKind() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetChildElemKind(t *testing.T) {
	var testIntPtr **int
	test := 5
	testPtr := &test
	testIntPtr = &testPtr
	testSlice := []interface{}{
		0, testIntPtr, nil,
	}
	valSlice := reflect.ValueOf(testSlice)
	var testNilPtr *int
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid kind for child elem kind",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "invalid kind for non child elem kind",
			args: args{
				val: reflect.ValueOf(5),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "test multiple pointer",
			args: args{
				val: reflect.ValueOf(testIntPtr),
			},
			wantRes: reflect.Int,
		},
		{
			name: "test nil ptr",
			args: args{
				val: reflect.ValueOf(testNilPtr),
			},
			wantRes: reflect.Int,
		},
		{
			name: "elem of slice interface",
			args: args{
				val: valSlice.Index(0),
			},
			wantRes: reflect.Int,
		},
		{
			name: "elem ptr of slice interface",
			args: args{
				val: valSlice.Index(1),
			},
			wantRes: reflect.Int,
		},
		{
			name: "elem nil of slice interface",
			args: args{
				val: valSlice.Index(2),
			},
			wantRes: reflect.Invalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetChildElemKind(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetChildElemKind() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
