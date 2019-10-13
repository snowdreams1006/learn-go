package error

import (
	"testing"
)

func TestFuncWithoutDefer(t *testing.T) {
	// 1 2
	t.Log(1)
	t.Log(2)
}

func TestFuncWithDefer(t *testing.T) {
	// 2 1
	defer t.Log(1)
	t.Log(2)
}