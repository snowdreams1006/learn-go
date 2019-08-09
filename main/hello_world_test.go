package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("Hello Test")
}

func TestFib(t *testing.T) {
	var a = 1
	var b = 1

	fmt.Print(a)
	for i := 0; i < 6; i++ {
		fmt.Print(" ", b)

		temp := a
		a = b
		b = temp + b
	}
	fmt.Println()
}
