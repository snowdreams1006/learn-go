package fib

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	// 1 2 3
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

func evalWithFunctionalStyle(a, b int, op func(int, int) (int, error)) func() (int, error) {
	return func() (int, error) {
		return op(a, b)
	}
}

func pow(a, b int) (int, error) {
	return int(math.Pow(float64(a), float64(b))),nil
}

func TestEvalWithFunctionalStyle(t *testing.T) {
	ef := evalWithFunctionalStyle(5, 2, pow)

	// Success: 25
	if result, err := ef(); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}
}

func powWithClosure(a int) func(b int)(int, error) {
	return func(b  int) (int, error) {
		return int(math.Pow(float64(a), float64(b))),nil
	}
}

func TestPowWithClosure(t *testing.T) {
	pow5 := powWithClosure(5)

	// Success: 25
	if result, err := pow5(2); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}

	pow2 := powWithClosure(2)

	// Success: 32
	if result, err := pow2(5); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}
}