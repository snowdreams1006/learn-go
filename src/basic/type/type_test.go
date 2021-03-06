package _type

import (
	"math"
	"math/cmplx"
	"testing"
)

func TestComplex(t *testing.T) {
	c := 3 + 4i

	// 5
	t.Log(cmplx.Abs(c))
}

func TestEuler(t *testing.T) {
	// (0+1.2246467991473515e-16i)
	t.Log(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	// (0+1.2246467991473515e-16i)
	t.Log(cmplx.Exp(1i*math.Pi) + 1)

	// (0.000+0.000i)
	t.Logf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

func TestExplicitTypeConvert(t *testing.T) {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	// 3 4 5
	t.Log(a, b, c)
}

func TestImplicitTypeConvert(t *testing.T) {
	type MyInt64 int64

	var a int = 1
	var b int64

	// b = a : cannot use a (type int) as type int64 in assignment
	b = int64(a)
	t.Log(a, b)

	var c MyInt64

	// c = b : cannot use b (type int64) as type MyInt64 in assignment
	c = MyInt64(b)
	t.Log(b, c)
}

func TestPointer(t *testing.T) {
	var a int = 1
	var pa *int = &a

	// 0xc0000921d0 1 1
	t.Log(pa, *pa, a)

	*pa = 2

	// 0xc0000901d0 2 2
	t.Log(pa, *pa, a)
}

func TestPointerShorter(t *testing.T) {
	a := 1
	pa := &a

	// 0xc0000e6010 1 1
	t.Log(pa, *pa, a)

	*pa = 2

	// 0xc0000e6010 2 2
	t.Log(pa, *pa, a)

	// pa = pa + 1 : invalid operation: pa + 1 (mismatched types *int and int)
	//pa = pa + 1

	// *int int int
	t.Logf("%T %T %T", pa, *pa,a)
}

func swapByVal(a, b int) {
	a, b = b, a
}

func TestSwapByVal(t *testing.T) {
	a, b := 3, 4

	swapByVal(a, b)

	// 3 4
	t.Log(a, b)
}

func swapByRef(a, b *int) {
	*a, *b = *b, *a
}

func TestSwapByRef(t *testing.T) {
	a, b := 3, 4

	swapByRef(&a, &b)

	// 4 3
	t.Log(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func TestSwap(t *testing.T) {
	a, b := 3, 4

	a, b = swap(a, b)

	// 4 3
	t.Log(a, b)
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

	// false true true
	t.Log(a > b, a < b, a != b)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	//b := [...]int{2, 4}
	c := [...]int{1, 2, 3}
	d := [...]int{1, 2, 4}

	// a == b --> invalid operation: a == b (mismatched types [3]int and [2]int)
	//t.Log(a == b)

	// true false
	t.Log(a == c,a == d)
}

func TestLogicalOperator(t *testing.T) {
	a, b := true, false
	t.Log(a, b)

	// false true false true
	t.Log(a && b, a || b, !a, !b)
}

func TestBitOperator(t *testing.T) {
	a, b := 1, 2
	t.Log(a, b)

	// 0 3 3
	t.Log(a&b, a|b, a^b)

	// 2 1
	t.Log(a<<1, b>>1)
}

func TestClearZeroOperator(t *testing.T) {
	// 0 0 1 0
	t.Log(1&^1, 0&^1, 1&^0, 0&^1)

	const (
		Readable = 1 << iota
		Writing
		Executable
	)
	// 0001 0010 0100 即 1 2 4
	t.Log(Readable, Writing, Executable)

	// 0111 即 7,表示可读,可写,可执行
	accessCode := 7
	t.Log(accessCode&Readable == Readable, accessCode&Writing == Writing, accessCode&Executable == Executable)

	// 0111 &^ 0001 = 0110 即清除可读权限
	accessCode = accessCode &^ Readable
	t.Log(accessCode&Readable == Readable, accessCode&Writing == Writing, accessCode&Executable == Executable)

	// 0110 &^ 0010 = 0100
	accessCode = accessCode &^ Writing
	t.Log(accessCode&Readable == Readable, accessCode&Writing == Writing, accessCode&Executable == Executable)
}
