package error

import (
	"bufio"
	"errors"
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

		//「雪之梦技术驿站」: 其实不用关心 defer 调用顺序,因为开始是顺序,而结束自然就是逆序.
		t.Log("「雪之梦技术驿站」: 其实不用关心 defer 调用顺序,因为开始是顺序,而结束自然就是逆序.")
	}
}

func TestCalculateWithDefer(t *testing.T) {
	//「雪之梦技术驿站」: 参与在defer语言时计算
	for i := 0; i < 10; i++ {
		// 5 4 3 2 1 0
		defer t.Log(i)

		if i == 5 {
			panic("「雪之梦技术驿站」: 参与在defer语言时计算")
		}
	}
}

func TestWriteFileErrorWithPanic(t *testing.T) {
	// 「雪之梦技术驿站」: 故意报错演示异常信息,panic 报错后程序已崩溃,后续程序不再执行!
	if file, err := os.OpenFile("fib.txt", os.O_EXCL|os.O_CREATE, 0666); err != nil {
		// panic: open fib.txt: file exists
		panic(err)
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fibonacciWithClosure()
		for i := 0; i < 10; i++ {
			fmt.Fprintln(writer, f())
		}
	}

	//「雪之梦技术驿站」: panic 报错后程序已崩溃,后续程序不再执行!
	t.Log("「雪之梦技术驿站」: panic 报错后程序已崩溃,后续程序不再执行!")
}

func TestWriteFileErrorWithoutPanic(t *testing.T) {
	// 「雪之梦技术驿站」: 故意报错演示异常信息,一般应该捕获而不是直接抛出panic,后续程序可以执行!
	if file, err := os.OpenFile("fib.txt", os.O_EXCL|os.O_CREATE, 0666); err != nil {
		// occur error with  'open fib.txt: file exists'
		t.Logf("occur error with  '%s'", err.Error())
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fibonacciWithClosure()
		for i := 0; i < 10; i++ {
			fmt.Fprintln(writer, f())
		}
	}

	//「雪之梦技术驿站」: 一般应该捕获而不是直接抛出panic,后续程序可以执行!
	t.Log("「雪之梦技术驿站」: 一般应该捕获而不是直接抛出panic,后续程序可以执行!")
}

func TestWriteFileErrorWithoutPanicAndExactError(t *testing.T) {
	// 「雪之梦技术驿站」: 故意报错演示异常信息,断言已知 error 进行针对性处理,无法处理时直接 panic 或者捕获错误信息.
	if file, err := os.OpenFile("fib.txt", os.O_EXCL|os.O_CREATE, 0666); err != nil {
		// operate = open,path = fib.txt,err = file exists
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			t.Logf("operate = %s,path = %s,err = %s", pathErr.Op, pathErr.Path, pathErr.Err)
		}
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fibonacciWithClosure()
		for i := 0; i < 10; i++ {
			fmt.Fprintln(writer, f())
		}
	}

	//「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.
	t.Log("「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.")
}

func TestWriteFileErrorWithoutPanicAndCustomError(t *testing.T) {
	// 「雪之梦技术驿站」: 故意报错演示异常信息,断言已知 error 进行针对性处理,无法处理时直接 panic 或者捕获错误信息.
	if file, err := os.OpenFile("fib.txt", os.O_EXCL|os.O_CREATE, 0666); err != nil {
		err = errors.New("「雪之梦技术驿站」: 自定义 error 错误信息")

		//「雪之梦技术驿站」: 自定义 error 错误信息
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			t.Logf("operate = %s,path = %s,err = %s", pathErr.Op, pathErr.Path, pathErr.Err)
		}
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fibonacciWithClosure()
		for i := 0; i < 10; i++ {
			fmt.Fprintln(writer, f())
		}
	}

	//「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.
	t.Log("「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.")
}