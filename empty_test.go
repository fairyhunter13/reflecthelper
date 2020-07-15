package reflecthelper

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MyInt int
type ZeroStruct struct{}

func TestZero(t *testing.T) {
	var zeroValues = []interface{}{
		false,
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		complex64(0),
		complex128(0),
		float32(0),
		float64(0),
		"",
		MyInt(0),
		reflect.ValueOf(0),
		reflect.ValueOf(""),
		nil,
		time.Time{},
		&time.Time{},
		nilTime,
		ZeroStruct{},
		&ZeroStruct{},
	}

	for _, v := range zeroValues {
		t.Run(fmt.Sprintf("%#v", v), func(t *testing.T) {
			assert.True(t, IsZero(v))
		})
	}
}

func TestIsValueZero(t *testing.T) {
	var test chan int
	type complexStruct struct {
		Hello *string
		Test  string
	}
	type cascadeStruct struct {
		Hello *string
		Test  struct {
			Success *bool
			Failed  *bool
		}
	}
	var zeroReflectValues = []reflect.Value{
		reflect.ValueOf(false),
		reflect.ValueOf(int(0)),
		reflect.ValueOf(int8(0)),
		reflect.ValueOf(int16(0)),
		reflect.ValueOf(int32(0)),
		reflect.ValueOf(int64(0)),
		reflect.ValueOf(uint(0)),
		reflect.ValueOf(uint8(0)),
		reflect.ValueOf(uint16(0)),
		reflect.ValueOf(uint32(0)),
		reflect.ValueOf(uint64(0)),
		reflect.ValueOf(uintptr(0)),
		reflect.ValueOf(float32(0)),
		reflect.ValueOf(float64(0)),
		reflect.ValueOf(complex64(0)),
		reflect.ValueOf(complex128(0)),
		reflect.ValueOf(""),
		reflect.ValueOf([0]int{}),
		reflect.ValueOf([5]int{}),
		reflect.ValueOf(complexStruct{}),
		reflect.ValueOf(&complexStruct{}),
		reflect.ValueOf(cascadeStruct{}),
		reflect.ValueOf(&cascadeStruct{}),
		reflect.ValueOf(test),
		reflect.ValueOf(MyInt(0)),
		reflect.ValueOf(time.Time{}),
		reflect.ValueOf(&time.Time{}),
		reflect.ValueOf(nilTime),
		reflect.ValueOf(ZeroStruct{}),
		reflect.ValueOf(&ZeroStruct{}),
	}

	for _, v := range zeroReflectValues {
		t.Run(fmt.Sprintf("%#v", v), func(t *testing.T) {
			assert.True(t, IsValueZero(v))
		})
	}
}

func TestIsValueNotZero(t *testing.T) {
	now := time.Now()
	nowPtr := &now
	success := true
	man := "man"
	test := make(chan int, 1)
	type complexStruct struct {
		Hello *string
		Test  string
	}
	type Test struct {
		Success *bool
		Failed  *bool
	}
	type cascadeStruct struct {
		Hello *string
		Test  Test
	}
	var zeroReflectValues = []reflect.Value{
		reflect.ValueOf(true),
		reflect.ValueOf(int(1)),
		reflect.ValueOf(int8(1)),
		reflect.ValueOf(int16(1)),
		reflect.ValueOf(int32(1)),
		reflect.ValueOf(int64(1)),
		reflect.ValueOf(uint(1)),
		reflect.ValueOf(uint8(1)),
		reflect.ValueOf(uint16(1)),
		reflect.ValueOf(uint32(1)),
		reflect.ValueOf(uint64(1)),
		reflect.ValueOf(uintptr(1)),
		reflect.ValueOf(float32(1)),
		reflect.ValueOf(float64(1)),
		reflect.ValueOf(complex64(1)),
		reflect.ValueOf(complex128(1)),
		reflect.ValueOf("1"),
		reflect.ValueOf([1]string{"1"}),
		reflect.ValueOf(complexStruct{Hello: &man}),
		reflect.ValueOf(&complexStruct{Hello: &man}),
		reflect.ValueOf(cascadeStruct{Test: Test{Success: &success}}),
		reflect.ValueOf(&cascadeStruct{Test: Test{Success: &success}}),
		reflect.ValueOf(test),
		reflect.ValueOf(MyInt(1)),
		reflect.ValueOf(now),
		reflect.ValueOf(&now),
		reflect.ValueOf(nowPtr),
	}

	for _, v := range zeroReflectValues {
		t.Run(fmt.Sprintf("%#v", v), func(t *testing.T) {
			assert.False(t, IsValueZero(v))
		})
	}
}

type structTest struct{}

func (st structTest) Hello() {}
func TestIsReflectZero(t *testing.T) {
	type test interface {
		Hello()
	}
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
