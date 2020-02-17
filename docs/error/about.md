# 错误管理

`Go` 语言中有个 `defer` 关键字,常用于实现**延迟函数**来保证关键代码的最终执行,常言道: "未雨绸缪方可有备无患".

**延迟函数**就是这么一种机制,无论程序是正常返回还是异常报错,只要存在延迟函数都能保证这部分关键逻辑最终执行,所以用来做些**资源清理**等操作再合适不过了.

![go-error-about-defer.jpg](../images/go-error-about-defer.jpg)

## 出入成双有始有终

日常开发编程中,有些操作总是**成双成对**出现的,有开始就有结束,有打开就要关闭,还有一些**连续依赖**关系等等.

一般来说,我们需要控制结束语句,在**合适的位置和时机**控制结束语句,**手动保证整个程序有始有终**,不遗漏清理收尾操作.

最常见的拷贝文件操作大致流程如下:

1. 打开源文件

```go
srcFile, err := os.Open("fib.txt")
if err != nil {
	t.Error(err)
	return
}	
```

2. 创建目标文件

```go
dstFile, err := os.Create("fib.txt.bak")
if err != nil {
	t.Error(err)
	return
}
```

3. 拷贝源文件到目标文件

```go
io.Copy(dstFile, srcFile)
```

4. 关闭目标文件

```go
dstFile.Close()
srcFile.Close()
```

5. 关闭源文件

```go
srcFile.Close()
```

**值得注意的是**: 这种拷贝文件的操作需要特别注意操作顺序而且也不要忘记释放资源,比如先打开再关闭等等!

```go
func TestCopyFileWithoutDefer(t *testing.T) {
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
```

> 「雪之梦技术驿站」: 上述代码逻辑还是清晰简单的,可能不会忘记释放资源也能保证操作顺序,但是如果逻辑代码比较复杂的情况,这时候就有一定的实现难度了!

可能是为了简化类似代码的逻辑,`Go` 语言引入了 `defer` 关键字,创造了"**延迟函数**"的概念.

- 无 `defer` 的文件拷贝

```go
func TestCopyFileWithoutDefer(t *testing.T) {
	if srcFile, err := os.Open("fib.txt"); err != nil {
		t.Error(err)
		return
	} else {
		if dstFile,err := os.Create("fib.txt.bak");err != nil{
			t.Error(err)
			return
		}else{
			io.Copy(dstFile,srcFile)
	
			dstFile.Close()
			srcFile.Close()
		}
	}
}
```

- 有 `defer` 的文件拷贝

```go
func TestCopyFileWithDefer(t *testing.T) {
	if srcFile, err := os.Open("fib.txt"); err != nil {
		t.Error(err)
		return
	} else {
		defer srcFile.Close()

		if dstFile, err := os.Create("fib.txt.bak"); err != nil {
			t.Error(err)
			return
		} else {
			defer dstFile.Close()

			io.Copy(dstFile, srcFile)
		}
	}
}
```

上述示例代码简单展示了 `defer` 关键字的基本使用方式,显著的好处在于 `Open/Close` 是一对操作,不会因为写到最后而忘记 `Close` 操作,而且连续依赖时也能正常保证延迟时机.

简而言之,如果函数内部存在连续依赖关系,也就是说创建顺序是 `A->B->C` 而销毁顺序是 `C->B->A`.这时候使用 `defer` 关键字最合适不过. 

## 懒人福音延迟函数

