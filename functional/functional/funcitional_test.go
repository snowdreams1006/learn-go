package functional

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
	// 3 -1 2 0 unsupported operator: %
	t.Log(
		eval(1, 2, "+"),
		eval(1, 2, "-"),
		eval(1, 2, "*"),
		eval(1, 2, "/"),
		//eval(1, 2, "%"),
	)
}

func evalWithStandardStyle(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}

func TestEvalWithStandardStyle(t *testing.T) {
	// Success: 2
	if result, err := evalWithStandardStyle(5, 2, "/"); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}

	// Error: unsupported operator: %
	if result, err := evalWithStandardStyle(5, 2, "%"); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}
}

func evalWithApplyStyle(a, b int, op func(int, int) (int, error)) (int, error) {
	return op(a, b)
}

func divide(a, b int) (int, error) {
	return a / b, nil
}

func mod(a, b int) (int, error) {
	return a % b, nil
}

func TestEvalWithApplyStyle(t *testing.T) {
	// Success: 2
	if result, err := evalWithApplyStyle(5, 2, divide); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}

	// Success: 1
	if result, err := evalWithApplyStyle(5, 2, mod); err != nil {
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
