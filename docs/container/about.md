# 内建容器

上篇文章中详细介绍了 `Go` 的基础语言,指出了 `Go` 和其他主流的编程语言的差异性,比较侧重于语法细节,相信只要稍加记忆就能轻松从已有的编程语言切换到 `Go` 语言的编程习惯中,尽管这种切换可能并不是特别顺畅,但多加练习尤其是多多试错,总是可以慢慢感受 `Go` 语言之美!

在学习 `Go` 的内建容器前,同样的,我们先简单回顾一下 `Go` 的基本语言,温度而知新可以为师矣!

## 上节知识回顾

> 如需了解详情,请于微信公众号[雪之梦技术驿站]内查看 [go 学习笔记之值得特别关注的基础语法有哪些](https://mp.weixin.qq.com/s/8Ijk3FGMo9fCSTNGbx8R3Q) 文章,觉得有用的话,顺手转发一下呗!

### 内建类型种类

- `bool` 

> 布尔类型,可选 `true|false`,默认初始化零值 `false` .

- `(u)int` ,`(u)int8` , `(u)int16`, `(u)int32`,`(u)int64`,`uintptr`

> `2^0=1`,`2^1=2` ,`2^2=4` 个**字节**长度的整型,包括有符号整型和无符号整型以及 `uintptr` 类型的指针类型,默认初始化零值 `0` .

- `byte(uint8)` ,`rune(int32)`,`string`

> `byte` 是最基础**字节**类型,是 `uint8` 类型的**别名**,而 `rune` 是 `Go` 中的**字符**类型,是 `int32` 的别名.最常用的字符串类型 `string` 应该不用介绍了吧?

- `float32` ,`float64` ,`complex64` ,`complex128`

> 只有 `float` 类型的浮点型,没有 `double` 类型,同样是以字节长度来区分,`complex64` 是复数类型,实部和虚部由 `float32` 类型复合而成,因此写作 `complex64` 这种形式.

### 内建类型特点

- 类型转换只有显示转换,不存在任何形式的隐式类型转换

> 不同变量类型之间不会自动进行隐式类型转换,`Go` 语言的类型转换只有强制的,只能显示转换.

- 虽然提供指针类型,但指针本身不能进行任何形式的计算.

> 指针类型的变量不能进行计算,但是可以重新改变内存地址的指向.

- 变量声明后有默认初始化零值,变量零值视具体类型而定

> `int` 类型的变量的初始化零值是 `0`,`string` 类型的初始化零值是空字符串,并不是 `nil`

### 基本运算符

- 算术运算符没有 `++i` 和`--i` 

> 只有 `i++` 和 `i--` 这种自增操作,再也不用担心两种方式的差异性了!

- 比较运算符 `==` 可以比较数组是否相等

> 当两个数组的维度和数组长度相等时,两个数组可以进行比较,顺序完全一致时,结果为 `true`,其他情况则是 `false` .

- 位运算符新增按位清零运算符 `&^`

> 其他主流的编程语言虽然没有这种操作符,通过组合命令也可以实现类似功能,但既然提供了按位清零运算符,再也不用自己进行组合使用了!

### 流程控制语句

- `if`  条件表达式不需要小括号并支持变量赋值操作

> 先定义临时变量并根据该变量进行逻辑判断,然后按照不同情况进行分类处理,`Go` 处理这种临时变量的情况,直接对条件表达式进行增强,这种情况以后会很常见!

- `if` 条件表达式内定义的变量作用域仅限于当前语句块

> 条件表达式内定义的变量是为了方便处理不同分支的逻辑,既然是临时变量,出了当前的 `if` 语句块就无法使用,也变得可以理解.

- `switch` 语句可以没有 `break`,除非使用了 `fallthrough`

> `switch` 语句的多个 `case` 结尾处可以没有 `break`,系统会自动进行 `break` 处理.

- `switch` 条件表达式不限制为常数或整数

> 和其他主流的编程语言相比,`Go` 语言的 `switch` 条件表达式更加强大,类型也较为宽松.

- `switch` 条件表达式可以省略,分支逻辑转向 `case` 语言实现.

> 省略 `switch` 条件表达式,多个 `case` 语言进行分支流程控制,功能效果和多重 `if else` 一样.

- 省略 `switch` 条件表达式后,每个 `case` 条件可以有多个条件,用逗号分隔.

> `swicth` 语句本质上是根据不同条件进行相应的流程控制,每个 `case` 的条件表达式支持多个,更是增强了流程控制的能力.

- `for` 循环的条件表达式也不需要小括号,且没有其他形式的循环.

> `Go` 语言只有 `for` 循环,没有 `while` 等其他形式的循环.

- `for` 循环的初始条件,终止条件和自增表达式都可以省略或者同时省略

> 条件表达式进行省略后可以实现 `while` 循环的效果,全部省略则是死循环.

### 函数和参数传递

- 函数声明按照函数名,入参,出参顺序定义,并支持多返回值

> 不论是变量定义还是函数定义,`Go` 总是和其他主流的编程语言反着来,如果按照输入输出的顺序思考就会发现,这种定义方式其实挺有道理的.

- 函数有多个返回值时可以给返回值命名,但对调用者而言没有差别

> 函数返回多个值时可以有变量名,见名知意方便调用者快速熟悉函数声明,但调用者并非一定要按照返回值名称接收调用结果.

- 函数的入参没有必填参数,可选参数等复杂概念,只支持可变参数列表

> 可变参数列表和其他主流的编程语言一样,必须是入参的最后一个.

- 函数参数传递只有值传递,没有引用传递,即全部需要重新拷贝变量

> 参数传递只有值传递,逻辑上更加简单,但是处理复杂情况时可以传递指针实现引用传递的效果.

## 内建容器有哪些

复习了 `Go` 语言的基础语法后,开始继续学习变量类型的承载者也就是容器的相关知识.

承载一类变量最基础的底层容器就是数组了,大多数高级的容器底层都可以依靠数组进行封装,所以先来了解一下 `Go` 的数组有何不同?

### 数组和切片

- 数组的声明和初始化 

数组的明显特点就是一组特定长度的连续存储空间,声明数组时必须指定数组的长度,声明的同时可以进行初始化,当然不指定数组长度时也可以使用 `...` 语法让编译器帮我们确定数组的长度.

```go
func TestArray(t *testing.T) {
    var arr1 [3]int
    arr2 := [5]int{1, 2, 3, 4, 5}
    arr3 := [...]int{2, 4, 6, 8, 10}

    // [0 0 0] [1 2 3 4 5] [2 4 6 8 10]
    t.Log(arr1, arr2, arr3)

    var grid [3][4]int

    // [[0 0 0 0] [0 0 0 0] [0 0 0 0]]
    t.Log(grid)
}
```

> `[3]int` 指定数组长度为 `3`,元素类型为 `int`,当然也可以声明时直接赋值 `[5]int{1, 2, 3, 4, 5}` ,如果懒得指定数组长度,可以用 `[...]int{2, 4, 6, 8, 10}` 表示.

- 数组的遍历和元素访问

最常见的 `for` 循环进行遍历就是根据数组的索引进行访问,`range arr` 方式提供了简化遍历的便捷方法.

```go
func TestArrayTraverse(t *testing.T) {
    arr := [...]int{2, 4, 6, 8, 10}

    for i := 0; i < len(arr); i++ {
        t.Log(arr[i])
    }

    for i := range arr {
        t.Log(arr[i])
    }

    for i, v := range arr {
        t.Log(i, v)
    }

    for _, v := range arr {
        t.Log(v)
    }
}
```

> `range arr` 可以返回索引值和索引项,如果仅仅关心索引项而不在乎索引值的话,可以使用 `_` 占位符表示忽略索引值,如果只关心索引值,那么可以不写索引项.这种处理逻辑也就是函数的多返回值顺序接收,不可以出现未使用的变量.

- 数组是值类型可以进行比较

数组是值类型,这一点和其他主流的编程语言有所不同,因此相同纬度且相同元素个数的数组可以比较,关于这方面的内容前面也已经强调过,这里再次简单回顾一下.

```go
func printArray(arr [5]int) {
    arr[0] = 666
    for i, v := range arr {
        fmt.Println(i, v)
    }
}

func TestPrintArray(t *testing.T) {
    var arr1 [3]int
    arr2 := [5]int{1, 2, 3, 4, 5}
    arr3 := [...]int{2, 4, 6, 8, 10}

    // [0 0 0] [1 2 3 4 5] [2 4 6 8 10]
    t.Log(arr1, arr2, arr3)

    // cannot use arr1 (type [3]int) as type [5]int in argument to printArray
    //printArray(arr1)

    fmt.Println("printArray(arr2)")
    printArray(arr2)

    fmt.Println("printArray(arr3)")
    printArray(arr3)

    // [1 2 3 4 5] [2 4 6 8 10]
    t.Log(arr2, arr3)
}
```

> 因为参数传递是值传递,所以 `printArray` 函数无法更改调用者传递的外部函数值,如果想要在函数 `printArray` 内部更改传递过来的数组内容,可以通过指针来实现,但是有没有更简单的做法?

想要在 `printArrayByPointer` 函数内部修改参数数组,可以通过数组指针的方式,如果有不熟悉的地方,可以翻看上一篇文章回顾查看.

```go
func printArrayByPointer(arr *[5]int) {
    arr[0] = 666
    for i, v := range arr {
        fmt.Println(i, v)
    }
}

func TestPrintArrayByPointer(t *testing.T) {
    var arr1 [3]int
    arr2 := [5]int{1, 2, 3, 4, 5}
    arr3 := [...]int{2, 4, 6, 8, 10}

    // [0 0 0] [1 2 3 4 5] [2 4 6 8 10]
    t.Log(arr1, arr2, arr3)

    fmt.Println("printArrayByPointer(arr2)")
    printArrayByPointer(&arr2)

    fmt.Println("printArrayByPointer(arr3)")
    printArrayByPointer(&arr3)

    // [666 2 3 4 5] [666 4 6 8 10]
    t.Log(arr2, arr3)
}
```

> 修改数组的元素可以通过传递数组指针来实现,除此之外,`Go` 语言中数组还有一个近亲 `slice`,也就是切片,它可以实现类似的效果.

- 切片的声明和初始化

切片和数组非常类似,创建数组时如果没有指定数组的长度,那么最终创建的其实是切片并不是数组.

```go
func TestSliceInit(t *testing.T) {
    var s1 [5]int
    // [0 0 0 0 0]
    t.Log(s1)

    var s2 []int
    // []
    t.Log(s2,len(s2))
}
```

> `[]int` 没有指定长度,此时创建的是切片,默认初始化零值是 `nil`,并不是空数组!

同理,数组可以声明并初始化,切片也可以,并且语法也很类似,稍不注意还以为是数组呢!

```go
func TestSliceInitValue(t *testing.T) {
    var s1 = [5]int{1, 3, 5, 7, 9}
    // [1 3 5 7 9]
    t.Log(s1)

    var s2 = []int{1, 3, 5, 7, 9}
    // [1 3 5 7 9]
    t.Log(s2)
}
```

> 仅仅是没有指定 `[]` 中的长度,最终创建的结果就变成了切片,真的让人眼花缭乱!

数组和切片如此相像,让人不得不怀疑两者之间有什么见不得人的勾当?其实可以从数组中得到切片,下面举例说明:

```go
func TestSliceFromArray(t *testing.T) {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    // arr =  [0 1 2 3 4 5 6 7 8 9]
    t.Log("arr = ", arr)

    // arr[2:6] =  [2 3 4 5]
    t.Log("arr[2:6] = ", arr[2:6])
    // arr[:6] =  [0 1 2 3 4 5]
    t.Log("arr[:6] = ", arr[:6])
    // arr[2:] =  [2 3 4 5 6 7 8 9]
    t.Log("arr[2:] = ", arr[2:])
    // arr[:] =  [0 1 2 3 4 5 6 7 8 9]
    t.Log("arr[:] = ", arr[:])
}
```

> `arr[start:end]` 截取数组的一部分得到的结果就是切片,切片的概念也是很形象啊!

和其他主流的编程语言一样,`[start:end]` 是一个左闭右开区间,切片的含义也非常明确:

忽略起始索引 `start` 时,`arr[:end]` 表示原数组从头开始直到终止索引 `end` 的前一位;
忽略终止索引 `end` 时,`arr[ start:]` 表示原数组从起始索引 `start` 开始直到最后一位;
既忽略起始索引又忽略终止索引的情况,虽然不常见但是含义上将应该就是原数组,但是记得类型是切片不是数组哟!

目前为止,我们知道切片和数组很相似,切片相对于数组只是没有大小,那么切片和数组的操作上是否一样呢?

```go
func updateSlice(s []int) {
    s[0] = 666
}

func TestUpdateSlice(t *testing.T) {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    // arr =  [0 1 2 3 4 5 6 7 8 9]
    t.Log("arr = ", arr)
    
    s1 := arr[2:6]
    // s1 =  [2 3 4 5]
    t.Log("s1 = ", s1)

    s2 := arr[:6]
    // s2 =  [0 1 2 3 4 5]
    t.Log("s2 = ", s2)

    updateSlice(s1)
    // s1 =  [666 3 4 5]
    t.Log("s1 = ", s1)
    // arr =  [0 1 666 3 4 5 6 7 8 9]
    t.Log("arr = ", arr)

    updateSlice(s2)
    // s2 =  [666 1 666 3 4 5]
    t.Log("s2 = ", s2)
    // arr =  [666 1 666 3 4 5 6 7 8 9]
    t.Log("arr = ", arr)
}
```

> 切片竟然可以更改传递参数,这一点可是数组没有做到的事情啊!除非使用数组的指针类型,切片竟然可以轻易做到?除非切片内部是指针,因为参数传递只有值传递,根本没有引用传递方式!

切片和数组在参数传递的表现不同,具体表现为数组进行参数传递时无法修改数组,想要想改数组只有传递数组指针才行,而切片却实现了数组的改变!

由于参数传递只有值传递一种方式,因此推测切片内部肯定存在指针,参数传递时传递的是指针,所以函数内部的修改才能影响到到函数外部的变量.

![go-container-about-slice-struct.png](../images/go-container-about-slice-struct.png)

> `slice` 的内部实现中有三个变量,指针 `ptr`,个数 `len` 和容量 `cap` ,其中 `ptr` 指向真正的数据存储地址.

正是由于切片这种内部实现,需要特性也好表现形式也罢才使得切换和数组有着千丝万缕的联系,其实这种数据结果就是对静态数组的扩展,本质上是一种动态数组而已,只不过 `Go` 语言叫做**切片**!

切片是动态数组,上述问题就很容易解释了,参数传递时传递的是内部指针,因而虽然是值传递拷贝了指针,但是指针指向的真正元素毕竟是一样的,所以切片可以修改外部参数的值.

数组可以在一定程度上进行比较,切片是动态数组,能不能进行比较呢?让接下来的测试方法来验证你的猜想吧!

![go-container-about-slice-compare.png](../images/go-container-about-slice-compare.png)

> 不知道你有没有猜对呢?切片并不能进行比较,只能与 `nil` 进行判断.

- 切片的添加和删除

数组是静态结构,数组的大小不能扩容或缩容,这种数据结构并不能满足元素个数不确定场景,因而才出现动态数组这种切片,接下来重点看下切片怎么添加或删除元素.

```go
func printSlice(s []int) {
    fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func TestSliceAutoLonger(t *testing.T) {
    var s []int
    // []
    t.Log(s)

    for i := 0; i < 10; i++ {
        s = append(s, i)

        printSlice(s)
    }

    // [0 1 2 3 ...,98,99]
    t.Log(s)

    for i := 0; i < 10; i++ {
        s = s[1:]
        
        printSlice(s)
    }

    // [0 1 2 3 ...,98,99]
    t.Log(s)
}
```

> 添加元素 `s = append(s, i)` 需要扩容时,每次以 `2` 倍进行扩容,删除元素 `s[1:]` 时,递减缩容.

`s = append(s, i)` 向切片中添加元素并返回新切片,由于切片是动态数组,当切片内部的数组长度不够时会自动扩容以容纳新数组,扩容前后的内部数组会进行元素拷贝过程,所以 `append` 会返回新的地址,扩容后的地址并不是原来地址,所以需要用变量接收添加后的切片.

当不断进行切片重新截取时 `s[1:]` ,切片存储的元素开始缩减,个数递减,容量也递减.

![go-container-about-slice-add-and-delete.png](../images/go-container-about-slice-add-and-delete.png)

其实除了基于数组创建切片和直接创建切片的方式外,还存在第三种创建切片的方式,也是使用比较多的方式,那就是 `make` 函数.

```go
func TestMakeSlice(t *testing.T) {
    s1 := make([]int,10)

    // s1 = [0 0 0 0 0 0 0 0 0 0], len(s1) = 10, cap(s1) = 10
    t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

    s2 := make([]int, 10, 32)

    // s2 = [0 0 0 0 0 0 0 0 0 0], len(s2) = 10, cap(s2) = 32
    t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}
```

通过 `make` 方式可以设置初始化长度和容量,这是字面量创建切片所不具备的能力,并且这种方式创建的切片还支持批量拷贝功能!

```go
func TestCopySlice(t *testing.T) {
    var s1 = []int{1, 3, 5, 7, 9}
    var s2 = make([]int, 10, 32)

    copy(s2, s1)

    // s2 = [1 3 5 7 9 0 0 0 0 0], len(s2) = 10, cap(s2) = 32
    t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

    var s3 []int

    copy(s3, s1)

    // s3 = [], len(s3) = 0, cap(s3) = 0
    t.Logf("s3 = %v, len(s3) = %d, cap(s3) = %d", s3, len(s3), cap(s3))
}
```

> `func copy(dst, src []Type) int` 是切片之间拷贝的函数,神奇的是,只有目标切片是 `make` 方式创建的切片才能进行拷贝,不明所以,有了解的小伙伴还请指点一二!

切片的底层结构是动态数组,如果切片是基于数组截取而成,那么此时的切片从效果上来看,切片就是原数组的一个视图,对切片的任何操作都会反映到原数组上,这也是很好理解的.

那如果对切片再次切片呢,或者说切片会不会越界,其实都比较简单了,还是稍微演示一下,重点就是动态数组的底层结构.

```go
func TestSliceOutOfBound(t *testing.T) {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

    s1 := arr[2:6]
    // s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
    t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

    s2 := s1[3:5]
    // s2 = [5 6], len(s2) = 2, cap(s2) = 3
    t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}
```

> `[]` 只能访问 `len(arr)` 范围内的元素,`[:]` 只能访问 `cap(arr)` 范围内的元素,一般而言 `cap >= len` 所以某些情况看起来越界,其实并不没有越界,只是二者的标准不同!

我们知道切片 `slice` 的内部数据结构是基于动态数组,存在三个重要的变量,分别是指针 `ptr`,个数 `len` 和容量 `cap` ,理解了这三个变量如何实现动态数组就不会掉进切片的坑了!

个数 `len` 是通过下标访问时的有效范围,超过 `len` 后会报越界错误,而容量 `cap` 是往后能看到的最大范围,动态数组的本质也是控制这两个变量实现有效数组的访问.

![go-container-about-slice-outOfBound-len.png](../images/go-container-about-slice-outOfBound-len.png)

> 因为 `s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6` ,所以 `[]` 访问切片 `s1` 元素的范围是`[0,4)` ,因此最大可访问到`s1[3]`,而 `s1[4]` 已经越界了!

![go-container-about-slice-outOfBound-cap.png](../images/go-container-about-slice-outOfBound-cap.png)

> 因为 `s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6` ,所以 `[:]` 根据切片 `s1` 创建新切片的范围是 `[0,6]` ,因此最大可访问范围是 `s1[0:6]` ,而 `s1[3:7]` 已经越界!

### 集合 `map`

集合是一种键值对组成的数据结构,其他的主流编程语言也有类似概念,相比之下,`Go` 语言的 `map` 能装载的数据类型更加多样化.

- 字面量创建 `map` 换行需保留逗号 `,`

```go
func TestMap(t *testing.T) {
    m1 := map[string]string{
        "author": "snowdreams1006",
        "website": "snowdreams1006",
        "language": "golang",
    }

    // map[name:snowdreams1006 site:https://snowdreams1006.github.io]
    t.Log(m1)
}
```

> 一对键值对的结尾处加上逗号 `,` 可以理解,但是最后一个也要有逗号这就让我无法理解了,`Why` ?

- `make` 创建的 `map` 和字面量创建的 `map` 默认初始化零值不同

```go
func TestMapByMake(t *testing.T) {
    // empty map
    m1 := make(map[string]int)

    // map[] false
    t.Log(m1, m1 == nil)

    // nil
    var m2 map[string]int

    // map[] true
    t.Log(m2, m2 == nil)
}
```

> `make` 函数创建的 `map` 是空 `map`,而通过字面量形式创建的 `map` 是 `nil`,同样的规律也适合于切片 `slice`.

- `range` 遍历 `map` 是无序的

```go
func TestMapTraverse(t *testing.T) {
    m := map[string]string{
        "name": "snowdreams1006",
        "site": "https://snowdreams1006.github.io",
    }

    // map[name:snowdreams1006 site:https://snowdreams1006.github.io]
    t.Log(m)

    for k, v := range m {
        t.Log(k, v)
    }

    t.Log()

    for k := range m {
        t.Log(k)
    }

    t.Log()

    for _, v := range m {
        t.Log(v)
    }
}
```

> 这里再一次遇到 `range` 形式的遍历,忽略键或值时用 `_` 占位,也是和数组,切片的把遍历方式一样,唯一的差别就是 `map` 没有索引,遍历结果也是无序的!

- 获取元素时需判断元素是否存在

```go
func TestMapGetItem(t *testing.T) {
    m := map[string]string{
        "name": "snowdreams1006",
        "site": "https://snowdreams1006.github.io",
    }

    // snowdreams1006
    t.Log(m["name"])

    // zero value is empty string
    t.Log(m["author"])

    // https://snowdreams1006.github.io
    if site, ok := m["site"]; ok {
        t.Log(site)
    } else {
        t.Log("key does not exist ")
    }
}
```

> `Go` 语言的 `map` 获取不存在的键时,返回的是值对应类型的零值,`map[string]string` 返回的默认零值就是空字符串,由于不会报错进行强提醒,这也就要求我们调用时多做一步检查.当键值对存在时,第二个返回值返回 `true`,不存在时返回 `false`.

- 删除键值对时用 `delete` 函数

```go
func TestMapDeleteItem(t *testing.T) {
    m := map[string]string{
        "name": "snowdreams1006",
        "site": "https://snowdreams1006.github.io",
    }

    // map[name:snowdreams1006 site:https://snowdreams1006.github.io]
    t.Log(m)

    delete(m, "name")

    // map[site:https://snowdreams1006.github.io]
    t.Log(m)

    delete(m, "id")

    // map[site:https://snowdreams1006.github.io]
    t.Log(m)
}
```

> `delete(map,key)` 用于删除 `map` 的键值对,如果想要验证是否删除成功,别忘了使用 `value,ok := m[k]` 确定是否存在指定键值对

- 除 `slice`,`map`,`func` 外,其余类型均可键

> 因为 `map` 是基于哈希表实现,所以遍历是无序的,另一方面因为 `slice`,`map`,`func` 不可比较,因为也不能作为键.当然若自定义类型 `struc` 不包含上述类型,也可以作为键,并不要求实现 `hashcode` 和 `equal` 之类的.

- `value` 可以承载函数 `func` 类型

```go
func TestMapWithFunValue(t *testing.T) {
    m := map[int]func(op int) int{}

    m[1] = func(op int) int {
        return op
    }
    m[2] = func(op int) int {
        return op * op
    }
    m[3] = func(op int) int {
        return op * op * op
    }

    // 1 4 27
    t.Log(m[1](1), m[2](2), m[3](3))
}
```

> 再一次说明函数是一等公民,这部分会在以后的函数式编程中进行详细介绍.

### 没有 `set`

`Go` 的默认类型竟然没有 `set` 这种数据结构,这在主流的编程语言中算是特别的存在了!

正如 `Go` 的循环仅支持 `for` 循环一样,没有 `while` 循环一样可以玩出 `while` 循环的效果,靠的就是增强的 `for` 能力.

所以,即使没有 `set` 类型,基于现有的数据结构一样能实现 `set` 效果,当然直接用 `map` 就可以封装成 `set`.

```go
func TestMapForSet(t *testing.T) {
    mySet := map[int]bool{}

    mySet[1] = true

    n := 3

    if mySet[n] {
        t.Log("update", mySet[n])
    } else {
        t.Log("add", mySet[n])
    }

    delete(mySet, 1)
}
```

> 使用 `map[type]bool` 封装实现 `set` 禁止重复性元素的特性,等到讲解到面向对象部分再好好封装,这里仅仅列出核心结构.

## 知识点总结梳理

`Go` 语言是十分简洁的,不论是基础语法还是这一节的内建容器都很好的体现了这一点.

数组作为各个编程语言的基础数据结构,`Go` 语言和其他主流的编程语言相比没有什么不同,都是一片连续的存储空间,不同之处是数组是值类型,所以也是可以进行比较的.

这并不是新鲜知识,毕竟上一节内容已经详细阐述过该内容,这一节的重点是数组的衍生版切片 `slice` .

因为数组本身是特定长度的连续空间,因为是不可变的,其他主流的编程语言中有相应的解决方案,其中就有不少数据结构的底层是基于数组实现的,`Go` 语言的 `slice` 也是如此,因此个人心底里更愿意称其为动态数组!

切片 `slice` 的设计思路非常简单,内部包括三个重要变量,包括数组指针 `ptr`,可访问元素长度 `len` 以及已分配容量 `cap` .

当新元素不断添加进切片时,总会达到已最大分配容量,此时切片就会自动扩容,反之则会缩容,从而实现了动态控制的能力!

- 指定元素个数的是数组,未指定个数的是切片

```go
func TestArrayAndSlice(t *testing.T) {
    // array
    var arr1 [3]int
    // slice
    var arr2 []int

    // [0 0 0] []
    t.Log(arr1,arr2)
}
```

- 基于数组创建的切片是原始数组的视图

```go
func TestArrayAndSliceByUpdate(t *testing.T) {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // arr =  [0 1 2 3 4 5 6 7 8 9]
    t.Log("arr = ", arr)

    s := arr[2:6]

    // before update s = [2 3 4 5], arr = [0 1 2 3 4 5 6 7 8 9]
    t.Logf("before update s = %v, arr = %v", s, arr)

    s[0] = 666

    // after update s = [666 3 4 5], arr = [0 1 666 3 4 5 6 7 8 9]
    t.Logf("after update s = %v, arr = %v", s, arr)
}
```

- 添加或删除切片元素都返回新切片

```go
func TestArrayAndSliceIncreasing(t *testing.T) {
    var s []int

    fmt.Println("add new item to slice")

    for i := 0; i < 10; i++ {
        s = append(s, i)

        fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
    }

    fmt.Println("remove item from slice")

    for i := 0; i < 10; i++ {
        s = s[1:]

        fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
    }
}
```

- `[index]` 访问切片元素仅仅和切片的 `len` 有关,`[start:end]` 创建新切片仅仅和原切片的 `cap` 有关

```go
func TestArrayAndSliceBound(t *testing.T) {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    s1 := arr[5:8]

    // s1[0] = 5, s1[2] = 7
    t.Logf("s1[0] = %d, s1[%d] = %d", s1[0], len(s1)-1, s1[len(s1)-1])
    // s1 = [5 6 7], len(s1) = 3, cap(s1) = 5
    t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

    s2 := s1[3:5]

    // s2[0] = 8, s2[1] = 9
    t.Logf("s2[0] = %d, s2[%d] = %d", s2[0], len(s2)-1, s2[len(s2)-1])
    // s2 = [8 9], len(s2) = 2, cap(s2) = 2
    t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}
```

- 只有 `map` 没有 `set` 

```go
func TestMapAndSet(t *testing.T) {
    m := map[string]string{
        "name": "snowdreams1006",
        "site": "https://snowdreams1006.github.io",
        "lang": "go",
    }

    // https://snowdreams1006.github.io
    if site, ok := m["site"]; ok {
        t.Log(site)
    } else {
        t.Log("site does not exist ")
    }

    s := map[string]bool{
        "name": true,
        "site": true,
        "lang": true,
    }

    // Pay attention to snowdreams1006
    if _, ok := m["isFollower"]; ok {
        t.Log("Have an eye on snowdreams1006")
    } else {
        s["isFollower"] = true
        t.Log("Pay attention to snowdreams1006")
    }
}
```

- `delete` 函数删除集合 `map` 键值对

```go
func TestMapAndSetByDelete(t *testing.T) {
    m := map[string]string{
        "name": "snowdreams1006",
        "site": "https://snowdreams1006.github.io",
        "lang": "go",
    }

    delete(m, "lang")

    // delete lang successfully
    if _,ok := m["lang"];!ok{
        t.Log("delete lang successfully")
    }
}
```

关于 `Go` 语言中内建容器是不是都已经 `Get` 了呢?如果有表述不对的地方,还请指正哈,欢迎一起来公众号[雪之梦技术驿站]学习交流,每天进步一点点!


