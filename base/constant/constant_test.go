package constant

import (
	"math"
	"testing"
)

func TestConstant(t *testing.T) {
	const a, b = 3, 4
	const s = "hello Go"

	// 3 4 hello Go
	t.Log(a, b, s)

	var c int
	c = int(math.Sqrt(a*a + b*b))

	// 3 4 5
	t.Log(a, b, c)
}

func TestConstant2Enum(t *testing.T) {
	const (
		java = 0
		golang = 1
		cpp = 2
		python = 3
	)

	// 0 1 2 3
	t.Log(java, golang, cpp,python)
}

func TestConstant2EnumShorter(t *testing.T) {
	const (
		java = iota
		golang
		_
		python
		javascript
	)

	// 0 1 3 4
	t.Log(java, golang,python,javascript)
}

func TestConstantIota(t *testing.T) {
	const (
		B = 1 << (10 * iota)
		Kb
		Mb
		Gb
		Tb
		Pb
	)

	// 1 1024 1048576 1073741824 1099511627776 1125899906842624
	t.Log(B, Kb, Mb, Gb, Tb, Pb)
}

func TestConstantIotaBitCalculate(t *testing.T){
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	// 0001 0010 0100 即 1 2 4
	t.Log(Readable, Writable, Executable)

	// 0111 即 7,表示可读,可写,可执行
	accessCode := 7
	t.Log(accessCode&Readable == Readable, accessCode&Writable == Writable, accessCode&Executable == Executable)
}