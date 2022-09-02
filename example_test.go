package xyerror_test

import (
	"errors"
	"fmt"

	"github.com/xybor-x/xyerror"
)

func ExampleException() {
	// To create a root Exception, call xyerror.NewException with the its name.
	var RootError = xyerror.NewException("RootError")

	// You can create an Exception by inheriting from another one.
	var ChildError = RootError.NewException("ChildError")

	fmt.Println(RootError)
	fmt.Println(ChildError)

	// Output:
	// RootError
	// ChildError
}

func ExampleError() {
	// You can compare an Error with an Exception by using the built-in method
	// errors.Is.
	var NegativeIndexError = xyerror.IndexError.NewException("NegativeIndexError")

	var err1 = xyerror.ValueError.New("some value error")
	if errors.Is(err1, xyerror.ValueError) {
		fmt.Println("err1 is a ValueError")
	}
	if !errors.Is(err1, NegativeIndexError) {
		fmt.Println("err1 is not a NegativeIndexError")
	}

	var err2 = NegativeIndexError.Newf("some negative index error %d", -1)
	if errors.Is(err2, NegativeIndexError) {
		fmt.Println("err2 is a NegativeIndexError")
	}
	if errors.Is(err2, xyerror.IndexError) {
		fmt.Println("err2 is a IndexError")
	}
	if !errors.Is(err2, xyerror.ValueError) {
		fmt.Println("err2 is not a ValueError")
	}

	// Output:
	// err1 is a ValueError
	// err1 is not a NegativeIndexError
	// err2 is a NegativeIndexError
	// err2 is a IndexError
	// err2 is not a ValueError
}

func ExampleCombinedException() {
	// CombinedException allows you to create an Exception with multiparents.
	var KeyValueError = xyerror.
		Combine(xyerror.KeyError, xyerror.ValueError).
		NewException("KeyValueError")

	var err = KeyValueError.New("something is wrong")

	if errors.Is(err, xyerror.KeyError) {
		fmt.Println("err is a KeyError")
	}

	if errors.Is(err, xyerror.ValueError) {
		fmt.Println("err is a ValueError")
	}

	// Output:
	// err is a KeyError
	// err is a ValueError
}
