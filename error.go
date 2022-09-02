package xyerror

import (
	"errors"
	"fmt"
)

// Error is a special error belongs to a generic error class.
type Error struct {
	exc Exception
	msg string
}

// Error returns the Exception name along with the Error message.
func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.exc.name, err.msg)
}

// Is returns true if Error is created by the target Exception.
func (err Error) Is(target error) bool {
	return errors.Is(err.exc, target)
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
