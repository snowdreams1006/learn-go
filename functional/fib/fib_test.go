package fib

import (
	"fmt"
	"testing"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestFibonacci(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f()," ")
	}
	fmt.Println()
}
