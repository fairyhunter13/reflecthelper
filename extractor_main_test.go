package reflecthelper

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{
			name: "true bool",
			args: args{
				val: reflect.ValueOf("true"),
			},
			wantRes: true,
		},
		{
			name: "false bool",
			args: args{
				val: reflect.ValueOf("0"),
			},
			wantRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetBool(tt.args.val, tt.args.fnOpts...); gotRes != tt.wantRes {
				t.Errorf("GetBool() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
	}{
		{
			name: "int64 value: 10",
			args: args{
				val: reflect.ValueOf(uint64(10)),
			},
			wantResult: 10,
		},
		{
			name: "int64 value: 0",
			args: args{
				val: reflect.ValueOf("0"),
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetInt(tt.args.val, tt.args.fnOpts...); gotResult != tt.wantResult {
				t.Errorf("GetInt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetUint(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint64
	}{
		{
			name: "uint64 value: 10",
			args: args{
				val: reflect.ValueOf(uint64(10)),
			},
			wantResult: 10,
		},
		{
			name: "uint64 value: 0",
			args: args{
				val: reflect.ValueOf("0"),
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetUint(tt.args.val, tt.args.fnOpts...); gotResult != tt.wantResult {
				t.Errorf("GetUint() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetFloat(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{
			name: "float64 value: 10",
			args: args{
				val: reflect.ValueOf(int64(10)),
			},
			wantResult: 10.0,
		},
		{
			name: "float64 value: 0",
			args: args{
				val: reflect.ValueOf("0"),
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetFloat(tt.args.val, tt.args.fnOpts...); gotResult != tt.wantResult {
				t.Errorf("GetFloat() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetComplex(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult complex128
	}{
		{
			name: "complex value: (10, 0)",
			args: args{
				val: reflect.ValueOf(int64(10)),
			},
			wantResult: complex(10, 0),
		},
		{
			name: "float64 value: 0",
			args: args{
				val: reflect.ValueOf("0"),
			},
			wantResult: complex(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetComplex(tt.args.val, tt.args.fnOpts...); gotResult != tt.wantResult {
				t.Errorf("GetComplex() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: `string value: "hello"`,
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantResult: `hello`,
		},
		{
			name: "string value: 27",
			args: args{
				val: reflect.ValueOf(float64(27)),
			},
			wantResult: "27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetString(tt.args.val, tt.args.fnOpts...); gotResult != tt.wantResult {
				t.Errorf("GetString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: `invalid value: "hello"`,
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantResult: false,
		},
		{
			name: "valid value: 2012-11-01T22:08:41+00:00",
			args: args{
				val: reflect.ValueOf("2012-11-01T22:08:41+00:00"),
			},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := GetTime(tt.args.val, tt.args.fnOpts...)
			if tt.wantResult {
				assert.NotEmpty(t, gotResult)
			} else {
				assert.Empty(t, gotResult)
			}
		})
	}
}

func TestTryGet(t *testing.T) {
	type args struct {
		val    reflect.Value
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult interface{}
	}{
		{
			name: `valid value: "hello"`,
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantResult: "hello",
		},
		{
			name: "valid value: int",
			args: args{
				val: reflect.ValueOf(5),
			},
			wantResult: int64(5),
		},
		{
			name: "valid value: false",
			args: args{
				val: reflect.ValueOf(false),
			},
			wantResult: false,
		},
		{
			name: "valid value: float64",
			args: args{
				val: reflect.ValueOf(float64(5.0212134)),
			},
			wantResult: 5.0212134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := TryGet(tt.args.val, tt.args.fnOpts...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TryGet() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	type args struct {
		val    interface{}
		fnOpts []FuncOption
	}
	tests := []struct {
		name       string
		args       args
		wantResult time.Duration
	}{
		{
			name: "valid duration syntax",
			args: args{
				val: reflect.ValueOf("5s"),
			},
			wantResult: 5 * time.Second,
		},
		{
			name: "invalid duration syntax",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantResult: 0,
		},
		{
			name: "valid duration 1 second",
			args: args{
				val: time.Second,
			},
			wantResult: time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetDuration(tt.args.val, tt.args.fnOpts...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetDuration() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
