package type_test

import (
	"testing"
)

type Myint int64

func TestImlicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)

	var c Myint
	c = Myint(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
