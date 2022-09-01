[![Xybor founder](https://img.shields.io/badge/xybor-huykingsofm-red)](https://github.com/huykingsofm)
[![Go Reference](https://pkg.go.dev/badge/github.com/xybor-x/xyerror.svg)](https://pkg.go.dev/github.com/xybor-x/xyerror)
[![GitHub Repo stars](https://img.shields.io/github/stars/xybor-x/xyerror?color=yellow)](https://github.com/xybor-x/xyerror)
[![GitHub top language](https://img.shields.io/github/languages/top/xybor-x/xyerror?color=lightblue)](https://go.dev/)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xybor-x/xyerror)](https://go.dev/blog/go1.18)
[![GitHub release (release name instead of tag name)](https://img.shields.io/github/v/release/xybor-x/xyerror?include_prereleases)](https://github.com/xybor-x/xyerror/releases/latest)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/5fae4ad52c184d6dbec00f44312f20d6)](https://www.codacy.com/gh/xybor-x/xyerror/dashboard?utm_source=github.com&utm_medium=referral&utm_content=xybor-x/xyerror&utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/5fae4ad52c184d6dbec00f44312f20d6)](https://www.codacy.com/gh/xybor-x/xyerror/dashboard?utm_source=github.com&utm_medium=referral&utm_content=xybor-x/xyerror&utm_campaign=Badge_Coverage)
[![Go Report](https://goreportcard.com/badge/github.com/xybor-x/xyerror)](https://goreportcard.com/report/github.com/xybor-x/xyerror)

# Introduction

Package xyerror supports to define error types conveniently.

The error is inspired by the idea of Python `Exception`.

# Features

## Python Exception Idea

Xyerror defined an error type used to create other errors called `Class`, it is
equivalent to `Exception` class in Python.

`Class` creates `XyError` objects by using `New` method. It looks like
`Exception` class creates `Exception` instances. Example:

```golang
// xyerror
var xerr xyerror.XyError
xerr = xyerror.ValueError.New("value is invalid")
fmt.Println(xerr)
```

```python
# python
print(ValueError("value is invalid"))
```

A `Class` can "inherits" from another `Class`. Example:

```golang
// xyerror
var ZeroDivisionError = xyerror.ValueError.NewClass("ZeroDivisionError")
fmt.Println(ZeroDivisionError.New("do not divide by zero"))
```

```python
# python
class ZeroDivisionError(ValueError):
    ...
print(ZeroDivisionError("do not divide by zero"))
```

A `Class` also "inherits" from some `Class` instances. Example:

```golang
// xyerror
var ValueTypeError = xyerror.
    Combine(xyerror.ValueError, xyerror.TypeError).
    NewClass("ValueTypeError")
```

```python
# python
class ValueTypeError(ValueError, TypeError):
    ...
```

A `XyError` created by `Class` can compare to this `Class`.

```golang
// xyerror
func foo() error {
    return xyerror.ValueError.New("value is invalid")
}

func bar() error {
    if xerr := foo(); errors.Is(xerr, xyerror.ValueError) {
        fmt.Println(xerr)
    } else {
        return xerr
    }
    return nil
}
```

```python
# python
def foo():
    raise ValueError("value is invalid")

def bar():
    try:
        foo()
    except ValueError as e:
        print(e)
```

## Error Number

Every `Class` has its own unique number called `errno`. This number is used to
compare a `XyError` with a `Class`. All `XyError` instances created by the same
`Class` have the same `errno`.

## Module-oriented Error

Xyerror is tended to create module-oriented errors. `Errno` of all `Class`
instances in a module need to be the same prefix.

For example, `ValueError` and `TypeError` are in `Default` module. So their
`errno` should be 100001 and 100002, respectively. In this case, 100000 is the
prefix number of `Default` module.

Every module needs to create an object called `Generator` to create its error
`Class` instances. Prefix of module is a self-defined value, it must be
divisible by 100000 and not equal to other modules' prefix.

```golang
var egen = xyerror.Register("XyExample", 200000)
var (
    FooError = egen.NewClass("FooError")
    BarError = egen.NewClass("BarError")
)
```

# Example

```golang
package xyerror_test

import (
	"errors"
	"fmt"

	"github.com/xybor-x/xyerror"
)

var exampleGen = xyerror.Register("example", 400000)

func ExampleClass() {
	// To create a root Class, call Generator.NewClass with the name of Class.
	var RootError = exampleGen.NewClass("RootError")

	// You can create a class from another one.
	var ChildError = RootError.NewClass("ChildError")

	fmt.Println(RootError)
	fmt.Println(ChildError)

	// Output:
	// [400001] RootError
	// [400002] ChildError
}

func ExampleXyError() {
	// You can compare a XyError with an Class by using the built-in method
	// errors.Is.
	var NegativeIndexError = xyerror.IndexError.NewClass("NegativeIndexError")

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

func ExampleGroup() {
	// Group allows you to create a class with multiparents.
	var KeyValueError = xyerror.
		Combine(xyerror.KeyError, xyerror.ValueError).
		NewClass(exampleGen, "KeyValueError")

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
```
