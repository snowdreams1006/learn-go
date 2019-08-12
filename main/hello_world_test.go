package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("Hello Test")
}

func TestFib(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var a = 1
	//var b = 1

	//var (
	//	a = 1
	//	b = 1
	//)

	//var a, b = 1, 1

	a, b := 1, 1

	fmt.Print(a)
	for i := 0; i < 6; i++ {
		fmt.Print(" ", b)

		temp := a
		a = b
		b = temp + b
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a, b := 1, 2
	t.Log(a, b)

	a, b = b, a
	t.Log(a, b)

	temp := a
	a = b
	b = temp
	t.Log(a, b)
}

func TestConst(t *testing.T) {
	const pi = 3.14
	t.Log(pi)

	// cannot assign to pi
	//pi = 2.828
	//t.Log(pi)
}

func TestConstForIota(t *testing.T) {
	const (
		Mon = 1 + iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	// 1 2 3 4 5 6 7
	t.Log(Mon, Tue, Wed, Thu, Fri, Sat, Sun)

	var days [7]int
	for i := 0; i < 7; i++ {
		days[i] = 1 + i
	}
	// [1 2 3 4 5 6 7]
	t.Log(days)

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
	// 0110 即 6,表示不可读,可写,可执行
	accessCode = 6
	t.Log(accessCode&Readable == Readable, accessCode&Writing == Writing, accessCode&Executable == Executable)
	// 0100 即 4,表示不可读,不可写,可执行
	accessCode = 4
	t.Log(accessCode&Readable == Readable, accessCode&Writing == Writing, accessCode&Executable == Executable)
	// 0000 即 0,表示不可读,不可写,不可执行
	accessCode = 0
	t.Log(accessCode&Readable == Readable, accessCode&Writing == Writing, accessCode&Executable == Executable)

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

	// 62.9 KB (64,411 字节)
	size := 64411.0
	t.Log(size, size/Kb)
}
