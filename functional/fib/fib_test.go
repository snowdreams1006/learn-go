package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

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

func TestFibonacci(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestIntGenRead(t *testing.T) {
	f := fib()
	printFileContents(f)
}

func autoIncrease() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestAutoIncrease(t *testing.T) {
	a := autoIncrease()

	// 1 2 3
	t.Log(a(), a(), a())
}

func timeSpend(fn func()) func() {
	return func() {
		start := time.Now()

		fn()

		fmt.Println("time spend : ", time.Since(start).Seconds())
	}
}

func slowFunc() {
	time.Sleep(time.Second * 1)

	fmt.Println("I am slowFunc")
}

func TestSlowFuncTimeSpend(t *testing.T) {
	slowFuncTimeSpend := timeSpend(slowFunc)

	slowFuncTimeSpend()
}

func count() []int {
	var arr []int
	for i := 1; i <= 3; i++ {
		arr = append(arr, i)
	}
	return arr
}

func TestCount(t *testing.T) {
	for _, c := range count() {
		t.Log(c)
	}
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
	for _, c := range countByClosureWithOk() {
		t.Log(c())
	}
}
