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
	"strings"
	"testing"

	"github.com/xybor-x/xycond"
	"github.com/xybor-x/xyerror"
)

func TestException(t *testing.T) {
	var c1 = xyerror.NewException("class1")
	var c2 = c1.NewException("class2")
	xycond.ExpectTrue(strings.Contains(c1.Error(), "class1")).Test(t)
	xycond.ExpectTrue(strings.Contains(c2.Error(), "class2")).Test(t)
}

func TestExceptionSameName(t *testing.T) {
	var c = xyerror.NewException(t.Name())
	xycond.ExpectNotPanic(func() { c.NewException(t.Name()) }).Test(t)
}

func TestExceptionNewError(t *testing.T) {
	var e1 = xyerror.ValueError.New(t.Name())
	var e2 = xyerror.TypeError.New(e1)
	xycond.ExpectEqual(e2.Error(), "TypeError: "+t.Name()).Test(t)
}
