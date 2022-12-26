package try_test

import "testing"

type MyInt int64

func TestComplicit(t *testing.T) {
	var a int32 = 1
	var b MyInt
	b = MyInt(a)
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
