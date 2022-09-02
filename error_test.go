package xyerror_test

import (
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
