package error

import (
	"testing"
)

func TestFuncWithoutDefer(t *testing.T) {
	// 1 2,「雪之梦技术驿站」: 正常顺序
	t.Log(1)
	t.Log(2)
}

func TestFuncWithDefer(t *testing.T) {
	// 2 1,「雪之梦技术驿站」: 正常顺序执行完毕后才执行 defer 代码
	defer t.Log(1)
	t.Log(2)
}

func TestFuncWithMultipleDefer(t *testing.T) {
	// 3 2 1,「雪之梦技术驿站」: 猜测 defer 底层实现数据结构可能是栈,先进后出.
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
}

func TestFuncWithMultipleDeferAndReturn(t *testing.T) {
	// 3 2 1,「雪之梦技术驿站」: defer 在于保证无论正常代码如何执行,该逻辑代码也一定会执行!
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
	return
	t.Log(4)
}

func TestFuncWithMultipleDeferAndPanic(t *testing.T) {
	// 3 2 1,「雪之梦技术驿站」: defer 在于保证无论正常代码如何执行,该逻辑代码也一定会执行!
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
	panic("「雪之梦技术驿站」: defer 在于保证无论正常代码如何执行,该逻辑代码也一定会执行!")
	t.Log(4)
}