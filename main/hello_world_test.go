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

	//pi = 2.828
	//t.Log(pi)
}
