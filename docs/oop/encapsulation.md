# 再说封装

>   关注公众号[雪之梦技术驿站]查看上篇文章 [猜猜看go是不是面向对象语言？能不能面向对象编程？](https://mp.weixin.qq.com/s/l-nqzYjcre2g__BrVptR9A)

虽然在上篇文章中,我们通过尝试性学习探索了 `Go` 语言中关于面向对象的相关概念,更确切的说是关于封装的基本概念以及相关实现.

但那还远远不够,不能满足于一条路,而是应该尽可能地多走几条路,只有这样才能为以后可能遇到的问题积攒下来经验,所以这一节我们将继续探索封装.

## 何为探索性学习

![go-oop-encapsulation-exploratory-learning.jpeg](../images/go-oop-encapsulation-exploratory-learning.jpeg)

通过现有知识加上思想规则指导不断猜想假设逐步验证的学习过程是探索性学习,这样既有利于我们思考又能加深我们对新知识的理解,何乐而不为?

学习 `Go` 语言的过程越发觉得吃力,倒不是因为语法晦涩难懂而是因为语法习惯背后蕴藏的思维习惯差异性太大!

`Go` 语言相对于其他主流的编程语言来说是一种新语言,不仅体现在语法层面更重要的是实现思路的差异性.

尤其是对于已有其他编程经验的开发者而言,这种体会更加深刻,原本可能觉得理所应当的事情到了 `Go` 语言这里基本上都变了模样,很大程度上都换了一种思路去实现,这其实是一件好事,不同的思维碰撞才能促进思考进步,一成不变的话,谈何创新发展?

在这里不得不感谢强大的 `IDE` 开发工具,没有它我们就不能及时发现错误,正是这种快速试错的体验才给我们足够的反馈,运用已有的编程经验逐步接近 `Go` 语言编程的真相.

上篇文章中已经确定主线方向,基本上弄清楚了面向对象中的封装概念以及实现,为了不遗漏任何可能重要的知识点,本文将继续开放性探索,力争讲解清楚封装的知识点.

如果这种学习的过程用走迷宫来比喻的话,一条道走到黑这种策略就是算法理论中的深度优先算法.如果边走边看,四处观望周围的风景就是广度优先算法.

![go-oop-encapsulation-exploratory-maze.jpg](../images/go-oop-encapsulation-exploratory-maze.jpg)

所以,聪明的你肯定已经猜到了,上文采用的正是深度优先算法而本文则采用广度优先算法继续探索封装对象之旅!

## 定义结构体

结构体的定义方式只有一种,或者不存在简化形式吗?

个人觉得不会不存在简化形式,当结构体存在多个字段,标准定义方式是合理使用的,但要是字段只有一个,仍然以标准形式定义结构体未免有种杀鸡焉用牛刀的感觉.

```go
type MyDynamicArray struct {
    ptr *[10]int
    len int
    cap int
}
```

所谓的结构体只不过是实现封装的一种手段,当封装的对象只有一个字段时,这个字段也就不存在字段名或者说这个唯一的字段名应该就可以由编译器自动定义,因此字段名可以省略.

字段类型肯定是不可或缺的,这么想的话,对于封装只有一个字段的对象来说,只需要考虑的是这个唯一字段的类型.

基于上述原因,个人觉得是这种猜想是合情合理的,但是按照已有的知识能否实现呢?

简单起见,暂时先以上篇文章中关于动态数组的结构体声明为例作为测试案例.

```go
type MyDynamicArray struct {
    ptr *[10]int
    len int
    cap int
}
```

如果一定要从三个字段中选择一个字段,那只能是保留内部数组,排除其余字段了,同时最终结果上可能实现不了动态数组的功能,语义上会有所欠缺,那就不论语义,只谈技术!

由于只保留内部数组,动态数组就变成下面这样.失去了动态数组的语义,命名上也做了改变,姑且称之为 `MyArray` 吧!

```go
type MyArray struct {
    arr [10]int
}
```

很明显,现在仍然是结构体的标准语法形式,请随我一起思考一下如何简化这种形式?

因为这种简化形式的内部字段只有一个,所以字段名必须省略而字段类型可能不同,因此应该在简化形式中只保留声明内部字段类型的部分.

```go
type MyArray struct {
    [10]int
}
```

由于多个字段时才需要换行分隔,一个字段自然是不需要换行的,因此大括号也是没必要存在的,这也是符合 `Go` 设计中尽可能精简的情况下保证语义清晰的原则.

当然如果你问我是否真的有这个原则的话,我的回答是可能有也可能没有.

因为我也不知道,只是近期学习 `Go` 语言的一种感觉,处处体现了这么一种哲学思想,也不用较真,只是个人看法.

```go
type MyArray struct [10]int
```

现在这种形式应该可以算是只有一种字段的结构体的简化形式,`struct` 语义上指明了 `MyArray` 是结构体,紧随后面的 `[10]int` 语义上表示结构体的类型,整体上就是说 `MyArray` 结构体的类型是 `[10]int` .

现在让我们在编辑器中测试一下,看一看 `Go` 的编译会不会报错,能否验证我们的猜测呢?

![go-oop-encapsulation-struct-sole-error.png](../images/go-oop-encapsulation-struct-sole-error.png)

很遗憾,`IDE` 编辑器告诉我们 `[10]int` 不合法,必须是类型或类型指针!

可 `[10]int` 确实是我们需要的类型啊,既然报错也就是说`Go` 编译器不支持这种简化形式!

个人猜测可能是 `struct` 关键字不支持这种简化形式,那就去掉这个关键字好了!

![go-oop-encapsulation-struct-sole-success.png](../images/go-oop-encapsulation-struct-sole-success.png)

没想到真的可以!

至少现在看来 `Go` 编译器是支持简化形式的,至于这种支持的形式和我们预期实现的语义是否一致,暂时还不好说,继续做实验探索吧!

![go-oop-encapsulation-struct-sole-test.png](../images/go-oop-encapsulation-struct-sole-test.png)

通过声明变量后直接打印,初步证明了我们这种简化形式是可以正常工作的,输出结果也是我们定义的内部数组!

接下来看一看能不能对这个所谓的内部数组进行操作呢?

这种简化形式只有一个字段,只指明了字段的类型,没有字段名,因而访问该字段应该直接通过结构体变量访问,不知道这种猜测是否正确,依旧做实验来证明.

```go
type MyArray [10]int

func TestMyArray(t *testing.T) {
	var myArr MyArray

	// [0 0 0 0 0 0 0 0 0 0]
	t.Log(myArr)

	myArr[0] = 1
	myArr[9] = 9

	// [1 0 0 0 0 0 0 0 0 9]
	t.Log(myArr)
}
```

这一次猜想也得到了验证,`Go` 编译器就是通过结构体变量直接操作内部字段,看来我们离真相更进一步!

先别急着高兴,将唯一的字段换成其他类型,多测试几遍看看是否依然正常?

```go
type MyBool bool

func TestMyBool(t *testing.T) {
	var myBool MyBool

	// false
	t.Log(myBool)

	myBool = true

	// true
	t.Log(myBool)
}
```

一番测试后并没有报错,很有可能这是 `Go` 所支持的结构体简化形式,也和我们的预期一致.

关于结构体属性的语法规则暂时没有其他探索的新角度,接下来开始探索结构体的方法.

探索的过程中要尽可能的设身处地思考 `Go` 语言应该如何设计才能方便使用者,尽可能地把自己想象成 `Go` 语言的设计者!

![go-oop-encapsulation-struct-sole-method.png](../images/go-oop-encapsulation-struct-sole-method.png)

结构体的简化形式下可能并不支持方法,如果真的是这样的话,这样做也有一定道理.

首先就语法层面分析,为什么单字段的结构体不支持方法?

还记得我们想要简化单字段结构体遇到的报错提示吗?

```go
type MyArray struct [10]int
```

如果直接将单字段类型放到 `struct` 关键字后面,`Go` 编译器就会报错,当我们省略 `struct` 关键字时上述报错自然就消失了.

从`Go` 编译器的角度上来讲,`struct` 是系统关键字,告诉编译器只要遇到这个关键字就解析成结构体语法,现在没有遇到 `sruct` 关键字也就意味着不是结构体语法.

这里关键字和结构体是一一对应关系,也就是充分必要条件,由关键字可以推测到结构体,由结构体也可以推测到关键字.

再回来看一看,我们的单字段结构体是怎么定义的呢?

```go
type MyArray [10]int
```

因为没有关键字 `struct` ,所以编译器推断 `MyArray` 不是结构体,既然不是结构体,也不能用结构体的接收者函数去定义方法.

```go
func (myBool *MyBool) IsTrue() bool{
	return myBool
}
```

所以这种方法就会报错,由此可见 ,`Go` 语言如果真的不支持单字段结构体方法也有理可循.

然后我们再从语义的角度上解释一下为什么不支持方法?

回到探索的初衷,当正在定义的结构体有多个字段时,应该按照标准写法为每个字段指定字段的名称和类型.

假如该字段有且只有一个时,再按照标准写法定义当然可以,但也应该提供更加简化的写法.

只有一个字段的结构体,字段名称是没有意义的也是不应该出现的,因为完全可以用结构体变量所代替,此时这个结构体唯一有存在价值的就是字段的类型了!

字段类型包括内建类型和用户自定义结构体类型,不论哪种类型,这种简化形式的结构体的语义上完全可以由该结构体的字段类型所决定,所以简化形式的结构体还需要方法吗?

自然是不需要的!

字段类型可以由字段类型自己定义的,也能确保职责清晰,彼此分离!

综上,个人觉得即便 `Go` 真的不支持单字段结构体的方法,背后的设计还是有章可循的,有理可依的!

上文中定义动态数组时,内部使用的数组是静态数组,现在为了方便继续探索方法,应该提供重载方法使其支持动态数组.

```go
func NewMyDynamicArray() *MyDynamicArray {
	var myDynamicArray MyDynamicArray

	myDynamicArray.len = 0
	myDynamicArray.cap = 10
	var arr [10]int
	myDynamicArray.ptr = &arr

	return &myDynamicArray
}
```

内部数组 `arr` 是静态数组,应该提供可以让外部调用者初始化指定数组的接口,按照已知的面向对象中关于方法的定义来重载方法.

![go-oop-encapsulation-struct-method-overload.png](../images/go-oop-encapsulation-struct-method-overload.png)

初次尝试方法的重载就遇到了问题,报错提示该方法已声明,所以说 `Go` 可能并不支持方法重载,这样就有点麻烦了.

想要实现类似的功能要么通过定义不同的方法名,要么定义一个非常大的函数,接收最全的参数,再根据调用者参数进行对应的逻辑处理.

用惯了方法的重载,突然发现这种特性在 `Go` 语言中无法实现,顿时有点沮丧,和其他主流的面向对象语言差异性也太大了吧!

不支持构造函数,不支持方法重载,原来以为理所应当的特性并不理所应当.

还是先冷静下来想一想,`Go` 为什么不支持方法重载呢?难不成和构造函数那样,怕是滥用干脆禁用的逻辑?

因为我不是设计者,无法体会也不想猜测原因,但可以肯定的是,`Go` 语言是一门全新的语言,有着独特的设计思路,不与众人同!

吐槽时间结束,既然上了贼船就得一条道走到黑,不支持方法重载就换个函数名或者按参数名区分.

![go-oop-encapsulation-struct-method-error.png](../images/go-oop-encapsulation-struct-method-error.png)

天啊撸,刚刚解决方法重载问题又冒出数组初始化不能是变量只能是常量表达式?

简直不可思议!

既然数组初始化长度只是常量表达式,也就无法接收外部传递的容量 `cap`,没有了容量只能接收长度 `len` ,而初始化内部数组长度又没办法确定了,两个变量都无法对外暴露!

一切又回到原点,想要实现动态数组的功能只能靠具体的方法中去动态扩容和缩容,不能初始化指定长度了.

这样的话,关于方法也是一条死路,停止探索.

## 声明结构体

结构体定义基本已经探索完毕,除了发现一种单字段结构体的简化形式外,暂时没有新的发现.

再次回到使用者的角度上,声明结构体有没有其他方式呢?

```go
var myDynamicArray MyDynamicArray
	
t.Log(myDynamicArray)
```

这是变量的声明方式,除了这种形式,还记得在学习 `Go` 的变量时曾经介绍过声明并初始化变量方式,是否也适用于结构体变量呢?

```go
var myDynamicArray = MyDynamicArray{
		
}

t.Log(myDynamicArray)
```

编译器没有报错,证明这种字面量形式也是适用的,不过空数据结构没有太大的意义,怎么能初始化对应的结构呢?

和多字段结构体最为相似的数据结构莫过于映射 `map` 了!

回忆一下 `map` 如何进行字面量初始化的吧!

```go
var m = map[string]string{
	"id":   "1006",
	"name": "雪之梦技术驿站",
}

t.Log(m)
```

模仿这种结构看看能不能对结构体也这么初始化,果然就没有那么顺利!

![go-oop-encapsulation-struct-init-field-error.png](../images/go-oop-encapsulation-struct-init-field-error.png)

我还没定义,你就不行了?

`IDE` 编辑器提示字段名称无效,结构体明明就有 `len` 字段啊,除非是没有正确识别!

`"len"` 与 `len` 是不一样的吧?

那就去掉双引号 `""` 直接使用字段名进行定义看看.

```go
var myDynamicArray = MyDynamicArray{
	len: 10,
}

t.Log(myDynamicArray)
```

此时报错消失了,成功解锁一种新的隐藏技能.

```go
var myDynamicArray = MyDynamicArray{
	ptr: &[10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	len: 10,
	cap: 10,
}

t.Log(myDynamicArray)
```

除了这种指定字段名称注入方式,能不能不指定字段名称而是按照字段顺序依次初始化?

![go-oop-encapsulation-struct-init-field-in-order.png](../images/go-oop-encapsulation-struct-init-field-in-order.png)

借助编辑器可以看到确实是按照顺序注入的,这样的话,其实有点意思了,明明不支持构造函数,采用字面量实例化时却看起来像构造函数的无参,有参数和全参形式?

可以预想到的是,这种全参注入的方式一定是严格按照定义顺序相匹配的,当参数不全时可能按位插入也可能不支持,真相如何,一试便知!

![go-oop-encapsulation-struct-init-field-lack-order.png](../images/go-oop-encapsulation-struct-init-field-lack-order.png)

事实上并不支持这种参数不全的形式,因此个人觉得要么无参要么全参要么指定初始化字段这三种语义上还是比较清楚的.

除了字面量的方式,`Go` 是否支持创建 `slice` 或 `map` 时所使用的 `make` 函数呢?

![go-oop-encapsulation-struct-init-field-make-error.png](../images/go-oop-encapsulation-struct-init-field-make-error.png)

看样子,`make` 函数并不支持创建结构体,至于为什么不支持,原因就不清楚了,也是个人的一个疑惑点.

既然 `make` 可以创建 `slice` ,`map` 这种内建类型,语义上就是用来创建类型的变量,而结构体也是一种类型,唯一的差别可能就是结构体大多是自定义类型而不是内建类型.

如果我来设计的话,可能会一统天下,因为语义上一致的功能只使用相同的关键字.

回到面向对象的传统编程规范上,一般实例化对象用的是关键字 `new`,而 `new` 并不是 `Go` 中的关键字.

`Go` 语言中的函数是一等公民,正如刚才说的 `make` 也不是关键字,同样是函数.

![go-oop-encapsulation-struct-init-field-new-error.png](../images/go-oop-encapsulation-struct-init-field-new-error.png)

即便对于同一个目标,`Go` 也是有着自己的独到见解!

`new` 不是以关键字形式出现而是以函数的身份登场,初步推测应该也具备实例化对象的能力吧?

![go-oop-encapsulation-struct-init-field-new-assignment-error.png](../images/go-oop-encapsulation-struct-init-field-new-assignment-error.png)

难道 `new` 函数不能实例化对象?为什么报错说赋值错误,难不成姿势不对?

吓得我赶紧看一下 `new` 的文档注释.

```go

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
```

根据注释说明,果然是使用姿势不对,并不像其他的面向对象语言那样可以重复赋值,`Go` 不支持这种形式,还是老老实实初始化声明吧!

```go
var myDynamicArray2 = new(MyDynamicArray)
	
t.Log(myDynamicArray2)	
```

既然存在着两种方式来实例化对象,那么总要看一下有什么区别.

```go
func TestNewMyDynamicArray(t *testing.T) {
	var myDynamicArray = MyDynamicArray{
		ptr: &[10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		len: 10,
		cap: 10,
	}
	myDynamicArray = MyDynamicArray{
		&[10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		10,
		10,
	}
	t.Log(myDynamicArray)
	t.Logf("%[1]T %[1]v", myDynamicArray)

	var myDynamicArray2 = new(MyDynamicArray)
	myDynamicArray2.ptr = &[10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myDynamicArray2.len = 10
	myDynamicArray2.cap = 10

	t.Log(myDynamicArray2)

	t.Logf("%[1]T %[1]v", myDynamicArray2)
}
```

![go-oop-encapsulation-struct-init-field-diff.png](../images/go-oop-encapsulation-struct-init-field-diff.png)

这里简单解释下 `t.Logf("%[1]T %[1]v", myDynamicArray)` 语句是什么意思?

`%[1]T` 其实是 `%T` 的变体,`%[1]v` 也是 `%v` 的变体,仔细观察的话就会发现占位符刚好都是同一个变量,这里也就是第一个参数,所以就用 `[1]` 替代了,再次体现了 `Go` 语言设计的简洁性.

下面再举一个简单的例子加深印象,看仔细了哦!

```go
test := "snowdreams1006"

// string snowdreams1006
t.Logf("%T %v", test, test)
t.Logf("%[1]T %[1]v", test)
```

> `%T` 是打印变量的类型,应该是类型 `type` 的缩写,`v` 应该是值 `value` 的缩写.

解释清楚了测试代码的含义,再回头看看测试结果,发现采用字面量方式得到的变量类型和 `new` 函数得到的变量类型明显不同!

具体表现为 `_struct.MyDynamicArray {0xc0000560f0 10 10}` 是结构体类型,而 `*_struct.MyDynamicArray &{0xc000056190 10 10}` 是结构体类型的指针类型.

这种差异也是可以预期的差异,也是符合语义的差异.

字面量实例化的对象是值对象,而 `new` 实例化对象开辟了内存,返回的是实例对象到引用,正如其他编程语言的 `new` 关键字一样,不是吗?

既然说到了值对象和引用对象,再说一遍老生常谈的问题,函数或者说方法传递时应该传递哪一种类型?

## 值传递还是引用传递

接下来的示例和动态数组并没有什么关系,简单起见,新开一个结构体叫做 `Employee`,顺便回顾一下目前学到的封装知识.

```
type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{
		"0",
		"Bob",
		20,
	}
	t.Logf("%[1]T %[1]v", e)

	e1 := Employee{
		Name: "Mike",
		Age:  30,
	}
	t.Logf("%[1]T %[1]v", e1)

	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 18
	t.Logf("%[1]T %[1]v", e2)
}
```

首先测试引用传递,这也是结构体常用的传递方式,行为表现上和其他的主流编程语言表现一致,方法内的修改会影响调用者的参数.

```go
func (e *Employee) toStringPointer() string {
	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToStringPointer(t *testing.T) {
	e := &Employee{"0", "Bob", 20}

	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	t.Log(e.toStringPointer())
}
```

![go-oop-encapsulation-struct-method-diff-pointer.png](../images/go-oop-encapsulation-struct-method-diff-pointer.png)

> `unsafe.Pointer(&e.Name)` 是查看变量的内存地址,可以看出来调用前后的地址是同一个.

```go
func (e Employee) toStringValue() string {
	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToStringValue(t *testing.T) {
	e := Employee{"0", "Bob", 20}

	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	t.Log(e.toStringValue())
}
```

![go-oop-encapsulation-struct-method-diff-value.png](../images/go-oop-encapsulation-struct-method-diff-value.png)

> 调用者发送的内存地址和接收者接收的内存地址不一样,符合期望,值传递都是拷贝变量进行传递的嘛!

值类型还是引用类型的区分无需赘述,接下来请关注一个神奇的事情,方法的接收者是值类型,方法的调用者是不是一定要传递值类型呢?

```go
func (e Employee) toString() string {
	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
```

方法的调用者分别传递值类型和引用类型,两者均能正常工作,是不是很神奇,好像和方法的定义没什么关系一样!

```go
func TestToString(t *testing.T) {
	e := Employee{"0", "Bob", 20}

	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	t.Log(e.toString())
	t.Log((&e).toString())
}
```

![go-oop-encapsulation-struct-method-value-diff.png](../images/go-oop-encapsulation-struct-method-value-diff.png)

> 虽然方法的接收者要求的是值类型,调用者传递的是值类型还是引用类型均可!

![go-oop-encapsulation-struct-method-pointer-diff.png](../images/go-oop-encapsulation-struct-method-value-diff.png)

仅仅更改了方法接收者的类型,调用者不用做任何更改,依然可以正常运行!

这样就很神奇了,方法的接受者不论是值类型还是指针类型,调用者既可以是值类型也可以是指针类型,为什么?

同样的,基于语义进行分析,方法的设计者和调用者之间可以说是松耦合的,设计者的更改对于调用者来说没有太大影响,这也就意味着以后设计者觉得用值类型接收参数不好,完全可以直接更改为指针类型而不用通知调用者调整逻辑!

这其实要归功于 `Go` 语言到设计者很好的处理了值类型和指针类型的调用方式,不论是值类型还是引用类型,一律使用点操作符 `.` 调用方法,并不像有的语言指针类型是 `->` 或 `*` 前缀才能调用指针类型的方法.

有所为有所不为,可能正是看到了这两种调用方式带来的差异性,`Go` 全部统一成点操作符了!

虽然形式上两种调用方式是一样的,但是设计方法或者函数时到底应该是值类型还是指针类型呢?

这里有三点建议可供参考:

- 如果接收者需要更改调用者的值,只能使用指针类型
- 如果参数本身非常大,拷贝参数比较占用内存,只能用指针类型
- 如果参数本身具有状态,拷贝参数可能会影响对象的状态,只能用指针类型
- 如果是内建类型或者比较小的结构体,完全可以忽略拷贝问题,推荐用值类型.

当然,实际情况可能还和业务相关,具体用什么类型还要自行判断,万一选用不当也不用担心,更改一下参数类型就好了也不会影响调用者的代码逻辑.

## 封装后如何访问

封装问题基本上讲解清楚了,一般来说,封装之后的结构体不仅是我们自己使用还有可能提供给外界使用,与此同时要保证外界不能随意修改我们的封装逻辑,这一部分就涉及到访问的控制权限了.

`Go` 语言的访问级别有两种,一种是公开的另一种就是私有的,由于没有继承特性,也不涉及子类和父类之间访问权限继承问题,顿时觉得没有继承也不是什么坏事嘛,少了很多易错的概念!

虽然现在理解起来很简单,具体实际使用上是否便利还不好判断.

关于可见性的命名规范如下:

- 名称一般使用大驼峰命名法即 `CamelCase`
- 首字母大写表示公开的 `public` ,小写表示私有的 `private` .
- 上述规则不仅适用于方法,包括结构体,变量和常量等几乎是 `Go` 语言的全部.

那么问题了,这里的 `public` 和 `private` 是针对谁来说?

`Go` 语言中的基本结构是包 `package`,这里的包和目录有区别,并不像 `Java` 语言那样包和目录严格相关联的,这一点对于 `Java` 小伙伴来说需要特别注意.

包是相关代码的集合,这些代码可能存放于不同的目录文件中,就是通过包 `package` 的声明告诉 `Go`编译器说:我们是一个家族整体.

如果不同的文件目录可以声明在同一个包中,这样相当于允许家族外迁,只保留姓氏就好.

还是用代码说话吧,散落在各地的小伙伴能不能有共同的姓氏!

```go
package main

import (
	"fmt"
	"github.com/snowdreams1006/learn-go/oop/pack"
)

func main() {
	var l = new(pack.Lang)
	l.SetName("Go")
	l.SetWebsite("https://golang.google.cn/")

	fmt.Println(l.ToString())
}

```

![go-oop-encapsulation-package-access-same-directory.png](../images/go-oop-encapsulation-package-access-same-directory.png)

`pack.go` 源码文件和 `pack_test` 测试文件都位于相同的目录 `pack` 下且包的声明也相同都是 `pack`.

这种情况相当于一家氏族位于一个村落中一起生活,和其他语言到表现一致.

现在试一下这个氏族的一部分人能不能搬到其他村落居住呢?

![go-oop-encapsulation-package-access-other-directory-error.png](../images/go-oop-encapsulation-package-access-other-directory-error.png)

难不成跨域地域有点大,不支持定义方法吗?那移动一下使其离 `pack` 目录近一点试试看!

![go-oop-encapsulation-package-access-same-directory-error.png](../images/go-oop-encapsulation-package-access-other-directory-error.png)

还是不行,不能新建子目录,那么和原来在一个目录下呢?

![go-oop-encapsulation-package-access-same-directory-success.png](../images/go-oop-encapsulation-package-access-other-directory-error.png)

只有这样是可以被标识位结构体的方法的,如果不是方法,完全可以任意存放,这一点就不再演示了,小伙伴可自行测试一下哟!

```go
package main

import (
	"fmt"
	"github.com/snowdreams1006/learn-go/oop/pack"
)

func main() {
	var l = new(pack.Lang)
	l.SetName("Go")
	l.SetWebsite("https://golang.google.cn/")

	fmt.Println(l.ToString())

	l.PrintLangName()
}
```

`"github.com/snowdreams1006/learn-go/oop/pack"` 是当前文件中导入依赖包路径,因此调用者能否正常访问到我们封装的结构体.

在当前结构体中的属性被我们设置成了小写字母开头,所以不在同一包是无法访问该属性的.

![go-oop-encapsulation-package-access-private.png](../images/go-oop-encapsulation-package-access-private.png)

## 封装后如何扩展

设计者封装好对象供其他人使用,难免会有疏忽不足之处,此时使用者就需要扩展已存在的结构体了.

如果是面向对象的设计思路,最简单的实现方式可能就是继承了,重写扩展什么的都不在话下,可是 `Go` 并不这么认为,不支持继承!

所以剩下的方法就是组合了,这也是学习面向对象时的前人总结的一种经验: 多用组合少用继承!

现在想一想,`Go` 语言不但贯彻了这一思想,更是严格执行了,因为 `Go` 直接取消了继承特性.

```go
type MyLang struct {
    l *Lang
}

func (ml *MyLang) Print() {
    if ml == nil || ml.l == nil {
        return
    }

    fmt.Println(ml.l.ToString())
}

func TestMyLangPrint(t *testing.T) {
    var l = new(Lang)
    l.SetName("Go")
    l.SetWebsite("https://golang.google.cn/")

    var ml = MyLang{l}

    ml.Print()
}
```

![go-oop-encapsulation-combination-custom.png](../images/go-oop-encapsulation-combination-custom.png)

> 通过自定义结构体内部属性是 `Lang` 类型,进而扩展原来 `Lang` 不具备的方法或者重写原来的方法.



如果我们的自定义结构体刚好只有这么一个属性,完全可以使用简化形式,说到这里其实有必要特别说明一下,专业叫法称之为别名.

```go
type Lan Lang

func (l *Lan) PrintWebsite(){
    fmt.Println(l.website)
}

func TestLanPrintWebsite(t *testing.T) {
    var la = new(Lan)
    la.name = "GoLang"
    la.website = "https://golang.google.cn/"

    la.PrintWebsite()
}
```

作为设计者和使用者都已经考虑到了,封装的基本知识也要告一段落了,由于 `Go` 不支持继承,也没必要演示相关代码,唯一剩下的只有接口了.

虽然 `Go` 同样是不支持多态,但是 `Go` 提供的接口确实与众不同,别有一番滋味在心头,下一节将开始探索接口.

## 关于封装的复盘

- 定义结构体字段

```go
type Lang struct {
    name    string
    website string
}
```

>  结构体有多个字段时彼此直接换行,不用逗号也不用分号之类的,不要多此一举.

- 定义结构体方法

```go
func (l *Lang) GetName() string {
    return l.name
}
```

> 原本是普通的函数,函数名前面加入指向当前结构体的参数时,函数不再是函数而是方法,同时当前结构体参数叫做接收者,类似于其他面向对象语言中的 `this` 或 `self` 关键字实现的效果.

- 字面量声明结构体

```go
func TestInitLang(t *testing.T) {
    l := Lang{
        name:    "Go",
        website: "https://golang.google.cn/",
    }

    t.Log(l.ToString())
}
```

> 字面量声明结构体除了这种类似于有参构造函数使用方式,还有无参和全参构造函数使用方式,这里说的构造函数只是看起来像并不真的是构造函数.

- `new` 声明结构体

```go
func TestPack(t *testing.T) {
    var l = new(Lang)
    l.SetName("Go")
    l.SetWebsite("https://golang.google.cn/")

    t.Log(l.ToString())
}
```

> `new` 函数和其他主流的编程语言 `new` 关键字类似,用于声明结构体,不同于字面量声明方式,`new` 函数的输出对象是指针类型.

- 首字母大小写控制访问权限

> 不论是变量名还是方法名,名称首字母大写表示公开的,小写表示私有的.

- 代码的基本组织单元是包

> 访问控制权限也是针对代码包而言,一个目录下只有一个代码包,包名和目录名没有必然联系.

- 复合扩展已有类型

```go
type MyLang struct {
    l *Lang
}

func (ml *MyLang) Print() {
    if ml == nil || ml.l == nil {
        return
    }

    fmt.Println(ml.l.ToString())
}

func TestMyLangPrint(t *testing.T) {
    var l = new(Lang)
    l.SetName("Go")
    l.SetWebsite("https://golang.google.cn/")

    var ml = MyLang{l}

    ml.Print()
}
```

> 自定义结构体内嵌其他结构体,通过复合而不是继承的方式实现对已有类型的增强控制,也是一种推荐的编程规范.

- 别名扩展已有类型

```go
type Lan Lang

func (l *Lan) PrintWebsite() {
    fmt.Println(l.website)
}
```

> 别名可以看成单字段结构体的简化形式,可以用来扩展已存在的结构体类型,也支持方法等特性.

最后,非常感谢你的阅读,鄙人知识浅薄,如有描述不当的地方,还请各位看官指出,你的每一次留言我都会认真回复,你的转发就是对我最大的鼓励!

![go-oop-encapsulation-hi.png](../images/go-oop-encapsulation-hi.png)

>   如果需要查看相关源码,可以直接访问 [https://github.com/snowdreams1006/learn-go](https://github.com/snowdreams1006/learn-go),同时也推荐关注公众号与我交流.

