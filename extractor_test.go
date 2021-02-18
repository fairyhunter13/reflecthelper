package reflecthelper

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type booler struct {
	val bool
}

func (b *booler) Bool() (bool, error) {
	return b.val, nil
}

type boolerVal struct {
	val bool
}

func (b boolerVal) Bool() (bool, error) {
	return b.val, nil
}

func TestExtractBool(t *testing.T) {
	type test struct {
		hello string
	}
	fieldVal := reflect.ValueOf(test{"hello"}).FieldByIndex([]int{0})
	var nilBool *bool
	valBool := true
	valBooler := booler{true}
	ptrValBooler := &valBooler
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
		wantErr    bool
	}{
		{
			name: "invalid interface value",
			args: args{
				val: fieldVal,
			},
			wantErr: true,
		},
		{
			name: "invalid reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "invalid bool string",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantErr: true,
		},
		{
			name: "nil bool pointer",
			args: args{
				val: reflect.ValueOf(nilBool),
			},
			wantErr: true,
		},
		{
			name: "bool value",
			args: args{
				val: reflect.ValueOf(true),
			},
			wantResult: true,
		},
		{
			name: "bool pointer value",
			args: args{
				val: reflect.ValueOf(&valBool),
			},
			wantResult: true,
		},
		{
			name: "bool string value",
			args: args{
				val: reflect.ValueOf("true"),
			},
			wantResult: true,
		},
		{
			name: "elem ptr booler value",
			args: args{
				val: reflect.ValueOf(ptrValBooler).Elem(),
			},
			wantResult: true,
		},
		{
			name: "ptr booler value",
			args: args{
				val: reflect.ValueOf(ptrValBooler),
			},
			wantResult: true,
		},
		{
			name: "boolerVal value",
			args: args{
				val: reflect.ValueOf(boolerVal{true}),
			},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractBool(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractBool() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

type int64er struct {
	val int64
}

func (i *int64er) Int64() (int64, error) {
	return i.val, nil
}

type int64Val struct {
	val int64
}

func (i int64Val) Int64() (int64, error) {
	return i.val, nil
}

type anon struct {
	hello string
}

func TestExtractInt(t *testing.T) {
	valInt := 10
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
		wantErr    bool
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "invalid struct value",
			args: args{
				val: reflect.ValueOf(anon{hello: "test"}),
			},
			wantErr: true,
		},
		{
			name: "bool true value",
			args: args{
				val: reflect.ValueOf(true),
			},
			wantResult: 1,
		},
		{
			name: "bool false value",
			args: args{
				val: reflect.ValueOf(false),
			},
			wantResult: 0,
		},
		{
			name: "int value",
			args: args{
				val: reflect.ValueOf(5),
			},
			wantResult: 5,
		},
		{
			name: "uint8 value",
			args: args{
				val: reflect.ValueOf(uint8(5)),
			},
			wantResult: 5,
		},
		{
			name: "uint8 value",
			args: args{
				val: reflect.ValueOf(uint8(5)),
			},
			wantResult: 5,
		},
		{
			name: "int ptr value",
			args: args{
				val: reflect.ValueOf(&valInt),
			},
			wantResult: 10,
		},
		{
			name: "string value",
			args: args{
				val: reflect.ValueOf("20"),
			},
			wantResult: 20,
		},
		{
			name: "ptr int64er value",
			args: args{
				val: reflect.ValueOf(&int64er{15}),
			},
			wantResult: 15,
		},
		{
			name: "elem ptr int64er value",
			args: args{
				val: reflect.ValueOf(&int64er{15}).Elem(),
			},
			wantResult: 15,
		},
		{
			name: "int64Val value",
			args: args{
				val: reflect.ValueOf(int64Val{15}),
			},
			wantResult: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractInt(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractInt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

type uinter struct {
	val uint
}

func (u *uinter) Uint() (uint, error) {
	return u.val, nil
}

type uinterVal struct {
	val uint
}

func (u uinterVal) Uint() (uint, error) {
	return u.val, nil
}

func TestExtractUint(t *testing.T) {
	uintPtr := uint(15)
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint64
		wantErr    bool
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "overflow int value",
			args: args{
				val: reflect.ValueOf(-100),
			},
			wantErr: true,
		},
		{
			name: "invalid string value",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantErr: true,
		},
		{
			name: "invalid struct value",
			args: args{
				val: reflect.ValueOf(anon{hello: "test"}),
			},
			wantErr: true,
		},
		{
			name: "bool value true",
			args: args{
				val: reflect.ValueOf(true),
			},
			wantResult: 1,
		},
		{
			name: "bool value false",
			args: args{
				val: reflect.ValueOf(false),
			},
			wantResult: 0,
		},
		{
			name: "int value",
			args: args{
				val: reflect.ValueOf(int(5)),
			},
			wantResult: 5,
		},
		{
			name: "uint value",
			args: args{
				val: reflect.ValueOf(uint(10)),
			},
			wantResult: 10,
		},
		{
			name: "uint ptr value",
			args: args{
				val: reflect.ValueOf(&uintPtr),
			},
			wantResult: 15,
		},
		{
			name: "ptr uinter value",
			args: args{
				val: reflect.ValueOf(&uinter{15}),
			},
			wantResult: 15,
		},
		{
			name: "elem ptr uinter value",
			args: args{
				val: reflect.ValueOf(&uinter{15}).Elem(),
			},
			wantResult: 15,
		},
		{
			name: "uinterVal value",
			args: args{
				val: reflect.ValueOf(uinterVal{15}),
			},
			wantResult: 15,
		},
		{
			name: "string value",
			args: args{
				val: reflect.ValueOf("20"),
			},
			wantResult: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractUint(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractUint() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

type floater struct {
	val float64
}

func (f *floater) Float64() (float64, error) {
	return f.val, nil
}

type floaterVal struct {
	val float64
}

func (f floaterVal) Float64() (float64, error) {
	return f.val, nil
}

func TestExtractFloat(t *testing.T) {
	floatPtr := 13.0
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
		wantErr    bool
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "invalid string value",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantErr: true,
		},
		{
			name: "invalid struct value",
			args: args{
				val: reflect.ValueOf(anon{hello: "test"}),
			},
			wantErr: true,
		},
		{
			name: "bool true value",
			args: args{
				val: reflect.ValueOf(true),
			},
			wantResult: 1.0,
		},
		{
			name: "bool false value",
			args: args{
				val: reflect.ValueOf(false),
			},
			wantResult: 0.0,
		},
		{
			name: "int8 value",
			args: args{
				val: reflect.ValueOf(int8(-5)),
			},
			wantResult: -5.0,
		},
		{
			name: "uint8 value",
			args: args{
				val: reflect.ValueOf(uint8(5)),
			},
			wantResult: 5.0,
		},
		{
			name: "float value",
			args: args{
				val: reflect.ValueOf(10.0),
			},
			wantResult: 10.0,
		},
		{
			name: "float ptr value",
			args: args{
				val: reflect.ValueOf(&floatPtr),
			},
			wantResult: 13.0,
		},
		{
			name: "ptr floater value",
			args: args{
				val: reflect.ValueOf(&floater{11}),
			},
			wantResult: 11.0,
		},
		{
			name: "elem ptr floater value",
			args: args{
				val: reflect.ValueOf(&floater{11}).Elem(),
			},
			wantResult: 11.0,
		},
		{
			name: "floaterVal value",
			args: args{
				val: reflect.ValueOf(floaterVal{11}),
			},
			wantResult: 11.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractFloat(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractFloat() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

type complex128er struct {
	val complex128
}

func (c *complex128er) Complex128() (complex128, error) {
	return c.val, nil
}

type complex128Val struct {
	val complex128
}

func (c complex128Val) Complex128() (complex128, error) {
	return c.val, nil
}

func TestExtractComplex(t *testing.T) {
	testComplex := complex(15, 0)
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult complex128
		wantErr    bool
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "invalid string value",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantErr: true,
		},
		{
			name: "invalid struct value",
			args: args{
				val: reflect.ValueOf(anon{"yeah"}),
			},
			wantErr: true,
		},
		{
			name: "int8 value",
			args: args{
				val: reflect.ValueOf(int8(5)),
			},
			wantResult: complex(5, 0),
		},
		{
			name: "uint8 value",
			args: args{
				val: reflect.ValueOf(uint8(10)),
			},
			wantResult: complex(10, 0),
		},
		{
			name: "float64 value",
			args: args{
				val: reflect.ValueOf(float64(13)),
			},
			wantResult: complex(13, 0),
		},
		{
			name: "complex value",
			args: args{
				val: reflect.ValueOf(complex(15, 0)),
			},
			wantResult: complex(15, 0),
		},
		{
			name: "ptr complex value",
			args: args{
				val: reflect.ValueOf(&testComplex),
			},
			wantResult: complex(15, 0),
		},
		{
			name: "ptr complex128er value",
			args: args{
				val: reflect.ValueOf(&complex128er{complex(17, 0)}),
			},
			wantResult: complex(17, 0),
		},
		{
			name: "elem ptr complex128er value",
			args: args{
				val: reflect.ValueOf(&complex128er{complex(17, 0)}).Elem(),
			},
			wantResult: complex(17, 0),
		},
		{
			name: "complex128Val value",
			args: args{
				val: reflect.ValueOf(complex128Val{complex(17, 0)}),
			},
			wantResult: complex(17, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractComplex(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractComplex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractComplex() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

type stringer struct {
	val string
}

func (s *stringer) String() string {
	return s.val
}

type stringVal struct {
	val string
}

func (s stringVal) String() string {
	return s.val
}

func TestExtractString(t *testing.T) {
	intVal := 5
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			name: "invalid reflect value for input",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "unsupported reflect value for string conversion",
			args: args{
				val: reflect.ValueOf([]int{0, 1, 2, 3}),
			},
			wantErr: true,
		},
		{
			name: "bool to string",
			args: args{
				val: reflect.ValueOf(true),
			},
			wantResult: "true",
			wantErr:    false,
		},
		{
			name: "int to string",
			args: args{
				val: reflect.ValueOf(-5),
			},
			wantResult: "-5",
			wantErr:    false,
		},
		{
			name: "uint to string",
			args: args{
				val: reflect.ValueOf(uint(5)),
			},
			wantResult: "5",
			wantErr:    false,
		},
		{
			name: "float to string",
			args: args{
				val: reflect.ValueOf(0.005),
			},
			wantResult: "0.005",
			wantErr:    false,
		},
		{
			name: "int pointer to string",
			args: args{
				val: reflect.ValueOf(&intVal),
			},
			wantResult: "5",
			wantErr:    false,
		},
		{
			name: "string to string",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantResult: "hello",
			wantErr:    false,
		},
		{
			name: "byte slice to string",
			args: args{
				val: reflect.ValueOf([]byte("hello")),
			},
			wantResult: "hello",
			wantErr:    false,
		},
		{
			name: "ptr stringer value",
			args: args{
				val: reflect.ValueOf(&stringer{"hello"}),
			},
			wantResult: "hello",
			wantErr:    false,
		},
		{
			name: "elem ptr stringer value",
			args: args{
				val: reflect.ValueOf(&stringer{"hello"}).Elem(),
			},
			wantResult: "hello",
			wantErr:    false,
		},
		{
			name: "stringVal value",
			args: args{
				val: reflect.ValueOf(stringVal{"hello"}),
			},
			wantResult: "hello",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractString(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

type hello struct{ test string }

func (h *hello) String() string {
	return h.test
}

func TestExtractTime(t *testing.T) {
	type args struct {
		val func() reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes func() *time.Time
		wantErr bool
	}{
		{
			name: "invalid reflect value",
			args: args{
				val: func() reflect.Value { return reflect.ValueOf(nil) },
			},
			wantRes: func() *time.Time {
				return nil
			},
			wantErr: true,
		},
		{
			name: "invalid string value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf("test")
				},
			},
			wantRes: func() *time.Time {
				return nil
			},
			wantErr: true,
		},
		{
			name: "invalid slice value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf([]int{0, 1, 2, 3})
				},
			},
			wantRes: func() *time.Time {
				return nil
			},
			wantErr: true,
		},
		{
			name: "ptr time type value",
			args: args{
				val: func() reflect.Value {
					timeVal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
					return reflect.ValueOf(&timeVal)
				},
			},
			wantRes: func() *time.Time {
				timeVal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
				return &timeVal
			},
			wantErr: false,
		},
		{
			name: "time type value",
			args: args{
				val: func() reflect.Value {
					timeVal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
					return reflect.ValueOf(timeVal)
				},
			},
			wantRes: func() *time.Time {
				timeVal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
				return &timeVal
			},
			wantErr: false,
		},
		{
			name: "hello type value",
			args: args{
				val: func() reflect.Value {
					hi := &hello{test: "2006-01-02T15:04:05+07:00"}
					return reflect.ValueOf(hi)
				},
			},
			wantRes: func() *time.Time {
				timeVal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
				return &timeVal
			},
			wantErr: false,
		},
		{
			name: "string value",
			args: args{
				val: func() reflect.Value {
					return reflect.ValueOf("2006-01-02T15:04:05+07:00")
				},
			},
			wantRes: func() *time.Time {
				timeVal, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
				return &timeVal
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ExtractTime(tt.args.val())
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.wantRes(), gotRes)
		})
	}
}

func TestTryExtract(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult interface{}
		wantErr    bool
	}{
		{
			name: "unimplemented value",
			args: args{
				val: reflect.ValueOf([]int{0, 1, 2, 3}),
			},
			wantErr: true,
		},
		{
			name: "bool value",
			args: args{
				val: reflect.ValueOf(true),
			},
			wantResult: true,
		},
		{
			name: "int value",
			args: args{
				val: reflect.ValueOf(int(-5)),
			},
			wantResult: int64(-5),
		},
		{
			name: "uint value",
			args: args{
				val: reflect.ValueOf(uint(5)),
			},
			wantResult: uint64(5),
		},
		{
			name: "float value",
			args: args{
				val: reflect.ValueOf(3.15),
			},
			wantResult: 3.15,
		},
		{
			name: "complex value",
			args: args{
				val: reflect.ValueOf(complex(1, 0)),
			},
			wantResult: complex(1, 0),
		},
		{
			name: "string value",
			args: args{
				val: reflect.ValueOf("hello"),
			},
			wantResult: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := TryExtract(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryExtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantResult, gotResult)
		})
	}
}
