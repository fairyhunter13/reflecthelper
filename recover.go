package reflecthelper

import "fmt"

// recoverFnOpt is used to recover from panic using the opt.RecoverPanic toggle.
func recoverFnOpt(err *error, opt *Option) {
	if err == nil || opt == nil || !opt.RecoverPanic {
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

// RecoverFn is used to recover from panic situation by passing the pointer of the error.
func RecoverFn(err *error) {
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
