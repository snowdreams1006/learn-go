# 斐波那契数列

本篇文章是 `Go` 语言学习笔记之**函数式编程**系列文章的第二篇,上一篇介绍了函数基础,这一篇文章重点介绍函数的重要应用之一: **闭包**

空谈误国,实干兴邦,以具体代码示例为基础讲解**什么是闭包**以及为什么需要闭包等问题,下面我们沿用上篇文章的示例代码开始本文的学习吧!

> **斐波那契数列**是形如 `1 1 2 3 5 8 13 21 34 55` 的**递增**数列,即从第三个数开始,后一个数字是**前两个数字之和**,保持此规律无限递增...

![go-functional-programming-about-fib.png](../images/go-functional-programming-about-fib.png)

开门见山,直接给出**斐波那契数列生成器**,接下来的文章慢慢深挖背后隐藏的奥秘,一个例子讲清楚什么是闭包.

> 「雪之梦技术驿站」: 如果还不了解 `Go` 语言的函数用法,可以参考上一篇文章: [go 学习笔记之学习函数式编程前不要忘了函数基础](https://mp.weixin.qq.com/s/dprkCOvPZHr6fi_qC91dVw)

- `Go` 版本的斐波那契数列生成器

```go
// 1 1 2 3 5 8 13 21 34 55
//     a b
//       a b
func fibonacci() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
```

> 「雪之梦技术驿站」: `Go` 语言支持**连续赋值**,更加贴合思考方式,而其余主流的编程语言可能**不支持**这种方式,大多采用**临时变量**暂存的方式.

- `Go` 版本的单元测试用例

```go
// 1 1 2 3 5 8 13 21 34 55
func TestFibonacci(t *testing.T) {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Print(f(), " ")
  }
  fmt.Println()
}
```

> 「雪之梦技术驿站」: 循环调用 `10` 次**斐波那契数列生成器**,因此生成**前十位**数列: `1 1 2 3 5 8 13 21 34 55`

## 背后有故事

小小的**斐波那契数列生成器**背后蕴藏着丰富的 `Go` **语言特性**,该示例也是**官方示例**之一.

![go-functional-programming-fib-try-go.png](../images/go-functional-programming-fib-try-go.png)

- 支持**连续赋值**,无需中间变量

> 「雪之梦技术驿站」: `Go` 语言和其他主流的编程语言不同,它们大多数最多支持多变量的**连续声明**而不支持**连续赋值**.

这也是 `Go` 语言特有的交换变量方式,`a, b = b, a` 语义简单明确并不用引入额外的临时变量.

```go
func TestExchange(t *testing.T) {
  a, b := 1, 2
  t.Log(a,b)

  // 2,1
  a, b = b, a
  t.Log(a,b)
}
```

> 「雪之梦技术驿站」: `Go` 语言实现变量交互的示例,`a, b = b, a` 表示变量直接交换.

而其他主流的编程语言的惯用做法是需要引入临时变量,大多数代码类似如下方式:

```go
func TestExchange(t *testing.T) {
  a, b := 1, 2
  t.Log(a,b)

  // 2,1
  temp := a
  a = b
  b = temp
  t.Log(a,b)
}
```

> 「雪之梦技术驿站」: `Go` 语言的**多变量同时赋值**特性体现的更多是一种**声明式编程**思想,不关注具体实现,而引入临时变量这种体现的则是**命令式编程**思维.

- 函数的**返回值也可以是函数**

> 「雪之梦技术驿站」: `Go` 语言中的函数是**一等公民**,不仅函数的返回值可以是**函数,参数,变量**等等都可以是函数.

函数的**返回值可以是函数**,这样的实际意义在于**使用者**可以拥有更大的灵活性,有时可以用作**延迟计算**,有时也可以用作**函数增强**.

先来演示一下**延迟计算**的示例:

函数的返回值可以是函数,由此实现类似于 `i++` 效果的**自增函数**.因为 `i` 的初值是 `0`,所以每调用一次该函数, `i` 的值就会自增,从而实现 `i++` 的效果.

```go
func autoIncrease() func() int {
  i := 0
  return func() int {
    i = i + 1
    return i
  }
}
```

再小的代码片段也不应该忘记测试,单元测试继续走起,顺便看一下使用方法.

```go
func TestAutoIncrease(t *testing.T) {
  a := autoIncrease()

  // 1 2 3
  t.Log(a(),a(),a())
}
```

初始调用 `autoIncrease` 函数并没有直接得到结果而是**返回了函数引用**,等到使用者觉得**时机成熟**后再次调用返回的函数引用即变量`a` ,这时候才会真正计算结果,这种方式被称为**延迟计算也叫做惰性求值**.

继续演示一下**功能增强**的示例:

因为要演示函数增强功能,没有输入哪来的输出?

所以函数的入参应该也是函数,返回值就是增强后的函数.

这样的话接下来要做的函数就比较清晰了,这里我们定义 `timeSpend` 函数: 实现的功能是**包装特定类型的函数**,增加**计算函数运行时间**的新功能并包装成函数,最后**返回**出去给使用者.

```go
func timeSpend(fn func()) func() {
  return func()  {
    start := time.Now()

    fn()

    fmt.Println("time spend : ", time.Since(start).Seconds())
  }
}
```

为了演示包装函数 `timeSpend`,需要定义一个比较**耗时函数当做入参**,函数名称姑且称之为为 `slowFunc` ,睡眠 `1s` 来**模拟耗时**操作.

```go
func slowFunc() {
  time.Sleep(time.Second * 1)

  fmt.Println("I am slowFunc")
}
```

**无测试不编码**,继续运行单元测试用例,演示包装函数 `timeSpend` 是如何增强原函数 `slowFunc` 以实现功能增强?

```go
func TestSlowFuncTimeSpend(t *testing.T) {
  slowFuncTimeSpend := timeSpend(slowFunc)

  // I am slowFunc
  // time spend :  1.002530902
  slowFuncTimeSpend()
}
```

> 「雪之梦技术驿站」: 测试结果显示原函数 `slowFunc` 被当做入参传递给包装函数 `timeSpend` 后实现了功能增强,不仅保留了原函数功能还增加了计时功能.

- 函数嵌套可能是闭包函数

不论是引言部分的**斐波那契数列生成器**函数还是演示函数返回值的**自增函数**示例,其实这种形式的函数有一种专业术语称为"闭包".

一般而言,函数内部不仅**存在变量还有嵌套函数**,而嵌套函数又引用了这些外部变量的话,那么这种形式**很有可能**就是**闭包函数**.

## 什么是闭包

如果有一句话介绍什么是闭包,那么我更愿意称其为**流浪在外的人想念爸爸妈妈**!

![go-functional-programming-fib-go-home.jpg](../images/go-functional-programming-fib-go-home.jpg)

如果非要用比较官方的定义去**解释什么是闭**包,那只好翻开[维基百科](https://en.wikipedia.org/wiki/Closure_%28computer_programming%29) 看下有关闭包的定义: 

> In programming languages, a closure, also lexical closure or function closure, is a technique for implementing lexically scoped name binding in a language with first-class functions. Operationally, a closure is a record storing a function[a] together with an environment.[1] The environment is a mapping associating each free variable of the function (variables that are used locally, but defined in an enclosing scope) with the value or reference to which the name was bound when the closure was created.[b] Unlike a plain function, a closure allows the function to access those captured variables through the closure's copies of their values or references, even when the function is invoked outside their scope.

如果能够**直接理解英文**的同学可以略过这部分的中文翻译,要是不愿意费脑理解英文的小伙伴跟我一起**解读中文**吧!

### 闭包是一种技术

第一句话英文如下:

> In programming languages, a **closure**, also **lexical closure** or **function closure**, is a **technique** for implementing lexically scoped **name binding** in a language with **first-class functions**.

相应的中文翻译:

> **闭包**也叫做**词法闭包**或者**函数闭包**,是**函数优先**编程语言中用于实现词法范围的**名称绑定**技术.

概念性定义解释后可能还是不太清楚,那么就用代码解释一下什么是闭包?

> 「雪之梦技术驿站」: 编程语言千万种,前端后端和中台;心有余而力不足,大众化 `Js` 带上 `Go` .

- `Go` 实现斐波那契数列生成器

这是开篇引言的示例,直接照搬过来,这里主要说明 `Go` 支持闭包这种技术而已,所以并不关心具体实现细节.

```go
func fibonacci() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
```

单元测试用例函数,连续 `10` 次调用**斐波那契数列生成器**,输出斐波那契数列中的前十位数字.

```go
// 1 1 2 3 5 8 13 21 34 55
func TestFibonacci(t *testing.T) {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Print(f(), " ")
  }
  fmt.Println()
}
```

- `Js` 实现斐波那契数列生成器

仿照 `Go` 语言的实现方式写一个 `Js` 版本的**斐波那契数列生成器**,相关代码如下:

```js
function fibonacci() {
  var a, b;
  a = 0;
  b = 1;
  return function() {
    var temp = a;
    a = b;
    b = temp + b;
    return a;
  }
}
```

同样的,仿造测试代码写出 `Js` 版本的测试用例:

```js
// 1 1 2 3 5 8 13 21 34 55
function TestFibonacci() {
  var f = fibonacci();
  for(var i = 0; i < 10; i++ ){
    console.log(f() +" ");
  }
  console.log();
}
```

不仅仅是 `Js` 和 `Go` 这两种编程语言能够实现闭包,实际上很多编程语言都能实现闭包,就像是**面向对象编程**一样,也**不是某种语言专有的技术**,唯一的区别可能就是**语法细节上略有不同**吧,所以记住了: **闭包是一种技术**!

### 闭包存储了环境

第二句英文如下:

> Operationally, a closure is a **record** storing a **function**[a] together with an **environment**.

相应的中文翻译:

> 在操作上,闭包是将**函数[a]与环境**一起存储的**记录**。

第一句我们知道了**闭包是一种技术**,而现在我们有知道了闭包**存储了闭包函数所需要的环境**,而环境分为函数运行时所处的**内部环境**和依赖的**外部环境**,闭包函数被**使用者调用**时不会像普通函数那样丢失环境而是存储了环境.

如果是普通函数方式打开上述示例的**斐波那契数列生成器**:

```go
func fibonacciWithoutClosure() int {
  a, b := 0, 1
  a, b = b, a+b
  return a
}
```

可想而知,这样肯定是不行的,因为函数内部环境是无法维持的,使用者每次调用 `fibonacciWithoutClosure` 函数都会重新初始化变量 `a,b` 的值,因而**无法实现累加自增效果**.

```go
// 1 1 1 1 1 1 1 1 1 1 
func TestFibonacciWithoutClosure(t *testing.T) {
  for i := 0; i < 10; i++ {
    fmt.Print(fibonacciWithoutClosure(), " ")
  }
  fmt.Println()
}
```

很显然,函数内部定义的变量每次运行函数时都会重新初始化,为了避免这种情况,在不改变整体实现思路的前提下,只需要**提升变量的作用范围**就能实现**斐波那契数列生成器**函数:

```go
var a, b = 0, 1
func fibonacciWithoutClosure() int {
  a, b = b, a+b
  return a
}
```

此时再次运行 `10` 次**斐波那契数列生成器**函数,如我们所愿生成前 `10` 位斐波那契数列.

```go
// 1 1 2 3 5 8 13 21 34 55
func TestFibonacciWithoutClosure(t *testing.T) {
  for i := 0; i < 10; i++ {
    fmt.Print(fibonacciWithoutClosure(), " ")
  }
  fmt.Println()
}
```

所以说普通函数 `fibonacciWithoutClosure` 的运行环境要么是仅仅依赖内部变量维持的**独立环境**,每次运行都会**重新初始化**,无法实现变量的重复利用;要么是依赖了外部变量维持的具有**记忆功能的环境**,解决了重新初始化问题的同时引入了新的问题,那就是必须**定义作用范围更大的外部环境**,增加了维护成本.

既然函数内的变量无法维持而函数外的变量又需要管理,如果能两者结合的话,岂不是皆大欢喜,扬长补短?

![go-functional-programming-fib-balance.jpg](../images/go-functional-programming-fib-balance.jpg)

对的,闭包基本上就是这种实现思路!

```go
func fibonacci() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
```

**斐波那契数列生成器**函数 `fibonacci` 的**返回值是匿名函**数,而匿名函数的返回值就是斐波那契数字.

如果不考虑函数内部实现细节,整个函数的语义是十分明确的,**使用者**初始化调用 `fibonacci` 函数时得到**返回值是真正的斐波那契生成器函数**,用变量暂存起来,当需要生成斐波那契数字的时候再**调用刚才暂存的变量**就能真正生成斐波那契数列.

```go
// 1 1 2 3 5 8 13 21 34 55
func TestFibonacci(t *testing.T) {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Print(f(), " ")
  }
  fmt.Println()
}
```

现在我们再好好比较一下这种形式实现的闭包和普通函数的区别?

- 闭包函数 `fibonacci` 的**内部定义**了变量 `a,b`,最终返回的匿名函数中使用了变量 `a,b`,使用时**间接生成**斐波那契数字.
- 普通函数 `fibonacciWithoutClosure` 的**外部定义**了变量 `a,b`,调用该函数**直接生成**斐波那契数字.
- 闭包函数是**延迟计算也就是惰性求值**而普通函数是**立即计算**,两者的调用方式不一样.

但是如果把视角切换到**真正有价值部分**,你会发现闭包函数只不过是**普通函数嵌套**而已!

```go
func fibonacciDeduction() func() int {
  a, b := 0, 1

  func fibonacciGenerator() int {
    a, b = b, a+b
    return a
  }

  return fibonacciGenerator
}
```

只不过 `Go` 并**不支持函数嵌套**,只能使用匿名函数来实现函数嵌套的效果,所以上述示例是会直接报错的哟!

![go-functional-programming-fib-nested-error.png](../images/go-functional-programming-fib-nested-error.png)

但是某些语言是支持函数嵌套的,比如最常用的 `Js` 语言就支持函数嵌套,用 `Js` 重写上述代码如下:

```js
function fibonacciDeduction() {
  var a, b;
  a = 0;
  b = 1;

  function fibonacciGenerator() {
    var temp = a;
    a = b;
    b = temp + b;
    return a
  }

  return fibonacciGenerator
}
```

**斐波那契数列生成器**函数是 `fibonacciDeduction`,该函数内部真正实现生成器功能的却是 `fibonacciGenerator` 函数,正是这个函数使用了变量 `a,b` ,相当于把**外部变量打包绑定成运行环境的一部分**!

```js
// 1 1 2 3 5 8 13 21 34 55
function TestFibonacciDeduction() {
  var f = fibonacciDeduction();
  for(var i = 0; i < 10; i++ ){
    console.log(f() +" ");
  }
  console.log();
}
```

> 「雪之梦技术驿站」: 闭包并不是某一种语言特有的技术,虽然各个语言的实现细节上有所差异,但并不妨碍整体理解,正如定义的第二句那样: `storing a **function**[a] together with an **environment**.`

### 环境关联了自由变量

第三句英文如下:

> The environment is a mapping associating each **free variable** of the function (variables that are used locally, but defined in an enclosing scope) with **the value or reference** to which the name was bound **when the closure was created**

相应的中文翻译:

> 环境是一种映射,它将函数的每个**自由变量**(在本地使用但在封闭范围内定义的变量)与**创建闭包时**名称绑定到的**值或引用**相关联。

环境是闭包所处的环境,这里**强调的是外部环境**,更确切的说是**相对于匿名函数而言的外部变**量,像这种**被闭包函数使用但是定义在闭包函数外部**的变量被称为**自由变量**.

> 「雪之梦技术驿站」: 由于闭包函数内部使用了自由变量,所以闭包内部的也就关联了自由变量的**值或引用**,这种绑定关系是创建闭包时确定的,运行时环境也会一直存在并不会发生像普通函数那样无法维持环境.

- 自由变量

这里使用了一个比较陌生的概念: **自由变量**(在本地使用但在封闭范围内定义的变量)

很显然,根据括号里面的注释说明我们知道: 所谓的**自由变量**是相对于闭包函数或者说匿名函数来说的外部变量,由于该变量的定义不受自己控制,所以对闭包函数自己来说就是自由的,并不受闭包函数的约束!

那么按照这种逻辑继续延伸猜测的话,匿名函数内部定义的变量岂不是**约束变量**?对于闭包函数而言的自由变量对于定义函数来说岂不是约束变量?

```go
var a, b = 0, 1
func fibonacciWithoutClosure() int {
  a, b = b, a+b
  return a
}
```

> 「雪之梦技术驿站」: 这里的变量 `a,b` 相对于函数 `fibonacciWithoutClosure` 来说,是不是自由变量?或者说全局变量就是自由变量,对吗?

- 值或引用

```go
func fibonacci() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
```

变量 `a,b` 定义在函数 `fibonacci` 内部,相对于匿名函数 `func() int` 来说是自由变量,在匿名函数中直接使用了变量 `a,b` 并没有重新复制一份,所以这种形式的环境关联的自由变量是**引用**.

再举个**引用关联**的示例,加深一下闭包的环境理解.

```go
func countByClosureButWrong() []func() int {
  var arr []func() int
  for i := 1; i <= 3; i++ {
    arr = append(arr, func() int {
      return i
    })
  }
  return arr
}
```

上述示例的 `countByClosureButWrong` 函数内部定义了变量数组 `arr` ,存储的是匿名函数而匿名函数使用的是**循环变量** `i` .

这里的循环变量的定义部分是在匿名函数的外部就是所谓的**自由变量**,变量 `i` 没有进行拷贝所以也就是**引用关联**.

```go
func TestCountByClosure(t *testing.T) {
  // 4 4 4 
  for _, c := range countByClosureButWrong() {
    t.Log(c())
  }
}
```

运行这种闭包函数,最终的输出结果都是 `4 4 4`,这是因为闭包的环境关联的循环变量 `i` 是**引用方式而不是值传递方式**,所以闭包运行结束后的变量 `i` 已经是 `4`.

除了引用传递方式还有值传递方式,关联自由变量时拷贝一份到匿名函数,使用者调用闭包函数时就能如愿绑定到循环变量.

```go
func countByClosureWithOk() []func() int {
  var arr []func() int
  for i := 1; i <= 3; i++ {
    func(n int) {
      arr = append(arr, func() int {
        return n
      })
    }(i)
  }
  return arr
}
```

> 「雪之梦技术驿站」: 自由变量 `i` 作为参数传递给匿名函数,而 `Go` 中的**参数传递只有值传递**,所以匿名函数使用的变量 `n` 就可以**正确绑定循环变量**了,这也就是自由变量的**值绑定方式**.

```go
func TestCountByClosureWithOk(t *testing.T) {
  // 1 2 3
  for _, c := range countByClosureWithOk() {
    t.Log(c())
  }
}
```

> 「雪之梦技术驿站」: 自由变量通过值传递的方式传递给闭包函数,实现值绑定环境,正确绑定了循环变量 `1 2 3` 而不是 `4 4 4 `

### 访问被捕获自由变量

第四句英文如下:

> Unlike a plain function, a closure allows the function to access those **captured variables** through the closure's **copies of their values or references**, even when the function is invoked **outside their scope**.

相应的中文翻译:

> 与普通函数不同,闭包允许函数通过闭包的**值的副本或引用**访问那些**被捕获的变量**，即使函数在其**作用域之外**被调用

闭包函数和普通函数的不同之处在于,闭包提供一种**持续访问被捕获变量**的能力,简单的理解就是**扩大了变量的作用域**.

```go
func fibonacci() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
```

自由变量 `a,b` 的定义发生在函数 `fibonacci` 体内,一般而言,变量的作用域也仅限于函数内部,外界是无法访问该变量的值或引用的.

但是,闭包提供了持续暴露变量的机制,外界突然能够访问原本应该私有的变量,实现了全局变量的作用域效果!

```go
var a, b = 0, 1
func fibonacciWithoutClosure() int {
  a, b = b, a+b
  return a
}
```

> 「雪之梦技术驿站」: 普通函数想要访问变量 `a,b` 的值或引用,定义在函数内部是无法暴露给调用者访问的,只能提升成**全局变量**才能实现作用域范围的扩大.

由此可见,一旦变量被闭包捕获后,外界使用者是可以访问这些被捕获的变量的值或引用的,相当于**访问了私有变量**!

## 怎么理解闭包

闭包是一种**函数式编程**中实现**名称绑定**的技术,直观表现为**函数嵌套**提升变量的作用范围,使得原本寿命短暂的局部变量获得**长生不死的能力**,只要被捕获到的自由变量一直在使用中,系统就不会回收内存空间!

知乎上关于闭包的众多回答中,其中有一个回答言简意赅,特意分享如下:

> **我叫独孤求败**，我在一个**山洞**里，里面有世界上最好的剑法，还有最好的武器。我学习了里面的剑法，拿走了最好的剑。离开了这里。我来到这个**江湖**，快意恩仇。但是从来没有人知道我这把剑的来历，和我这一身的武功。。。那山洞就是一个闭包，而我，就是那个山洞里唯一一个可以与外界交汇的地方。这山洞的一切对外人而言就像不存在一样，只有我才拥有这里面的宝藏！

这也是闭包定义中最后一句话表达的意思: 山洞是闭包函数,里面的**剑法和武器**就是闭包的**内部环境**,而**独孤求败剑客**则是被捕获的**自由变量**,他出生在山洞之外的世界,学成归来后独自闯荡**江湖**.从此江湖上有了独孤求败的传说和那把剑以及神秘莫测的剑法.

![go-functional-programming-fib-swordsman.jpeg](../images/go-functional-programming-fib-swordsman.jpeg)

## 掌握闭包了么

- 问题: 请将下述普通函数改写成闭包函数?

```go
func count() []int {
  var arr []int
  for i := 1; i <= 3; i++ {
    arr = append(arr, i)
  }
  return arr
}

func TestCount(t *testing.T) {
  // 1 2 3
  for _, c := range count() {
    t.Log(c)
  }
}
```

- 回答: 闭包的错误示例以及正确示例

```go
func countByClosureButWrong() []func() int {
  var arr []func() int
  for i := 1; i <= 3; i++ {
    arr = append(arr, func() int {
      return i
    })
  }
  return arr
}

func TestCountByClosure(t *testing.T) {
  // 4 4 4 
  for _, c := range countByClosureButWrong() {
    t.Log(c())
  }
}

func countByClosureWithOk() []func() int {
  var arr []func() int
  for i := 1; i <= 3; i++ {
    func(n int) {
      arr = append(arr, func() int {
        return n
      })
    }(i)
  }
  return arr
}

func TestCountByClosureWithOk(t *testing.T) {
  // 1 2 3
  for _, c := range countByClosureWithOk() {
    t.Log(c())
  }
}
```

那么,问题来了,原本普通函数就能实现的需求更改成闭包函数实现后,一不小心就弄错了,为什么还需要闭包?

## 闭包归纳总结

现在再次回顾一下**斐波那契数列生成器**函数,相信你已经读懂了吧,有没有看到闭包的影子呢?

```go
// 1 1 2 3 5 8 13 21 34 55
//     a b
//       a b
func fibonacci() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
```

但是,有没有想过这么一个问题: 为什么需要闭包,闭包解决了什么问题?

- 闭包不是某一种语言特有的机制,但常出现在**函数式编程**中,尤其是**函数占据重要地位**的编程语言.
- 闭包的直观表现是**函数内部嵌套了函数**,并且**内部函数访问了外部变量**,从而使得**自由变量**获得**延长寿命**的能力.
- 闭包中使用的自由变量一般有**值传递和引用传递**两种形式,示例中的斐波那契数列生成器利用的是引用而循环变量示例用的是值传递.
- `Go` 不支持函数嵌套但**支持匿名函数**,语法层面的差异性掩盖不了闭包整体的统一性.

> 「雪之梦技术驿站」: 由于篇幅所限,为什么需要闭包以及闭包的优缺点等知识的相关分析打算另开一篇单独讨论,敬请期待...

## 相关资料参考

- [廖雪峰: `js` 的闭包教程](https://www.liaoxuefeng.com/wiki/1022910821149312/1023021250770016)
- [知乎: 什么是闭包](https://www.zhihu.com/question/34210214)
- [三点水博客: 再谈闭包](https://lotabout.me/2016/thoughts-of-closure/)

<script src="https://my.openwrite.cn/js/readmore.js" type="text/javascript"></script>
<script>
    const btw = new BTWPlugin();
    btw.init({
        id: 'container',
        blogId: '15702-1569305559839-744',
        name: '雪之梦技术驿站',
        qrcode: 'https://snowdreams1006.github.io/snowdreams1006-wechat-public.jpeg',
        keyword: 'vip',
    });
</script>