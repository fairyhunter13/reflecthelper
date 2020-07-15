package reflecthelper

import (
	"reflect"

	"github.com/Popog/deepcopy"
)

// Clone clones the current underlying value of val as a reflect.Value.
// Clone can't clone unexported struct field because it is inaccessible.
func Clone(val reflect.Value) (res reflect.Value) {
	res = deepcopy.DeepCopyValue(val)
	return
}

// InitNew initializes a new reflect.Value with reflect.Type of val.
func InitNew(val reflect.Value) (res reflect.Value) {
	if !val.IsValid() {
		return
	}

	res = reflect.Indirect(reflect.New(val.Type()))
	return
}
