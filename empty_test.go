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
	var elemInt *int
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
		reflect.ValueOf(nil),
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
		reflect.ValueOf([1]*int{elemInt}),
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
