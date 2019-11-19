package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"
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

func funcWithMultipleDeferAndReturn() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}

func TestFuncWithMultipleDeferAndReturn(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数会在包围函数正常return之前逆序执行.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数会在包围函数正常return之前逆序执行.")

	// 3 2 1
	funcWithMultipleDeferAndReturn()
}

func funcWithMultipleDeferAndEnd() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func TestFuncWithMultipleDeferAndEnd(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数会在包围函数到达函数体结尾之前逆序执行.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数会在包围函数到达函数体结尾之前逆序执行.")

	// 3 2 1
	funcWithMultipleDeferAndEnd()
}

func funcWithMultipleDeferAndPanic() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("panic")
	fmt.Println(4)
}

func TestFuncWithMultipleDeferAndPanic(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数会在包围函数panic惊慌失措之前逆序执行.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数会在包围函数panic惊慌失措之前逆序执行.")

	// 3 2 1
	funcWithMultipleDeferAndPanic()
}

func readFileWithDefer(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

var mu sync.Mutex
var m = make(map[string]int)
func lookupWithDefer(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func funcCallWithDefer() {
	fmt.Println("funcInvokeWithDefer function is called")
}

func TestFuncCallWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: defer 语句可以是函数调用.
	fmt.Println(" 「雪之梦技术驿站」: defer 语句可以是函数调用.")

	defer funcCallWithDefer()

	fmt.Println("TestFuncInvokeWithDefer function call has ended")
}

type Lang struct {
	name    string
	website string
}

func (l *Lang) ToString() {
	fmt.Printf("Lang:[name = %s,website = %s] \n", l.name, l.website)
}

func TestMethodCallWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: defer 语句也可以是方法调用.
	fmt.Println(" 「雪之梦技术驿站」: defer 语句也可以是方法调用.")

	var l = new(Lang)
	l.name = "Go"
	l.website = "https://snowdreams1006.github.io/go/"

	defer l.ToString()

	fmt.Println("TestMethodCallWithDefer method call has ended")
}

func TestParenthesizedCallWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: defer 语句不可以被括号包裹.
	fmt.Println(" 「雪之梦技术驿站」: defer 语句不可以被括号包裹.")

	// function must be invoked in defer statement
	//defer (fmt.Println("it cannot be parenthesized."))

	fmt.Println("TestParenthesizedCallWithDefer statement call has ended")
}

func TestBuiltinFuncCallWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: defer 语句不可以被括号包裹.
	fmt.Println(" 「雪之梦技术驿站」: defer 语句不可以被括号包裹.")

	arr := new([10]int)
	arr[4] = 5
	arr[7] = 8

	// defer discards result of len(arr)
	//defer len(arr)
	defer println("Calls of built-in functions are restricted as for expression statements.")

	fmt.Println("TestBuiltinFuncCallWithDefer function call has ended")
}

func trace(funcName string) func(){
	start := time.Now()
	fmt.Printf("function %s enter at %s \n",funcName,start)

	return func(){
		fmt.Printf("function %s exit at %s(elapsed %s)",funcName,time.Now(),time.Since(start))
	}
}

func foo(){
	fmt.Printf("foo begin at %s \n",time.Now())

	defer trace("foo")()
	time.Sleep(5*time.Second)

	fmt.Printf("foo end at %s \n",time.Now())
}

func TestFoo(t *testing.T) {
	foo()
}

func surroundingFuncEvaluatedNotInvoked(init int) int {
	fmt.Printf("1.init=%d\n",init)

	defer func() {
		fmt.Printf("2.init=%d\n",init)

		init ++

		fmt.Printf("3.init=%d\n",init)
	}()

	fmt.Printf("4.init=%d\n",init)

	return init
}

