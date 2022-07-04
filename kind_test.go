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
			wantRes: reflect.Int,
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
			wantRes: reflect.String,
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
				val: valSlice.Index(1).Elem(),
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
			if gotRes := GetKindElem(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetKindElem() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetChildElemTypeKind(t *testing.T) {
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
			name: "valid kind for int",
			args: args{
				val: reflect.ValueOf(5),
			},
			wantRes: reflect.Int,
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
			if gotRes := GetKindChildElemType(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetKindChildElemType() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetChildElemPtrKind(t *testing.T) {
	var k **int
	var testSlice []int
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid reflect kind",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "valid slice kind",
			args: args{
				val: reflect.ValueOf([]int{}),
			},
			wantRes: reflect.Slice,
		},
		{
			name: "valid ptr slice kind",
			args: args{
				val: reflect.ValueOf(&testSlice),
			},
			wantRes: reflect.Slice,
		},
		{
			name: "valid int ptr kind",
			args: args{
				val: reflect.ValueOf(k),
			},
			wantRes: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetKindChildElemPtr(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetKindChildElemPtr() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetChildElemValueKind(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "valid interface value",
			args: args{
				val: func() reflect.Value {
					var hello interface{} = 5
					return reflect.ValueOf(hello)
				},
			},
			wantRes: reflect.Int,
		},
		{
			name: "valid ptr value",
			args: args{
				val: func() reflect.Value {
					test := 5
					hello := &test
					return reflect.ValueOf(&hello)
				},
			},
			wantRes: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetKindChildElemValue(tt.args.val()); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetKindChildElemValue() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestIsKindBool(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid bool kind",
			args: args{
				kind: reflect.Bool,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindBool(tt.args.kind); got != tt.want {
				t.Errorf("IsKindBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindValueBytesSlice(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			want: false,
		},
		{
			name: "invalid array int value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([5]int{})
				},
			},
			want: false,
		},
		{
			name: "invalid slice int value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]int{})
				},
			},
			want: false,
		},
		{
			name: "valid slice bytes value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]byte{})
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindValueBytesSlice(tt.args.val()); got != tt.want {
				t.Errorf("IsKindValueBytesSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindSlice(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid chan kind",
			args: args{
				kind: reflect.Chan,
			},
			want: false,
		},
		{
			name: "valid slice kind",
			args: args{
				kind: reflect.Slice,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindSlice(tt.args.kind); got != tt.want {
				t.Errorf("IsKindSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindComplex(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid complex64 kind",
			args: args{
				kind: reflect.Complex64,
			},
			want: true,
		},
		{
			name: "valid complex128 kind",
			args: args{
				kind: reflect.Complex128,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindComplex(tt.args.kind); got != tt.want {
				t.Errorf("IsKindComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindFloat(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid float32 kind",
			args: args{
				kind: reflect.Float32,
			},
			want: true,
		},
		{
			name: "valid float64 kind",
			args: args{
				kind: reflect.Float64,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindFloat(tt.args.kind); got != tt.want {
				t.Errorf("IsKindFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindInt(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid int kind",
			args: args{
				kind: reflect.Int,
			},
			want: true,
		},
		{
			name: "valid int8 kind",
			args: args{
				kind: reflect.Int8,
			},
			want: true,
		},
		{
			name: "valid int16 kind",
			args: args{
				kind: reflect.Int16,
			},
			want: true,
		},
		{
			name: "valid int32 kind",
			args: args{
				kind: reflect.Int32,
			},
			want: true,
		},
		{
			name: "valid int64 kind",
			args: args{
				kind: reflect.Int64,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindInt(tt.args.kind); got != tt.want {
				t.Errorf("IsKindInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindUnsafePointer(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid unsafe ptr kind",
			args: args{
				kind: reflect.UnsafePointer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindUnsafePointer(tt.args.kind); got != tt.want {
				t.Errorf("IsKindUnsafePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindString(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid string kind",
			args: args{
				kind: reflect.String,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindString(tt.args.kind); got != tt.want {
				t.Errorf("IsKindString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindUint(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid kind",
			args: args{
				kind: reflect.Invalid,
			},
			want: false,
		},
		{
			name: "valid uint kind",
			args: args{
				kind: reflect.Uint,
			},
			want: true,
		},
		{
			name: "valid uint8 kind",
			args: args{
				kind: reflect.Uint8,
			},
			want: true,
		},
		{
			name: "valid uint16 kind",
			args: args{
				kind: reflect.Uint16,
			},
			want: true,
		},
		{
			name: "valid uint32 kind",
			args: args{
				kind: reflect.Uint32,
			},
			want: true,
		},
		{
			name: "valid uint64 kind",
			args: args{
				kind: reflect.Uint64,
			},
			want: true,
		},
		{
			name: "valid uint ptr kind",
			args: args{
				kind: reflect.Uintptr,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindUint(tt.args.kind); got != tt.want {
				t.Errorf("IsKindUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKindInterface(t *testing.T) {
	type args struct {
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "kind interface",
			args: args{
				kind: reflect.Interface,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKindInterface(tt.args.kind); got != tt.want {
				t.Errorf("IsKindInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
