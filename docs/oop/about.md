# 面向对象

![go-oop-about-oop-elephant.jpg](../images/go-oop-about-oop-elephant.jpg)

面向对象编程风格深受广大开发者喜欢,尤其是以 `C++`, `Java` 为典型代表的编程语言大行其道,十分流行!

有意思的是这两中语言几乎毫无意外都来源于 `C` 语言,却不同于 `C` 的面向过程编程,这种面向对象的编程风格给开发者带来了极大的便利性,解放了劳动,松耦合,高内聚也成为设计的标准,从而让我们能够更加愉快地复制粘贴,做代码的搬运工,很多第三方工具开箱即用,语义明确,职责清晰,这都是面向对象编程的好处!

`Go` 语言也是来源于 `C` 语言,不知道你是否也会好奇 `Go` 语言是否支持面向对象这种编程风格呢?

准确的说,`Go` 既支持面向对象编程又不是面向对象语言! 

是也不是,难道像是薛定谔的猫一样具有不确定性?

其实这个答案是官方的回答,并不是我个人凭空杜撰而来的,如需了解详情可参考 [Is Go an object-oriented language?](https://golang.google.cn/doc/faq#Is_Go_an_object-oriented_language)

![go-oop-about-schrodinger-cat.png](../images/go-oop-about-schrodinger-cat.png)

为什么这么说呢?

`Go` 支持封装,却不支持继承和多态,所以严格按照面向对象规范来说, `Go` 语言不是面向对象的编程语言.

但是,`Go` 提供的接口是一种非常简单上手且更加通用的方式,虽然和其他主流的编程语言表现形式上略有不同,甚至不能实现多态,但 `Go` 的接口不仅仅适用于结构体,也可以适用于任何数据类型,这无疑是很灵活的!

争议性比较大的当属继承,由于没有任何关键字支持继承特性,因此是找不到继承的痕迹.虽然的确存在着某些方式可以将类型嵌入到其他类型中以实现子类化,但那却不是真正的继承.

所以说,`Go` 既支持面向对象的编程风格又不完全是面向对象的编程语言.

如果换个角度看问题的话,正是由于没有继承特性使得`Go` 相对于面向对象编程语言更加轻量化,不妨想一想继承的特性,子类和父类的关系,单继承还是多继承,访问控制权限等问题吧!

![go-oop-about-go-lovely.png](../images/go-oop-about-go-lovely.png)

如果按照面向对象的编程规范,实现封装特性的那部分应该是类和对象,但这种概念与实现语言的关键字`class` 是密不可分的,然而 `Go` 并没有 `class` 关键字而是 `C` 语言家族的 `struct` 关键字,所以叫做类或对象也不是十分贴切,所以下面的讲解过程还是采用结构体吧!

## 如何定义结构体

`stuct` 关键字声明结构体,属性之间回车换行.

比如下面示例中定义了动态数组结构体,接下来的示例中都会以动态数组结构体作为演示对象.

```go
type MyDynamicArray struct {
	ptr *[]int
	len int
	cap int
}
```

`Go` 语言中定义对象的多属性时使用直接换行方式而不是分号来分隔?为什么和其他主流的编程语言不呢?

对于习惯分号结尾的开发者可能一时并不习惯 `Go` 的这种语法,于是决定探索一下 `Go` 的编程规范!

![go-oop-about-redundant-semicolon.png](../images/go-oop-about-redundant-semicolon.png)

> 如果手动添加分号的话,编辑器则会提示分号重复,所以猜想是可能是`Go`编译器已经自动添加了分号,并将分号作为语句声明的分隔符,手动添加分号后,`Go` 不管不顾还是添加了分号,于是就有了上述的报错.

这样做有什么好处呢?

自己添加分号和编译器无条件添加分号结果不都是一样的吗,更何况其他主流的编程语言都是手动添加分号的啊!

存在多个属性时直接换行而不是添加分号作为分隔符,对于从未接触过编程语言的小白来说,可能会省事儿,但是对于已有编程经验的开发者来说,却需要特别记住不能添加分号,这一点确实有些闹腾!

如果多个属性全部写在一行时,没有换行符我看你还怎么区分,此时用逗号分隔还是用分号分隔呢?

![go-oop-about-semicolon-or-new-line.png](../images/go-oop-about-semicolon-or-new-line.png)

首先空格肯定是不能分隔多个属性的,因此尝试分号或者逗号是否可以.

根据提示说需要分号或者新的换行符,而换行符是标准形式,所以接下来试一下分号能不能分隔?

![go-oop-about-semicolon-in-one-line.png](../images/go-oop-about-semicolon-in-one-line.png)

编辑器此时没有报错或警告信息,因此在一行上多个属性之间应该用分号分隔,也就是说 `Go` 编译器识别多个属性仍然是同其他主流的编程语言一样,使用分号分隔,而开发者却不能用!

![go-oop-about-semicolon-fire-by-official.jpeg](../images/go-oop-about-semicolon-fire-by-official.jpeg)

类似于上述的规则记忆很简单,验证也比较容易,难点在于理解为什么?

 `Go` 为什么会这么设计?或者说如何理解这种设计思路所代表的语义?

`Go` 作为一门新的编程语言,不仅体现在具体的语法差异上,更重要的是编程思想的特殊性.

正如面向对象中的接口概念一样,设计者只需要定义抽象的行为并不用关心行为的具体实现.

如果我们也采用这种思路来理解不同的编程语言,那么就能透过现象看本质了,否则真的很容易陷入语法细节上,进而可能忽略了背后的核心思想.

其实关于结构体的多属性分隔符问题上,实际上不论采用什么作为分隔符都行,哪怕就是一个逗号,句号都行,只要能让编译器识别到这是不同的属性就行.

由于大多数主流的编程语言一般采用分号作为分隔符,开发者需要手动编写分隔号以供编译器识别,而 `Go` 语言却不这么认为,算了吧,直接换行,我一样可以识别出来(尽管底层 `Go` 编译器进行编译时仍然是采用分号表示换行的)!

![go-oop-about-semicolon-ninja.png](../images/go-oop-about-semicolon-ninja.png)

添加或者不添加分号,对于开发者而言,仅仅是一种分隔多个属性的标志而已,如果能不添加就能实现,那为什么还要添加呢?

![go-oop-about-semicolon-ok.png](../images/go-oop-about-semicolon-ok.png)

是什么,为什么和怎么样是三个基本问题,如果是简单学习了解的话,学会是什么和怎么样就已经足够了,但是这样一来学着学着难免会陷入各自为政的场面,也就是说各个编程语言之间没有任何关系,每一种语言都是独立存在的?!

世界语言千千万,编程语言也不少,学了新语言却没有利用旧语言,那学习新语言时和纯小白有何差异?

学到是学会了,可惜却对旧语言没什么帮助并没有加深旧语言的理解,只是单纯的学习一种全新的语言罢了.

语言是演变创造出来的,不是空中楼阁,是建立在已有体系下逐渐发展演变而来,任何新语言都能或多或少找到旧语言的影子.

所以何不尝试一下,弄清楚新语言设计的初衷和以及设计时所面临的问题,然后再看该语言是如何解决问题的,解决的过程称之为实现细节,我想这种方式应该是一种比较好的学习方式吧!

虽然无法身处语言设计时环境,也不一定明白语言设计时所面临的挑战,但先问尝试着问一下为什么,不这么设计行不行诸如此类的问题,应该是一种不错的开端.

所以接下来的文章都会采用语义性分析的角度,尝试理解 `Go`语言背后的设计初衷,同时以大量的辅助性的测试验证猜想,不再是简单的知识罗列整理过程,当然必要的知识归纳还是很重要的,这一点自然也不会放弃.

![go-oop-about-go-cheer.png](../images/go-oop-about-go-cheer.png)

现在动态数组已经定义完毕,也就是作为设计者的工作暂时告一段落,那作为使用者,如何使用我们的动态数组呢?

按照面向对象的说法,由类创造出对象的过程叫做实例化,然而我们已经知道 `Go` 并不是完全的面向对象语言,因此为了尽可能避免用面向对象的专业术语去称呼 `Go` 的实现细节,我们暂时可以将其理解为结构体类型和结构体变量的关系,以后随着学习的深入,可能会对这部分有更加深刻的认识.

```go
func TestMyDynamicArray(t *testing.T){
	var arr MyDynamicArray

	// {<nil> 0 0}
	t.Log(arr)
}
```

上述写法并没有特殊强调过,完全是用前几篇文章中已经介绍过的语法规则实现的,`var arr MyDynamicArray` 表示声明类型为 `MyDynamicArray` 的变量 `arr` ,此时直接打印该变量的值,得到的是 `{<nil> 0 0}`.

后两个值都是 `0`,自然很好理解,因为在讲解 `Go` 语言中的变量时我们就已经介绍过,`Go` 的变量类型默认初始化都有相应的零值,`int` 类型的 `len cap` 属性自然就是 `0`,而 `ptr *[]int` 是指向数组的指针,所以是 `nil`.

等等,有点不对劲,这里有个设计错误,明明叫做动态数组结果内部却是切片,这算怎么回事?

先修正这个错误再说,由此可见,一时粗心影响多么恶劣以至于语义都变了,容我先改正过来!

![go-oop-about-myDynamicArray-array-size-with-cap.png](../images/go-oop-about-myDynamicArray-array-size-with-cap.png)

我们知道要使用数组必须指定数组的初始化长度,第一感觉是使用 `cap` 表示的容量来初始化 `*[cap]int` 数组,然而并不可以,编辑器提示说必须使用整型数字.

虽然 `cap` 是 `int` 类型的变量,但内部数组 `[cap]int` 并不能识别这种方式,可能是因为这两个变量时一块声明的,`cap` 和 `[cap]int` 都是变量,无法分配.

那如果指定初始化长度应该指定多少呢,如果是 `0` 的话,语义上正确但和实际使用情况不符合,因为这样一来内部数组根据就没办法插入了!

![go-oop-about-myDynamicArray-array-size-with-zero.png](../images/go-oop-about-myDynamicArray-array-size-with-zero.png)

所以数组的初始化长度不能为零,这样解决了无法操作数组的问题,但语义上又不正确了,因此这种情况下需要维护两个变量 `len` 和 `cap` 的值来确保语义和逻辑正确,其中 `len` 表示真正的数组个数,`cap` 表示内部数组实际分配的长度,由于这两个变量至关重要,不应该被调用者随意修改,最多只能查看变量的值,所以必须提供一种机制保护变量的值.

接下来,我们尝试用函数封装的思路来完成这种需求,代码实现如下:

```go
type MyDynamicArray struct {
	ptr *[10]int
	len int
	cap int
}

func TestMyDynamicArray(t *testing.T){
	var myDynamicArray MyDynamicArray

	t.Log(myDynamicArray)

	myDynamicArray.len = 0
	myDynamicArray.cap = 10
	var arr [10]int
	myDynamicArray.ptr = &arr

	t.Log(myDynamicArray)
	t.Log(*myDynamicArray.ptr)
}
```

`var myDynamicArray MyDynamicArray` 声明结构体变量后并设置了结构体的基本属性,然后操作了内部数组,实现了数组的访问修改.

![go-oop-about-myDynamicArray-array-size-with-init.png](../images/go-oop-about-myDynamicArray-array-size-with-init.png)

然而,我们犯了一个典型的错误,调用者不应该关注实现细节,这不是一个封装该干的事!

具体实现细节应该由设计者完成,将有关数据封装成一个整体对外提供相应的接口,这样调用者才能安全方便地调用.

第一步,先将与内部数组相关的两个变量进行封装,对外仅提供访问接口不提供设置接口,防止调用者随意修改.

很显然这部分应该是函数来实现,于是乎有了下面的改造过程.

![go-oop-about-myDynamicArray-method-inner.png](../images/go-oop-about-myDynamicArray-method-inner.png)

很遗憾,编辑器直接报错: 必须是类型名称或是指向类型名称的指针.

函数不可以放在结构体内,这一点倒是像极了 `C` 家族,但是 `Java` 这种衍生家族会觉得不可思议,不管怎么说,这意味着结构体只能定义结构而不能定义行为!

那我们就把函数移动到结构体外部吧,可是我们定义的函数名叫做 `len`,而系统也有 `len` 函数,此时能否正常运行呢?让我们拭目以待,眼见为实.

![go-oop-about-myDynamicArray-method-len.png](../images/go-oop-about-myDynamicArray-method-len.png)

除了函数本身报错外,函数内部的 `len` 也报错了,是因为此时的函数和结构体尚未建立起任何联系,怎么可能访问到 `len` 属性呢,不报错才怪呢!

解决这个问题很简单,直接将结构体的指针传递给 `len` 函数不就好了,这样一来函数内部就可以访问到结构体的属性了.

![go-oop-about-myDynamicArray-method-len-with-args.png](../images/go-oop-about-myDynamicArray-method-len-with-args.png)

从设计的角度上来讲,确实解决了函数定义的问题,但是使用者调用函数时的使用方法看起来和面向对象的写法有些不一样.

```go
func TestMyDynamicArray(t *testing.T) {
	var myDynamicArray MyDynamicArray

	t.Log(myDynamicArray)

	myDynamicArray.len = 0
	myDynamicArray.cap = 10
	var arr [10]int
	myDynamicArray.ptr = &arr

	t.Log(myDynamicArray)
	t.Log(*myDynamicArray.ptr)

	(*myDynamicArray.ptr)[0] = 1
	t.Log(*myDynamicArray.ptr)

	t.Log(len(&myDynamicArray))
}
```

面向对象的方法中一般都是通过点操作符 `.` 访问属性或方法的,而我们实现的属性访问是 `.` 但方法却是典型的函数调用形式?这看起来明显不像是方法嘛!

为了让普通函数看起来像是面向对象中的方法,`Go` 做了下面的改变,通过将当前结构体的变量声明移动到函数名前面,从而实现类似于面向对象语言中的 `this` 或 `self` 的效果.


```go
func len(myArr *MyDynamicArray) int {
	return myArr.len
}
```

![go-oop-about-myDynamicArray-method-len-ahead-name.png](../images/go-oop-about-myDynamicArray-method-len-ahead-name.png)

此时方法名和参数返回值又报错了,根据提示说函数名和字段名不能相同?

真的又是一件神奇的事情,难不成 `Go` 无法区分函数和字段?这就不得而知了.

那我们只好修改函数名,改成面向对象中喜闻乐见的方法命名规则,如下:

```go
func (myArr *MyDynamicArray) GetLen() int {
	return myArr.len
}
```

> 简单说一下 `Go` 的访问性规则,大写字母开头的表示公开的 `public` 权限,小写字母开头表示私有的 `private` 权限,`Go` 只有这两类权限,都是针对包 `package` 而言,以后会再细说,现在先这么理解就行了.

按照实验得到的方法规则,继续完善其他方法,补充 `GetCap` 和 `IsEmpty` 等方法.

现在我们已经解决了私有变量的访问性问题,对于初始化的逻辑还没有处理,一般来说,初始化逻辑可以放到构造函数中执行,那 `Go` 是否支持构造函数呢,以及怎么才能触发构造函数?

![go-oop-about-myDynamicArray-constuct.png](../images/go-oop-about-myDynamicArray-constuct.png)

尝试按照其他主流的编程语言中构造函数的写法来编写 `Go` 的构造函数 , 没想到 `Go` 编译器直接报错了,提示重新定义了 `MyDynamicArray` 类型,以至于影响了其余部分!

如果修改方法名称的话,理论上可以解决报错问题,但是这并不是构造函数的样子了,难不成 `Go` 不支持构造函数吗?

此时,面向对象形式的构造函数转变成自定义函数实现的构造函数,更加准确的说法,这是一种类似于工厂模式实现的构造函数方式.

```go
func NewMyDynamicArray() *MyDynamicArray {
	var myDynamicArray MyDynamicArray
	return &myDynamicArray
}
```

难道 `Go` 语言真的不支持构造函数?

至于是否支持构造函数或者说应该如何支持构造函数,真相不得而知,随着学习的深入,相信以后一定会有明确的答案,这里简单表达一下个人看法.

首先我们知道 `Go` 的结构体中只能定义数据,而结构体的方法肯定是在结构体外定义的,为了符合面向对象的使用习惯,也就是通过实例对象的点操作符来访问方法,`Go` 的方法只能是函数的变体,即普通函数中关于指向结构体变量的声明部分转移到函数名前面来实现方法,这种由函数转变成为方法的模式也符合 `Go` 一贯的命名规则: 向来是按照人的思维习惯命名,先有输入再有输出等逻辑.

结构体的方法从语法和语义的两个维度上支持了面向对象规范,那么构造函数想要实现面向对象应该如何做呢?

构造函数正如其名应该是函数,而不是方法,方法由指向自身的参数,这一点构造函数不应该有,否则都有实例对象了还构造毛线啊?

既然构造函数是普通函数,那么按照面向对象的命名习惯,方法名应该是结构体名,然而真的操作了,编辑器直接就报错了,所以这不符合面向对象的命名习惯!

如此一来,构造函数的名称可能并不是结构体类型的名称,有可能是其他特殊的名称,最好这个名称能够见名知义且具备实例化对象时自动调用的能力.

当然这个名称依赖于 `Go` 的设计者如何命名,这里靠猜测是很难猜对的,否则我就是设计者了啊!

除此之外,还有另外一种可能,那就是 `Go` 并没有构造函数,想要实现构造函数的逻辑只能另辟蹊径.

这么说有没有什么靠谱的依据呢?

我想大概是有的,构造函数虽然提供了自动初始化能力,但是如果真的在构造函数中加入复杂的初始化逻辑,无疑会增大以后出错的排查难度并给使用者带来一定的阅读障碍,所以说一定程度上,构造函数很有可能被滥用了!

那是否就意味着不需要构造函数了呢?

也不能这么说,构造函数除了基本的变量初始化以及简单的逻辑外,在实际编程中还是有一定用途的,为了避免滥用而直接禁用,多少有点饮鸩止渴的感觉吧?

因此,个人的看法是应该可以保留构造函数这种初始化逻辑,也可以换一种思路去实现,或者干脆直接放弃构造函数转而由编译器自动实现构造函数,正如编译器可以自动添加多字段之间的分号那样.

如果开发者真的有构造函数的需求,通过工厂模式或者单例模式等手段总是可以定制结构体初始化的逻辑,所以放弃也未尝不可!

最后,以上这些纯属个人猜想,目前并不知道 `Go` 是否存在构造函数,有了解的人,还请明确告诉我答案,个人倾向于不存在构造函数,最多只提供类似于构造函数初始化的逻辑!

现在,我们已经封装了结构体的数据,定义了结构体的方法以及实现了结构体的工厂函数.那么接下来让我们继续完善动态数组,实现数组的基本操作.

```go
func NewMyDynamicArray() *MyDynamicArray {
	var myDynamicArray MyDynamicArray

	myDynamicArray.len = 0
	myDynamicArray.cap = 10
	var arr [10]int
	myDynamicArray.ptr = &arr

	return &myDynamicArray
}

func TestMyDynamicArray(t *testing.T) {
	myDynamicArray := NewMyDynamicArray()

	t.Log(myDynamicArray)
}
```

首先将测试用例中的逻辑提取到工厂函数中,默认无参的工厂函数初始化的内部数组长度为 `10` ,后续再考虑调用者指定以及实现动态数组等功能,暂时先实现最基本的功能.

初始化的内部数组均是零值,因此需要首先提供给外界能够添加的接口,实现如下:

```go

func (myArr *MyDynamicArray) Add(index, value int) {
	if myArr.len == myArr.cap {
		return
	}

	if index < 0 || index > myArr.len {
		return
	}

	for i := myArr.len - 1; i >= index; i-- {
		(*myArr.ptr)[i+1] = (*myArr.ptr)[i]
	}

	(*myArr.ptr)[index] = value
	myArr.len++
}
```

由于默认的初始化工厂函数暂时是固定长度的数组,因此新增元素其实是操作固定长度的数组,不过这并不妨碍后续实现动态数组部分.

为了操作方便,再提供插入头部和插入尾部两种接口,可以基于动态数组实现比较高级的数据结构.

```go
func (myArr *MyDynamicArray) AddLast(value int) {
	myArr.Add(myArr.len, value)
}

func (myArr *MyDynamicArray) AddFirst(value int) {
	myArr.Add(0, value)
}
```

为了方便测试动态数组的算法是否正确,因此提供打印方法查看数组结构.

![go-oop-about-myDynamicArray-print.png](../images/go-oop-about-myDynamicArray-print.png)

由此可见,打印方法显示的数据结构和真实的结构体数据是一样的,接下来我们就比较有信心继续封装动态数组了!

```go
func (myArr *MyDynamicArray) Set(index, value int) {
	if index < 0 || index >= myArr.len {
		return
	}

	(*myArr.ptr)[index] = value
}

func (myArr *MyDynamicArray) Get(index int) int {
	if index < 0 || index >= myArr.len {
		return -1
	}

	return (*myArr.ptr)[index]
}
```

这两个接口更加简单,更新数组指定索引的元素以及根据索引查询数组的值.

接下来让我们开始测试一下动态数组的全部接口吧!

![go-oop-about-myDynamicArray-test.png](../images/go-oop-about-myDynamicArray-test.png)

动态数组暂时告一段落,不知道你是否好奇为什么以动态数组为例讲解面向对象?

其实主要是为了验证上一篇文章中的猜想,也就是切片和数组的到底是什么关系?

我觉得切片的底层是数组,只不过语法层面提供了支持以至于看不出数组的影子,仙子阿既然学习了面向对象,那么就用面向对象的方式实现下切片的功能,虽然无法模拟语法级别的实现,但是功能特性完全是可以模仿的啊!

下面还是梳理总结一下本文的只要知识点吧,也就是封装的实现.

## 如何封装结构体

之所以称之为结构体是因为 `Go` 的关键字是 `struct` 而不是 `class`,也是面向对象编程风格中唯一支持的特性,继承和多态都不支持,到时候另开文章细说.

结构体是对数据进行封装所使用的手段,结构体内只能定义数据而不能定义方法,这些数据有时候被称为字段,有时候叫做属性或者干脆叫做变量,至于什么叫法不是特别重要,如何命名和所处的环境语义有关.

```go
type MyDynamicArray struct {
	ptr *[10]int
	len int
	cap int
}
```

这种结构体内就有三个变量,变量之间直接换行进行分隔而不是分号并换行的形式,刚开始觉得有些怪,不过编辑器一般都很智能,假如习惯性添加了分号,会提示你进行删除,所以语法细节上不必在意.

结构体内不支持编写函数,仅支持数据结构,这样就意味着数据和行为是分开的,两者之间的关联是比较弱的.

```go
func (myArr *MyDynamicArray) IsEmpty() bool {
	return myArr.len == 0
}
```

这种方式的函数和普通函数略有不同,将包含结构体变量的参数提前到函数名前面,语义上也比较明确,表示的是结构体的函数,为了和普通函数有所区别,这种函数称之为方法.

其实,单纯地就实现功能上看,方法和函数并没有什么不同,无外乎调用者的使用方式不一样罢了!

```go
func IsEmpty(myArr *MyDynamicArray) bool {
	return myArr.len == 0
}
```

之所以是这种设计方式,一方面体现了函数的重要性,毕竟是 `Go` 语言中的一等公民嘛!

另一方面是为了实现面向对象的语法习惯,不论属性还是方法,统统用点 `.` 操作符进行调用.

官方的文档中将这种结构体参数称之为接收者,因为数据和行为是弱关联的,这里的接收者充当的就是关联数据的作用,接收者顾名思义就是接受数据的人,那发送数据的人又是谁呢?

不言而喻,发送者应该是调用者传递的结构体实例对象,结构体变量将数据结构发送给接收者方法,从而数据和行为联系在一起了.

```go
func TestMyDynamicArray(t *testing.T) {
	myDynamicArray := NewMyDynamicArray()

	fmt.Println(myDynamicArray.IsEmpty())
}
```

好了,以上就是面向对象初体验中的全部部分,仅仅是这么一小部分却耗费我整整三天的时间了,想说换种思维不简单,写好一篇文章也不容易啊!

下篇文章中将继续介绍面向对象的封装特性,讲解更多干货,如果觉得本文对你有所帮助,欢迎转发评论,感觉你的阅读!
