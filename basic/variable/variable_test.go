package variable

import "testing"

func TestVariableZeroValue(t *testing.T) {
	var a int
	var s string

	// 0
	t.Log(a, s)
	// 0 ""
	t.Logf("%d %q", a, s)
}

func TestVariableInitialValue(t *testing.T) {
	var a, b int = 1, 2
	var s string = "hello Go"

	// 1 2 hello Go
	t.Log(a, b, s)
}

func TestVariableShorter(t *testing.T) {
	var (
		a int    = 1
		b int    = 2
		s string = "hello go"
	)

	// 1 2 hello Go
	t.Log(a, b, s)
}

func TestVariableTypeDeduction(t *testing.T) {
	var a, b, s = 1, 2, "hello Go"

	// 1 2 hello Go
	t.Log(a, b, s)
}

func TestVariableTypeDeductionShorter(t *testing.T) {
	a, b, s := 1, 2, "hello Go"

	// 1 2 hello Go
	t.Log(a, b, s)

	s = "hello golang"

	// 1 2 hello golang
	t.Log(a, b, s)
}

var globalTestId = 2
// globalTestName := "type_test" is not supported
var globalTestName = "type_test"

func TestVariableScope(t *testing.T) {
	// 2 type_test
	t.Log(globalTestId, globalTestName)

	globalTestName = "TestVariableScope"

	// 2 TestVariableScope
	t.Log(globalTestId, globalTestName)
}
