package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestFuncWithoutDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 正常顺序
	t.Log("「雪之梦技术驿站」: 正常顺序")

	// 1 2
	t.Log(1)
	t.Log(2)
}

func TestFuncWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 正常顺序执行完毕后才执行 defer 代码
	t.Log(" 「雪之梦技术驿站」: 正常顺序执行完毕后才执行 defer 代码")

	// 2 1
	defer t.Log(1)
	t.Log(2)
}

func TestFuncWithMultipleDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 猜测 defer 底层实现数据结构可能是栈,先进后出.
	t.Log(" 「雪之梦技术驿站」: 猜测 defer 底层实现数据结构可能是栈,先进后出.")

	// 3 2 1
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
}

func TestFuncWithMultipleDeferOrder(t *testing.T) {
	// 「雪之梦技术驿站」: defer 底层实现数据结构类似于栈结构,依次倒叙执行多个 defer 语句
	t.Log(" 「雪之梦技术驿站」: defer 底层实现数据结构类似于栈结构,依次倒叙执行多个 defer 语句")

	// 2 3 1
	defer t.Log(1)
	t.Log(2)
	defer t.Log(3)
}

func TestFuncWithMultipleDeferAndReturn(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数会在包围函数正常return之前逆序执行.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数会在包围函数正常return之前逆序执行.")

	// 3 2 1
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
	return
	t.Log(4)
}

func TestFuncWithMultipleDeferAndPanic(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数会在包围函数panic惊慌失措之前逆序执行.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数会在包围函数panic惊慌失措之前逆序执行.")

	// 3 2 1
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
	panic("「雪之梦技术驿站」: defer 延迟函数会在包围函数panic惊慌失措之前逆序执行.")
	t.Log(4)
}

func noDeferFuncOrderWhenReturn() (result int) {
	func() {
		// 1. before : result = 0
		fmt.Printf("before : result = %v\n", result)

		result++

		// 2. after : result = 1
		fmt.Printf("after : result = %v\n", result)
	}()

	// 3. return : result = 1
	fmt.Printf("return : result = %v\n", result)

	return 0
}

func TestNoDeferFuncOrderWhenReturn(t *testing.T) {
	// 「雪之梦技术驿站」: 普通函数顺序执行,结果很明显,不需要解释.
	t.Log(" 「雪之梦技术驿站」: 普通函数顺序执行,结果很明显,不需要解释.")

	// 4. result = 0
	result := noDeferFuncOrderWhenReturn()
	t.Logf("result = %v", result)
}

func deferFuncOrderWhenReturn() (result int) {
	defer func() {
		// 2. before : result = 0
		fmt.Printf("before : result = %v\n", result)

		result++

		// 3. after : result = 1
		fmt.Printf("after : result = %v\n", result)
	}()

	// 1. return : result = 0
	fmt.Printf("return : result = %v\n", result)

	return 0
}

func TestDeferFuncOrderWhenReturn(t *testing.T) {
	// 「雪之梦技术驿站」: 包围函数具有显式返回语句时,延迟函数defer在结果参数赋值之后且在函数返回之前执行
	t.Log(" 「雪之梦技术驿站」: 普通函数顺序执行,结果很明显,不需要解释.")

	// 4. result = 1
	result := deferFuncOrderWhenReturn()
	t.Logf("result = %v", result)
}


func f11() (result int) {
	result = 0 //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	func() { //defer被插入到return之前执行，也就是赋返回值和ret指令之间
		result++
	}()
	return
}

func TestF11(t *testing.T) {
	// 1
	t.Log(f11())
}

func f111() (result int) {
	func() {
		result++
	}()
	return 0
}

func TestF111(t *testing.T) {
	// 0
	t.Log(f111())
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func TestF2(t *testing.T) {
	// 5
	t.Log(f2())
}

func f22() (r int) {
	t := 5
	r = t //赋值指令
	func() { //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		t = t + 5
	}()
	return //空的return指令
}

func TestF22(t *testing.T) {
	// 5
	t.Log(f22())
}

func f222() (r int) {
	t := 5
	func() {
		t = t + 5
	}()
	return t
}

func TestF222(t *testing.T) {
	// 10
	t.Log(f222())
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func TestF3(t *testing.T) {
	// 1
	t.Log(f3())
}

func f33() (r int) {
	r = 1 //给返回值赋值
	func(r int) { //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return //空的return
}

func TestF33(t *testing.T) {
	// 1
	t.Log(f33())
}

func f333() (r int) {
	func(r int) {
		r = r + 5
	}(r)
	return 1
}

func TestF333(t *testing.T) {
	// 1
	t.Log(f3())
}

// f returns 42
func f() (result int) {
	defer func() {
		fmt.Println("before ", result)

		// result is accessed after it was set to 6 by the return statement
		result *= 7

		fmt.Println("after ", result)
	}()

	fmt.Println("return ", result)

	return 6
}

func TestF(t *testing.T) {
	// 1
	t.Log(f())
}

func fibonacciWithClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestWriteFileWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 其实不用关心 defer 调用顺序,成对操作时记得defer结束即可!
	t.Log(" 「雪之梦技术驿站」: 其实不用关心 defer 调用顺序,成对操作时记得defer结束即可!")

	// 「雪之梦技术驿站」: 有始有终,打开过文件要及时关闭,defer 让我们操作变得有始有终!
	if file, err := os.Create("fib.txt"); err != nil {
		panic(err)
	} else {
		defer file.Close()

		// 「雪之梦技术驿站」: bufio 暂存内存要刷新到文件,defer 帮助我们及时进行资源管理,出入成双!
		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fibonacciWithClosure()
		for i := 0; i < 10; i++ {
			fmt.Fprintln(writer, f())
		}
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
	//「雪之梦技术驿站」: panic 报错后程序已崩溃,后续程序不再执行!
	t.Log("「雪之梦技术驿站」: panic 报错后程序已崩溃,后续程序不再执行!")

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
}

func TestWriteFileErrorWithoutPanic(t *testing.T) {
	//「雪之梦技术驿站」: 一般应该捕获而不是直接抛出panic,后续程序可以执行!
	t.Log("「雪之梦技术驿站」: 一般应该捕获而不是直接抛出panic,后续程序可以执行!")

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
}

func TestWriteFileErrorWithoutPanicAndExactError(t *testing.T) {
	//「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.
	t.Log("「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.")

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
}

func TestWriteFileErrorWithoutPanicAndCustomError(t *testing.T) {
	//「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.
	t.Log("「雪之梦技术驿站」: 明确 error 类型的前提下,可以针对性处理,否则要么捕获错误信息要么直接 panic 错误.")

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
}

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [0,100]")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if fib, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(fib)
	}
}

func TestPanicVxExit1(t *testing.T) {
	t.Log("start")
	os.Exit(-1)
}

func TestPanicVxExit2(t *testing.T) {
	defer func() {
		t.Log("finally")
	}()

	t.Log("start")
	panic(errors.New("something wrong"))
	os.Exit(-1)
}

func TestPanicVxExit3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("recover from %v", err)
		}
	}()

	t.Log("start")
	panic(errors.New("something wrong"))
}
