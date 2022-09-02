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
