package error

import (
	"bufio"
	"fmt"
	"os"
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

func fibonacciWithClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestWriteFileWithDefer(t *testing.T) {
	//「雪之梦技术驿站」: 有始有终,打开过文件要及时关闭,defer 让我们操作变得有始有终!
	if file, err := os.Create("fib.txt"); err != nil {
		panic(err)
	} else {
		defer file.Close()

		//「雪之梦技术驿站」: bufio 暂存内存要刷新到文件,defer 帮助我们及时进行资源管理,出入成双!
		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fibonacciWithClosure()
		for i := 0; i < 10; i++ {
			fmt.Fprintln(writer, f())
		}
	}
}
