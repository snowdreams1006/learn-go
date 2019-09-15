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