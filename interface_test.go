package reflecthelper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnwrapInterfaceValue(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name     string
		args     args
		wantKind reflect.Kind
	}{
		{
			name: "invalid val",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantKind: reflect.Invalid,
		},
		{
			name: "valid interface value",
			args: args{
				val: func() reflect.Value {
					test := []interface{}{5}
					testVal := reflect.ValueOf(test)
					return testVal.Index(0)
				},
			},
			wantKind: reflect.Int,
		},
		{
			name: "valid multi level interface value",
			args: args{
				val: func() reflect.Value {
					test := []interface{}{interface{}(5)}
					testVal := reflect.ValueOf(test)
					return testVal.Index(0)
				},
			},
			wantKind: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := UnwrapInterfaceValue(tt.args.val())
			assert.Equal(t, gotRes.Kind(), tt.wantKind)
		})
	}
}

func TestIsNil(t *testing.T) {
	var (
		interfaceVal boolInt
		checkVal     *checkBool
		checkSlice   []int
		checkFunc    func()
	)
	interfaceVal = checkVal
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil ptr inside interface",
			args: args{
				val: interfaceVal,
			},
			want: true,
		},
		{
			name: "nil interface",
			args: args{
				val: nil,
			},
			want: true,
		},
		{
			name: "nil slice",
			args: args{
				val: checkSlice,
			},
			want: true,
		},
		{
			name: "nil func",
			args: args{
				val: checkFunc,
			},
			want: true,
		},
		{
			name: "valid int",
			args: args{
				val: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.val); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPtr(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args func() args
		want bool
	}{
		{
			name: "a pointer",
			args: func() args {
				var newVar int
				return args{&newVar}
			},
			want: true,
		},
		{
			name: "invalid null pointer",
			args: func() args {
				type test struct{}
				var newVar *test
				return args{&newVar}
			},
			want: true,
		},

		{
			name: "valid value",
			args: func() args {
				var newVar string
				return args{newVar}
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args()
			assert.Equalf(t, tt.want, IsPtr(args.in), "IsPtr(%v)", args.in)
		})
	}
}

func TestGetTypeInterface(t *testing.T) {
	var val int

	assert.Equal(t, reflect.TypeOf(5), GetTypeInterface(val))
	assert.Equal(t, reflect.TypeOf(10), GetTypeInterface(reflect.ValueOf(val)))
}

func TestGetKindInterface(t *testing.T) {
	var val int

	assert.Equal(t, reflect.TypeOf(5).Kind(), GetKindInterface(val))
	assert.Equal(t, reflect.TypeOf(10).Kind(), GetKindInterface(reflect.ValueOf(val)))
}
