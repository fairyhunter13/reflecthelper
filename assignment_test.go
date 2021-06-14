package reflecthelper

import (
	"math"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestAssignReflect(t *testing.T) {
	now := time.Now()
	type args struct {
		assigner func() reflect.Value
		val      func() reflect.Value
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		wantAssigner func() reflect.Value
	}{
		{
			name: "invalid assigner and val",
			args: args{
				assigner: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "assigner can't set",
			args: args{
				assigner: func() reflect.Value {
					return reflect.ValueOf(5)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "nil ptr assigner",
			args: args{
				assigner: func() reflect.Value {
					var hello *int
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(5)
				},
			},
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(5)
			},
		},
		{
			name: "interface assigner",
			args: args{
				assigner: func() reflect.Value {
					hello := []interface{}{5, 4, 3}
					assigner := reflect.ValueOf(&hello).Elem().Index(0)
					return assigner
				},
				val: func() reflect.Value {
					return reflect.ValueOf(10)
				},
			},
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(10)
			},
		},
		{
			name: "ptr interface assigner",
			args: args{
				assigner: func() reflect.Value {
					test := 5
					hello := []interface{}{&test, 4, 3}
					assigner := reflect.ValueOf(&hello).Elem().Index(0)
					return assigner
				},
				val: func() reflect.Value {
					return reflect.ValueOf(10)
				},
			},
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(10)
			},
		},
		{
			name: "invalid val",
			args: args{
				assigner: func() reflect.Value {
					hello := 5
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(nil)
				},
			},
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "invalid bool",
			args: args{
				assigner: func() reflect.Value {
					hello := true
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "valid bool",
			args: args{
				assigner: func() reflect.Value {
					hello := true
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("false")
				},
			},
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(false)
			},
		},
		{
			name: "invalid int",
			args: args{
				assigner: func() reflect.Value {
					hello := int(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "overflow int8",
			args: args{
				assigner: func() reflect.Value {
					hello := int8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(math.MaxInt8 + 1)
				},
			},
			wantErr:      true,
			wantAssigner: nil,
		},
		{
			name: "set int8 succeeded",
			args: args{
				assigner: func() reflect.Value {
					hello := int8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(10)
				},
			},
			wantErr: false,
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(10)
			},
		},
		{
			name: "invalid uint value",
			args: args{
				assigner: func() reflect.Value {
					hello := uint(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "overflow uint8 value",
			args: args{
				assigner: func() reflect.Value {
					hello := uint8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(math.MaxUint8 + 1)
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid uint8 value",
			args: args{
				assigner: func() reflect.Value {
					hello := uint8(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("10")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(uint8(10))
			},
			wantErr: false,
		},
		{
			name: "invalid float value",
			args: args{
				assigner: func() reflect.Value {
					hello := float32(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "overflow float32 value",
			args: args{
				assigner: func() reflect.Value {
					hello := float32(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(math.MaxFloat64)
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid float32 value",
			args: args{
				assigner: func() reflect.Value {
					hello := float32(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(10)
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(float32(10))
			},
			wantErr: false,
		},
		{
			name: "invalid complex64 value",
			args: args{
				assigner: func() reflect.Value {
					hello := complex64(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "overflow complex64 value",
			args: args{
				assigner: func() reflect.Value {
					hello := complex64(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(complex(math.MaxFloat64, 0))
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid complex64 value",
			args: args{
				assigner: func() reflect.Value {
					hello := complex64(5)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(complex64(10))
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(complex64(10))
			},
			wantErr: false,
		},
		{
			name: "unimplemented val kind",
			args: args{
				assigner: func() reflect.Value {
					hello := []int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(make(chan int))
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "over length for the assigner",
			args: args{
				assigner: func() reflect.Value {
					hello := [1]int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "not supported multi level array",
			args: args{
				assigner: func() reflect.Value {
					hello := [3][3]int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "not supported multi level slice",
			args: args{
				assigner: func() reflect.Value {
					hello := [][]int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid array assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := [3]int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf([3]int{1, 2, 3})
			},
			wantErr: false,
		},
		{
			name: "valid slice assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := []int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([3]int{1, 2, 3})
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf([]int{1, 2, 3})
			},
			wantErr: false,
		},
		{
			name: "valid slice ptr int assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := []*int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([3]int{1, 2, 3})
				},
			},
			wantAssigner: func() reflect.Value {
				var (
					num1 = int(1)
					num2 = int(2)
					num3 = int(3)
				)
				return reflect.ValueOf([]*int{&num1, &num2, &num3})
			},
			wantErr: false,
		},
		{
			name: "slice int unimplemented string assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := []int{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "slice uint8 overlength string assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := [2]uint8{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "slice uint8 valid string assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := []uint8{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf([]uint8("hello"))
			},
			wantErr: false,
		},
		{
			name: "array uint8 valid string assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := [5]uint8{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf([5]uint8{byte('h'), byte('e'), byte('l'), byte('l'), byte('o')})
			},
			wantErr: false,
		},
		{
			name: "slice int32 valid string assignment",
			args: args{
				assigner: func() reflect.Value {
					hello := []int32{}
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf([]int32("hello"))
			},
			wantErr: false,
		},
		{
			name: "invalid string value",
			args: args{
				assigner: func() reflect.Value {
					hello := "check"
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{0, 1, 2, 3})
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "valid string value",
			args: args{
				assigner: func() reflect.Value {
					hello := "check"
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf("hello")
			},
			wantErr: false,
		},
		{
			name: "different chan type",
			args: args{
				assigner: func() reflect.Value {
					hello := make(chan *int)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(make(chan int))
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "same map type",
			args: args{
				assigner: func() reflect.Value {
					hello := make(map[string]int)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(map[string]int{"hello": 5})
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(map[string]int{"hello": 5})
			},
			wantErr: false,
		},
		{
			name: "same time type",
			args: args{
				assigner: func() reflect.Value {
					hello := now
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(now.Add(time.Second))
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(now.Add(time.Second))
			},
			wantErr: false,
		},
		{
			name: "time assignment invalid val slice type",
			args: args{
				assigner: func() reflect.Value {
					hello := now
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf([]int{1, 2, 3})
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "time assignment invalid string format",
			args: args{
				assigner: func() reflect.Value {
					hello := now
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "time assignment valid string format",
			args: args{
				assigner: func() reflect.Value {
					hello := now
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("2019-08-22T11:43:21+07:00")
				},
			},
			wantAssigner: func() reflect.Value {
				past, _ := time.Parse(time.RFC3339, "2019-08-22T11:43:21+07:00")
				return reflect.ValueOf(past)
			},
			wantErr: false,
		},
		{
			name: "unimplemented unsafe pointer",
			args: args{
				assigner: func() reflect.Value {
					var x *int
					hello := unsafe.Pointer(x)
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf("hello")
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "assign map from struct",
			args: args{
				assigner: func() reflect.Value {
					hello := make(map[string]interface{})
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					type test struct {
						Hello string
						Hi    int
					}
					return reflect.ValueOf(test{"hello", 5})
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(map[string]interface{}{
					"Hello": "hello",
					"Hi":    5,
				})
			},
			wantErr: false,
		},
		{
			name: "assign map from map with different type",
			args: args{
				assigner: func() reflect.Value {
					hello := make(map[string]interface{})
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(map[string]int{"hello": 5, "hi": 10})
				},
			},
			wantAssigner: func() reflect.Value {
				return reflect.ValueOf(map[string]interface{}{
					"hello": 5,
					"hi":    10,
				})
			},
			wantErr: false,
		},
		{
			name: "assign map from int - error",
			args: args{
				assigner: func() reflect.Value {
					hello := make(map[string]interface{})
					return reflect.ValueOf(&hello)
				},
				val: func() reflect.Value {
					return reflect.ValueOf(5)
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
		{
			name: "assign struct from map",
			args: args{
				assigner: func() reflect.Value {
					type test struct {
						Hello string
						Hi    int
					}
					return reflect.ValueOf(&test{})
				},
				val: func() reflect.Value {
					return reflect.ValueOf(map[string]interface{}{"hello": "hola!", "hi": 10})
				},
			},
			wantAssigner: func() reflect.Value {
				type test struct {
					Hello string
					Hi    int
				}
				return reflect.ValueOf(test{"hola!", 10})
			},
			wantErr: false,
		},
		{
			name: "assign struct from int - error",
			args: args{
				assigner: func() reflect.Value {
					type test struct {
						Hello string
						Hi    int
					}
					return reflect.ValueOf(&test{})
				},
				val: func() reflect.Value {
					return reflect.ValueOf(5)
				},
			},
			wantAssigner: nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assigner := tt.args.assigner()
			if err := AssignReflect(assigner, tt.args.val()); (err != nil) != tt.wantErr {
				t.Errorf("AssignReflect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantAssigner != nil && tt.wantAssigner().IsValid() {
				assert.EqualValues(t, tt.wantAssigner().Interface(), GetChildElem(assigner).Interface())
			}
		})
	}
}
