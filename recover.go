package reflecthelper

import "fmt"

// RecoverFn is a function used to recover from panic situation by passing the pointer of the error.
var RecoverFn = func(err *error) {
	if err == nil {
		return
	}

	if rec := recover(); rec != nil {
		switch val := rec.(type) {
		case error:
			*err = val
		default:
			*err = fmt.Errorf("%v", val)
		}
	}
}
