package functional

import (
	"fmt"
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

	// Success: 5
	if result, err := evalWithApplyStyle(5, 2, func(a int, b int) (result int, e error) {
		if a > b {
			return a, nil
		}
		return b, nil
	}); err != nil {
		t.Log("Error:", err)
	} else {
		t.Log("Success:", result)
	}
}
