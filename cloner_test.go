package reflecthelper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type str struct {
	Field string
	field string
}

func TestClone(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		s := &str{Field: "astr", field: "small field"}

		valClone := Clone(reflect.ValueOf(s))
		b := valClone.Interface().(*str)

		s.Field = "changed field"

		assert.NotEqual(t, s.field, b.field)
		fmt.Println(s, b)
	})
	t.Run("Clone nil value", func(t *testing.T) {
		var s *str

		valClone := Clone(reflect.ValueOf(s))
		b := valClone.Interface().(*str)

		s = &str{Field: "astr", field: "small field"}

		assert.NotEqual(t, s, b)
		fmt.Println(s, b)
	})
}

func TestInitNew(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		s := &str{field: "hello"}

		valInit := InitNew(reflect.ValueOf(s))
		b := valInit.Interface().(*str)
		assert.Nil(t, b)
	})
}
