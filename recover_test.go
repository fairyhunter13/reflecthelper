package reflecthelper

import (
	"errors"
	"testing"
)

func Test_RecoverFn(t *testing.T) {
	t.Run("err var is nil", func(t *testing.T) {
		RecoverFn(nil)
	})
	t.Run("panic mode, string panic", func(t *testing.T) {
		var err error
		defer RecoverFn(&err)
		panic("Random")
	})
	t.Run("panic mode, error panic", func(t *testing.T) {
		var err error
		defer RecoverFn(&err)
		panic(errors.New("hello"))
	})
}
