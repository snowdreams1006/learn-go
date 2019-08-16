package function

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"testing"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator: " + op)
	}
	return result
}

func TestEval(t *testing.T) {
	t.Log(
		eval(1, 2, "+"),
		eval(1, 2, "-"),
		eval(1, 2, "*"),
		eval(1, 2, "/"),
		eval(1, 2, "%"),
	)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func TestDivide(t *testing.T) {
	// 2 1
	t.Log(divide(5, 2))
}

func divideReturnName(a, b int) (q, r int) {
	return a / b, a % b
}

func TestDivideReturnName(t *testing.T) {
	q, r := divideReturnName(5, 2)

	// 2 1
	t.Log(q, r)
}

func evalByDivide(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := divide(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}

func TestEvalByDivide(t *testing.T) {
	// 2
	t.Log(evalByDivide(5, 2, "/"))

	// 0 unsupported operator: %
	t.Log(evalByDivide(5, 2, "%"))

	// Error: unsupported operator: %
	if result, err := evalByDivide(5, 2, "%"); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args (%d,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func TestApply(t *testing.T) {
	// 1
	t.Log(apply(func(a int, b int) int {
		return a % b
	}, 5, 2))

	// 25
	t.Log(apply(pow, 5, 2))
}

func sum(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func TestSum(t *testing.T) {
	// 15
	t.Log(sum(1, 2, 3, 4, 5))
}
