package function

import (
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

func evalByDivide(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result, _ = divide(a, b)
	default:
		panic("unsupported operator: " + op)
	}
	return result
}

func TestEvalByDivide(t *testing.T) {
	// 2
	t.Log(evalByDivide(5, 2, "/"))
}
