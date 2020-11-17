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
			wantTyp: nil,
		},
		{
			name: "valid string type",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantTyp: reflect.TypeOf("hello"),
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

func TestIsTypeValueElemable(t *testing.T) {
	number := 76
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil type not elemable",
			args: args{
				val: reflect.ValueOf(nil),
			},
			want: false,
		},
		{
			name: "value type not elemable",
			args: args{
				val: reflect.ValueOf(5),
			},
			want: false,
		},
		{
			name: "pointer type is elemable",
			args: args{
				val: reflect.ValueOf(&number),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTypeValueElemable(tt.args.val); got != tt.want {
				t.Errorf("IsTypeElemable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTypeElemable(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{
			name: "nil type",
			args: args{
				typ: nil,
			},
			wantRes: false,
		},
		{
			name: "slice type",
			args: args{
				typ: TypeByteSlice,
			},
			wantRes: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := IsTypeElemable(tt.args.typ); gotRes != tt.wantRes {
				t.Errorf("IsTypeElemable() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetChildElemType(t *testing.T) {
	var testInt *int
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantTyp reflect.Type
	}{
		{
			name: "nil type",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantTyp: nil,
		},
		{
			name: "normal int type",
			args: args{
				val: reflect.ValueOf(5),
			},
			wantTyp: reflect.TypeOf(10),
		},
		{
			name: "pointer int type",
			args: args{
				val: reflect.ValueOf(testInt),
			},
			wantTyp: reflect.ValueOf(testInt).Type().Elem(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTyp := GetChildElemType(tt.args.val); !reflect.DeepEqual(gotTyp, tt.wantTyp) {
				t.Errorf("GetChildElemType() = %v, want %v", gotTyp, tt.wantTyp)
			}
		})
	}
}

func TestGetChildElemPtrType(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantTyp reflect.Type
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantTyp: nil,
		},
		{
			name: "valid slice type",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantTyp: reflect.TypeOf([]int{}),
		},
		{
			name: "valid ptr type",
			args: args{
				val: func() reflect.Value {
					var x **int
					return reflect.ValueOf(x)
				},
			},
			wantTyp: reflect.TypeOf(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTyp := GetChildElemPtrType(tt.args.val()); !reflect.DeepEqual(gotTyp, tt.wantTyp) {
				t.Errorf("GetChildElemPtrType() = %v, want %v", gotTyp, tt.wantTyp)
			}
		})
	}
}

func TestGetChildElemValueType(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantTyp reflect.Type
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantTyp: nil,
		},
		{
			name: "valid slice type",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantTyp: reflect.TypeOf([]int{}),
		},
		{
			name: "valid ptr type",
			args: args{
				val: func() reflect.Value {
					var x **int
					return reflect.ValueOf(x)
				},
			},
			wantTyp: reflect.TypeOf(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTyp := GetChildElemValueType(tt.args.val()); !reflect.DeepEqual(gotTyp, tt.wantTyp) {
				t.Errorf("GetChildElemValueType() = %v, want %v", gotTyp, tt.wantTyp)
			}
		})
	}
}