func TestSurroundingFuncEvaluatedNotInvoked(t *testing.T) {
	// 「雪之梦技术驿站」: 1 4 2 3
	fmt.Println(" 「雪之梦技术驿站」: 普通函数顺序执行,结果很明显,不需要解释.")

	retVal := surroundingFuncEvaluatedNotInvoked(0)
	fmt.Printf("retVal = %d\n", retVal)
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

func deferFuncWithAnonymousReturnValue() int {
	var retVal int
	defer func() {
		retVal++
	}()
	return 0
}

func deferFuncWithNamedReturnValue() (retVal int) {
	defer func() {
		retVal++
	}()
	return 0
}

func TestDeferFuncWhenReturn(t *testing.T) {
	t.Log(deferFuncWithAnonymousReturnValue())
	t.Log(deferFuncWithNamedReturnValue())
}


//func deferFuncOrderWhenReturn() (result int) {
//	defer func() {
//		// 2. before : result = 0
//		fmt.Printf("before : result = %v\n", result)
//
//		result++
//
//		// 3. after : result = 1
//		fmt.Printf("after : result = %v\n", result)
//	}()
//
//	// 1. return : result = 0
//	fmt.Printf("return : result = %v\n", result)
//
//	return 0
//}
//
//func TestDeferFuncOrderWhenReturn(t *testing.T) {
//	// 「雪之梦技术驿站」: 包围函数具有显式返回语句时,延迟函数defer在结果参数赋值之后且在函数返回之前执行.
//	t.Log(" 「雪之梦技术驿站」: 包围函数具有显式返回语句时,延迟函数defer在结果参数赋值之后且在函数返回之前执行.")
//
//	// 4. result = 1
//	result := deferFuncOrderWhenReturn()
//	t.Logf("result = %v", result)
//}

func deferFuncOrderWhenReturnExplain() (result int) {
	result = 0

	func() {
		// 1. before : result = 0
		fmt.Printf("before : result = %v\n", result)

		result++

		// 2. after : result = 1
		fmt.Printf("after : result = %v\n", result)
	}()

	// 3. return : result = 0
	fmt.Printf("return : result = %v\n", result)

	return
}

func TestDeferFuncOrderWhenReturnExplain(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数中的包围函数具有显式返回语句时,return result 并不是原子操作,整个函数被划分为三部分,首先赋值结果变量,接着执行延迟函数,最后执行return.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数中的包围函数具有显式返回语句时,return result 并不是原子操作,整个函数被划分为三部分,首先赋值结果变量,接着执行延迟函数,最后执行return.")

	// 4. result = 1
	result := deferFuncOrderWhenReturnExplain()
	t.Logf("result = %v", result)
}

func deferFuncOrderWhenReturnDemo() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}

func TestDeferFuncOrderWhenReturnDemo(t *testing.T) {
	// 「雪之梦技术驿站」: 注意 defer 延迟函数中可能会访问并修改结果变量,最终结果是42并不是6.
	t.Log(" 「雪之梦技术驿站」: 注意 defer 延迟函数中可能会访问并修改结果变量,最终结果是42并不是6.")

	// 42
	t.Log(deferFuncOrderWhenReturnDemo())
}

func deferFuncWithRedundantVarWhenReturn() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func TestDeferFuncWithRedundantVarWhenReturn(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数修改过的变量并不一定是结果变量哟!
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数修改过的变量并不一定是结果变量哟!")

	// 5
	t.Log(deferFuncWithRedundantVarWhenReturn())
}

func deferFuncWithRedundantVarWhenReturnExplain() (r int) {
	t := 5

	// 1.结果变量赋值语句
	r = t
	func() {
		t += 5
	}()

	// 2.直接返回
	return
}

func TestDeferFuncWithRedundantVarWhenReturnExplain(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数修改过的变量不是结果变量并不会受到延迟函数的影响.
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数修改过的变量不是结果变量并不会受到延迟函数的影响.")

	// 5
	t.Log(deferFuncWithRedundantVarWhenReturnExplain())
}

func deferFuncWithClosureWhenReturn() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 5
}

func TestDeferFuncWithClosureWhenReturn(t *testing.T) {
	// 「雪之梦技术驿站」: defer 延迟函数变量传递后被修改,最终结果变量也不会受影响,因为是值传递!
	t.Log(" 「雪之梦技术驿站」: defer 延迟函数变量传递后被修改,最终结果变量也不会受影响,因为是值传递!")

	// 5
	t.Log(deferFuncWithClosureWhenReturn())
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

func TestCopyFileWithoutDefer(t *testing.T) {
	//if srcFile, err := os.Open("fib.txt"); err != nil {
	//	t.Error(err)
	//	return
	//} else {
	//	if dstFile,err := os.Create("fib.txt.bak");err != nil{
	//		t.Error(err)
	//		return
	//	}else{
	//		io.Copy(dstFile,srcFile)
	//
	//		dstFile.Close()
	//		srcFile.Close()
	//	}
	//}

	srcFile, err := os.Open("fib.txt")
	if err != nil {
		t.Error(err)
		return
	}

	dstFile, err := os.Create("fib.txt.bak")
	if err != nil {
		t.Error(err)
		return
	}

	io.Copy(dstFile, srcFile)

	dstFile.Close()
	srcFile.Close()
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
