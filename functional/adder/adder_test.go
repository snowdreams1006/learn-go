package adder

import (
	"fmt"
	"testing"
	"time"
)

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

type iAdder func(int) (int, iAdder)

func mAdder(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, mAdder(base + v)
	}
}

func TestMAdder(t *testing.T) {
	m := mAdder(0)

	for i := 0; i < 10; i++ {
		var s int
		s, m = m(i)

		t.Logf("0+...+%d=%d", i, s)
	}
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
	t.Log(a(),a(),a())
}

func timeSpend(fn func()) func() {
	return func()  {
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