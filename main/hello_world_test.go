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
	for i := 0; i < len(days); i++ {
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

	var access [3]int
	for i := 0; i < len(access); i++ {
		access[i] = 1 << uint(i)
	}
	// [1 2 4]
	t.Log(access)

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

	const (
		// iota=0 const=1+0=1 iota=0+1=1
		first = 1 + iota

		// iota=1 const=1+1=2 iota=1+1=2
		second

		// iota=2 const=2+2=4 iota=2+1=3
		third = 2 + iota

		// iota=3 const=2+3=5 iota=3+1=4
		forth

		// iota=4 const=2*4=8 iota=4+1=5
		fifth = 2 * iota

		// iota=5 const=2*5=10 iota=5+1=6
		sixth

		// iota=6 const=6 iota=6+1=7
		seventh = iota
	)
	// 1 2 4 5 8 10 6
	t.Log(first, second, third, forth, fifth, sixth, seventh)

	var rank [7]int
	for i := 0; i < len(rank); i++ {
		if i < 2 {
			rank[i] = 1 + i
		} else if i < 4 {
			rank[i] = 2 + i
		} else if i < 6 {
			rank[i] = 2 * i
		} else {
			rank[i] = i
		}
	}
	// [1 2 3 4 5 6 7]
	t.Log(rank)

	const currentIota = iota

	// 0
	t.Log(currentIota)
}
