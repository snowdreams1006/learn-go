package _type

import "testing"

type MyInt64 int64

func TestImplicitTypeConvert(t *testing.T) {
	var a int = 1
	var b int64

	b = int64(a)
	t.Log(a, b)

	var c MyInt64

	c = MyInt64(b)
	t.Log(b, c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPoi := &a

	//aPoi = aPoi + 1

	t.Log(a, aPoi)
	t.Logf("%T %T", a, aPoi)
}

func TestString(t *testing.T){
	var s string
	t.Log(len(s))

	if s == "" {
		t.Log("空字符串",s)
	}else{
		t.Log("非空字符串",s)
	}
}
