package closure

import (
	"fmt"
	"testing"
)

// 1 1 2 3 5 8 13 21 34 55
//     a b
//       a b
func fibonacciByNormal() {
	a, b := 0, 1

	for i := 0; i < 10; i++ {
		a, b = b, a+b

		fmt.Print(a, " ")
	}

	fmt.Println()
}

// 1 1 2 3 5 8 13 21 34 55
func TestFibonacciByNormal(t *testing.T) {
	fibonacciByNormal()
}

// 1 1 2 3 5 8 13 21 34 55
//     a b
//       a b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 1 1 2 3 5 8 13 21 34 55
func TestFibonacci(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
}

func autoIncrease() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func TestAutoIncrease(t *testing.T) {
	a := autoIncrease()

	// 1 2 3
	t.Log(a(), a(), a())

	b := autoIncrease()

	// 1 2 3
	t.Log(b(), b(), b())
}

func countByClosureButWrong() []func() int {
	var arr []func() int
	for i := 1; i <= 3; i++ {
		arr = append(arr, func() int {
			return i
		})
	}
	return arr
}

func TestCountByClosure(t *testing.T) {
	// 4 4 4
	for _, c := range countByClosureButWrong() {
		t.Log(c())
	}
}

func countByClosureWithOk() []func() int {
	var arr []func() int
	for i := 1; i <= 3; i++ {
		func(n int) {
			arr = append(arr, func() int {
				return n
			})
		}(i)
	}
	return arr
}

func TestCountByClosureWithOk(t *testing.T) {
	// 1 2 3
	for _, c := range countByClosureWithOk() {
		t.Log(c())
	}
}

func sumByNormal(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func TestSumByNormal(t *testing.T) {
	arr := []int{1, 2, 3}

	// 6
	t.Log(sumByNormal(arr))

	// 6
	t.Log(sumByNormal(arr))
}

func sumByClosure() func(arr []int) int {
	sum := 0
	return func(arr []int) int {
		for _, v := range arr {
			sum += v
		}
		return sum
	}
}

func TestSumClosure(t *testing.T) {
	arr := []int{1, 2, 3}

	s := sumByClosure()

	// 6
	t.Log(s(arr))

	// 12
	t.Log(s(arr))

	ns := sumByClosure()

	// 6
	t.Log(ns(arr))
}