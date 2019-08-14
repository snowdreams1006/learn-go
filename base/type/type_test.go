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

func TestString(t *testing.T) {
	var s string
	t.Log(len(s))

	if s == "" {
		t.Log("空字符串", s)
	} else {
		t.Log("非空字符串", s)
	}
}

func TestArithmeticOperator(t *testing.T) {
	a := 0
	// 0
	t.Log(a)

	a = a + 1
	// 1
	t.Log(a)

	a = a * 2
	// 2
	t.Log(a)

	a = a % 2
	// 0
	t.Log(a)

	a++
	// 1
	t.Log(a)
}

func TestComparisonOperator(t *testing.T) {
	a, b := 0, 1
	t.Log(a, b)

	t.Log(a > b)
	t.Log(a < b)
	t.Log(a != b)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	//b := [...]int{2, 4}
	c := [...]int{1, 3, 3}
	d := [...]int{1, 2, 4}

	// a == b --> invalid operation: a == b (mismatched types [3]int and [2]int)
	//t.Log(a == b)
	t.Log(a == c)
	t.Log(a == d)
}

func TestLogicalOperator(t *testing.T) {
	a, b := true, false
	t.Log(a, b)

	t.Log(a&&b,a||b,!a,!b)
}
