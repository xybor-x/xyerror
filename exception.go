package xyerror

import (
	"errors"
	"fmt"

	"github.com/xybor-x/xylock"
)

var counter = 0
var exceptions []string
var lock = xylock.Lock{}

// Exception is the generic type of Errors. It does not contain the error
// message. You should use an Error by creating from Exception, instead of using
// Exception directly.
type Exception struct {
	// The unique number is used to differentiate Exceptions.
	id int

	// The Exception name
	name string

	// The parent Exceptions
	parent []Exception
}

// NewException creates a root Exception.
func NewException(name string) Exception {
	lock.Lock()
	defer lock.Unlock()

	for i := range exceptions {
		if exceptions[i] == name {
			panic("Do not use an existed Exception name: " + name)
		}
	}
	counter++
	exceptions = append(exceptions, name)

	return Exception{
		id:     counter,
		name:   name,
		parent: nil,
	}
}

// NewException creates a new Exception by inheriting the called Exception.
func (exc Exception) NewException(name string) Exception {
	var class = NewException(name)
	class.parent = []Exception{exc}
	return class
}

// Newf creates an Error with a formatting message.
func (exc Exception) Newf(msg string, a ...any) Error {
	return Error{exc: exc, msg: fmt.Sprintf(msg, a...)}
}

// New creates an Error with default formatting objects.
func (exc Exception) New(a ...any) Error {
	return Error{exc: exc, msg: fmt.Sprint(a...)}
}

// Is returns true if the Exception is inherited from the target.
func (exc Exception) Is(target error) bool {
	if !errors.As(target, &Exception{}) {
		return false
	}

	var exception = target.(Exception)

	if exc.id == exception.id {
		return true
	}

	for i := range exc.parent {
		if errors.Is(exc.parent[i], target) {
			return true
		}
	}

	return false
}

// Error returns the Exception name.
func (exc Exception) Error() string {
	return exc.name
}

// CombinedException is an array of Exception. It supports to creates an
// Exception inherited from many parents.
type CombinedException []Exception

// Combine supports creating a group of Exceptions. This group can be used to
// create the Exception with multiparents.
func Combine(cs ...Exception) CombinedException {
	return cs
}

// NewException creates an Exception with multiparents.
func (combined CombinedException) NewException(name string) Exception {
	var child = NewException(name)
	child.parent = combined
	return child
}
