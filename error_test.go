// Copyright (c) 2022 xybor-x
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package xyerror_test

import (
	"errors"
	"testing"

	"github.com/xybor-x/xycond"
	"github.com/xybor-x/xyerror"
)

func TestError(t *testing.T) {
	var exc = xyerror.NewException("Error")
	var err1 = exc.Newf("error-%d", 1)
	var err2 = exc.New("error-2")

	xycond.ExpectEqual(err1.Error(), "Error: error-1").Test(t)
	xycond.ExpectEqual(err2.Error(), "Error: error-2").Test(t)
}

func TestErrorIs(t *testing.T) {
	var err1 = xyerror.ValueError.New("err1")
	var err2 = xyerror.TypeError.New("err2")

	xycond.ExpectError(err1, xyerror.ValueError).Test(t)
	xycond.ExpectError(err2, xyerror.TypeError).Test(t)

	xycond.ExpectErrorNot(err1, err1).Test(t)
	xycond.ExpectErrorNot(err1, err2).Test(t)
	xycond.ExpectErrorNot(err1, xyerror.TypeError).Test(t)
}

func TestOr(t *testing.T) {
	var err1 = xyerror.ValueError.New("err1")
	var err2 = xyerror.TypeError.New("err2")
	var err3 error

	xycond.ExpectError(xyerror.Or(err1, err2), xyerror.ValueError).Test(t)
	xycond.ExpectError(xyerror.Or(err2, err1), xyerror.TypeError).Test(t)
	xycond.ExpectNil(xyerror.Or(nil, err3)).Test(t)
}

func TestCombine(t *testing.T) {
	var exc = xyerror.Combine(xyerror.ValueError, xyerror.TypeError).
		NewException("ValueTypeError")
	var err = exc.New("foo")

	xycond.ExpectError(err, xyerror.ValueError).Test(t)
	xycond.ExpectError(err, xyerror.TypeError).Test(t)
	xycond.ExpectErrorNot(err, xyerror.IndexError).Test(t)
}

func TestMessage(t *testing.T) {
	var e1 = xyerror.ValueError.New("foo-bar")
	var e2 = errors.New("foo-foo")
	xycond.ExpectEqual(e1.Error(), "ValueError: foo-bar").Test(t)
	xycond.ExpectEqual(xyerror.Message(e1), "Foo-Bar").Test(t)
	xycond.ExpectEqual(xyerror.Message(e2), "Foo-Foo").Test(t)
}
