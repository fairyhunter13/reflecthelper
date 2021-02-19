package reflecthelper

import "time"

// ParseTime parses the timeTxt string to the time.Time using various formats.
func ParseTime(timeTxt string, fnOpts ...FuncOption) (result time.Time, err error) {
	opt := NewOption().Assign(fnOpts...)
	result, err = parseTime(timeTxt, opt)
	return
}
