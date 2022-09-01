package xyerror

import (
	"fmt"

	"github.com/xybor-x/xylock"
)

var counter = 0
var classes []string
var lock = xylock.Lock{}

// Class is the generic type for XyError.
type Class struct {
	// The unique number of each Class.
	errno int

	// The error name
	name string

	// The parent classes
	parent []Class
}

// NewClass creates a root Class.
func NewClass(name string) Class {
	lock.Lock()
	defer lock.Unlock()

	for i := range classes {
		if classes[i] == name {
			panic("Do not use an existed error class name: " + name)
		}
	}
	counter++
	classes = append(classes, name)

	return Class{
		errno:  counter,
		name:   name,
		parent: nil,
	}
}

// NewClass creates a new Class by inheriting the called Class.
func (c Class) NewClass(name string) Class {
	var class = NewClass(name)
	class.parent = []Class{c}
	return class
}

// Newf creates a XyError with a formatting message.
func (c Class) Newf(msg string, a ...any) XyError {
	return XyError{c: c, msg: fmt.Sprintf(msg, a...)}
}

// New creates a XyError with default formatting objects.
func (c Class) New(a ...any) XyError {
	return XyError{c: c, msg: fmt.Sprint(a...)}
}

// belongsTo checks if a Class is inherited from a target class. A class belongs
// to the target Class if it is created by the target itself or target's child.
func (c Class) belongsTo(t Class) bool {
	if c.errno == t.errno {
		return true
	}

	for i := range c.parent {
		if c.parent[i].belongsTo(t) {
			return true
		}
	}

	return false
}

// Error is the method to treat Class as an error.
func (c Class) Error() string {
	return c.name
}
