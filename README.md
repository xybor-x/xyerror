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

Package xyerror supports to define and compare errors conveniently.

The error is inspired by Python `Exception`.

# Features

Package xyerror defined an error type used to create other errors called
`Exception`, it is equivalent to `Exception` class in Python.

`Exception` creates `Error` objects by using `New` or `Newf` method. It looks
like `Exception` class creates `Exception` instances. Example:

```golang
// xyerror
var err xyerror.Error = xyerror.ValueError.New("value is invalid")
fmt.Println(err)
```

```python
# python
print(ValueError("value is invalid"))
```

An `Exception` can "inherits" from another `Exception`. Example:

```golang
// xyerror
var ZeroDivisionError = xyerror.ValueError.NewException("ZeroDivisionError")
fmt.Println(ZeroDivisionError.New("do not divide by zero"))
```

```python
# python
class ZeroDivisionError(ValueError):
    ...
print(ZeroDivisionError("do not divide by zero"))
```

An `Exception` also "inherits" from many `Exception` instances. Example:

```golang
// xyerror
var ValueTypeError = xyerror.
    Combine(xyerror.ValueError, xyerror.TypeError).
    NewException("ValueTypeError")
```

```python
# python
class ValueTypeError(ValueError, TypeError):
    ...
```

An `Error` created by an `Exception` can compare to that `Exception`.

```golang
// xyerror
func foo() error {
    return xyerror.ValueError.New("value is invalid")
}

func bar() error {
    if err := foo(); errors.Is(err, xyerror.ValueError) {
        fmt.Println(err)
    } else {
        return err
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

# Example

```golang
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
```
