package reflecthelper

import "errors"

// List of all errors for reflecthelper.
var (
	ErrAssignerCantSet = errors.New("Assigner doesn't have the ability to set the value")
)
