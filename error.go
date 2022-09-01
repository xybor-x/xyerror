package xyerror

import (
	"errors"
	"fmt"
)

// XyError is a special error belongs to a generic error class.
type XyError struct {
	// error class
	c Class

	// error message
	msg string
}

// Error is the method to treat XyError as an error.
func (xerr XyError) Error() string {
	return fmt.Sprintf("%s: %s", xerr.c.name, xerr.msg)
}

// Is is the method used to customize errors.Is method.
func (xerr XyError) Is(target error) bool {
	if !errors.As(target, &Class{}) {
		return false
	}

	var tc = target.(Class)

	return xerr.c.belongsTo(tc)
}

// Or returns the first not-nil error. If all errors are nil, return nil.
func Or(errs ...error) error {
	for i := range errs {
		if errs[i] != nil {
			return errs[i]
		}
	}

	return nil
}
