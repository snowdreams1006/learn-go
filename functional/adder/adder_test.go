package adder

import "testing"

func adder() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func TestAdder(t *testing.T) {
	a := adder()
	for i := 0; i < 10; i++ {
		t.Logf("0+...+%d=%d", i, a(i))
	}
}