> 官方文档相关表述见 [Defer statements](https://golang.google.cn/ref/spec#Defer_statements) 

如果没有 `defer` 延迟函数前,普通函数正常运行:

```go
func TestFuncWithoutDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 正常顺序
	t.Log("「雪之梦技术驿站」: 正常顺序")

	// 1 2
	t.Log(1)
	t.Log(2)
}
```

当添加 `defer` 关键字实现延迟后,原来的 `1` 被推迟到 `2` 后面而不是之前的 `1 2` 顺序.

```go
func TestFuncWithDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 正常顺序执行完毕后才执行 defer 代码
	t.Log(" 「雪之梦技术驿站」: 正常顺序执行完毕后才执行 defer 代码")

	// 2 1
	defer t.Log(1)
	t.Log(2)
}
```

如果存在多个 `defer` 关键字,执行顺序可想而知,越往后的越先执行,这样才能保证按照依赖顺序依次释放资源.

```go
func TestFuncWithMultipleDefer(t *testing.T) {
	// 「雪之梦技术驿站」: 猜测 defer 底层实现数据结构可能是栈,先进后出.
	t.Log(" 「雪之梦技术驿站」: 猜测 defer 底层实现数据结构可能是栈,先进后出.")

	// 3 2 1
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
}
```

相信你已经明白了多个 `defer` 语句的执行顺序,那就测试一下吧!

```go
func TestFuncWithMultipleDeferOrder(t *testing.T) {
	// 「雪之梦技术驿站」: defer 底层实现数据结构类似于栈结构,依次倒叙执行多个 defer 语句
	t.Log(" 「雪之梦技术驿站」: defer 底层实现数据结构类似于栈结构,依次倒叙执行多个 defer 语句")

	// 2 3 1
	defer t.Log(1)
	t.Log(2)
	defer t.Log(3)
}
```

初步认识了 `defer` 延迟函数的使用情况后,我们再结合文档详细解读一下相关定义.

- 英文原版文档

> A "**defer**" statement invokes a function whose execution **is deferred to** the moment the surrounding function returns,either because the **surrounding function** executed a **return** statement,reached the **end** of its function body,or because the corresponding goroutine is **panicking**.

- 中文翻译文档 

> "**defer**"语句调用一个函数,该函数的执行**被推迟到**周围函数返回的那一刻,这是因为**周围函数**执行了一个**return**语句,到达了函数体的**末尾**,或者是因为相应的协程正在**惊慌**.

具体来说,延迟函数的执行时机大概分为三种情况:

### 周围函数执行return

> because the **surrounding function** executed a **return** statement

`return` 后面的 `t.Log(4)` 语句自然是不会运行的,程序最终输出结果为 `3 2 1` 说明了 `defer` 语句会在周围函数执行 `return` 前依次逆序执行.

```go
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
```

### 周围函数到达函数体

> reached the **end** of its function body

周围函数的函数体运行到结尾前逆序执行多个 `defer` 语句,即先输出 `3` 后依次输出 `2 1`.
最终函数的输出结果是 `3 2 1` ,也就说是没有 `return` 声明也能保证结束前执行完 `defer` 延迟函数.

```go
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
```

### 当前协程正惊慌失措

> because the corresponding goroutine is **panicking**

周围函数万一发生 `panic` 时也会先运行前面已经定义好的 `defer` 语句,而 `panic` 后续代码因为没有特殊处理,所以程序崩溃了也就无法运行.

函数的最终输出结果是 `3 2 1 panic` ,如此看来 `defer` 延迟函数还是非常尽忠职守的,虽然心里很慌但还是能保证老弱病残先行撤退!

```go
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
```

通过解读 `defer` 延迟函数的定义以及相关示例,相信已经讲清楚什么是 `defer` 延迟函数了吧?

简单地说,延迟函数就是一种**未雨绸缪**的规划机制,帮助开发者编程程序时**及时做好收尾善后工作**,提前做好预案以准备随时应对各种情况.

- 当周围函数正常执行到到达函数体结尾时,如果发现存在延迟函数自然会逆序执行延迟函数.
- 当周围函数正常执行遇到return语句准备返回给调用者时,存在延迟函数时也会执行,同样满足善后清理的需求.
- 当周围函数异常运行不小心 `panic` 惊慌失措时,程序存在延迟函数也不会忘记执行,提前做好预案发挥了作用.

所以不论是正常运行还是异常运行,**提前做好预案总是没错的**,基本上可以保证万无一失,所以不妨考虑考虑 `defer` 延迟函数?

![go-error-about-lovely.png](../images/go-error-about-lovely.png)

## 延迟函数应用场景

基本上**成双成对**的操作都可以使用延迟函数,**尤其是申请的资源前后存在依赖关系时**更应该使用 `defer` 关键字来简化处理逻辑.

下面举两个常见例子来说明延迟函数的应用场景.

- Open/Close

文件操作一般会涉及到打开和开闭操作,尤其是文件之间拷贝操作更是有着严格的顺序,只需要按照申请资源的顺序紧跟着`defer` 就可以满足资源释放操作.

```go
func readFileWithDefer(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
```

- Lock/Unlock

锁的申请和释放是保证同步的一种重要机制,需要申请多个锁资源时可能存在依赖关系,不妨尝试一下延迟函数!

```go
var mu sync.Mutex
var m = make(map[string]int)
func lookupWithDefer(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
```

## 总结以及下节预告

`defer` 延迟函数是保障关键逻辑正常运行的一种机制,如果存在多个延迟函数的话,一般会按照**逆序**的顺序运行,类似于栈结构.

延迟函数的运行时机一般有三种情况:

- 周围函数遇到返回时

```go
func funcWithMultipleDeferAndReturn() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}
```

- 周围函数函数体结尾处

```go
func funcWithMultipleDeferAndEnd() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
```

- 当前协程惊慌失措中

```go
func funcWithMultipleDeferAndPanic() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("panic")
	fmt.Println(4)
}
```

本文主要介绍了什么是 `defer` 延迟函数,通过解读官方文档并配套相关代码认识了延迟函数,但是延迟函数中存在一些可能令人比较迷惑的地方.

![go-error-about-question.png](../images/go-error-about-question.png)

读者不妨看一下下面的代码,将心里的猜想和实际运行结果比较一下,我们下次再接着分享,感谢你的阅读.

```go
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
```

## 延伸阅读参考文档

- [Defer_statements](https://golang.google.cn/ref/spec#Defer_statements)
- [go语言的defer语句](https://www.jianshu.com/p/5b0b36f398a2)
- [Go defer实现原理剖析](https://studygolang.com/articles/16067)
- [go语言 defer 你不知道的秘密!](https://www.cnblogs.com/baizx/p/5024547.html)
- [Go语言中defer的一些坑](https://www.jianshu.com/p/79c029c0bd58)
- [go defer (go延迟函数)](https://www.cnblogs.com/ysherlock/p/8150726.html)
