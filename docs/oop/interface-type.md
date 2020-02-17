# 再谈接口

   如果你还了解编程概念中的接口概念,那么我建议你最好还是先阅读上一篇文章.详情请点击 [go 学习笔记之万万没想到宠物店竟然催生出面向接口编程?](https://mp.weixin.qq.com/s/Y2ZfdFUO2QMgWQY5jP0VaA) ,否则的话,请自动忽略上文,继续探索 `Go` 语言的接口有什么不同之处.

> 如无法自动跳转到公众号「雪之梦技术驿站」文章,可以点击我的头像,动动你的小手翻翻历史文章,相信聪明的你一定可以找到相关文章.

   接口是面向对象编程风格中继**封装**概念后的另一个重要概念,封装包含两方面含义:**数据和行为**的封装.

> 关于封装的概念这里同样不再赘述,有兴趣的话,可以阅读[go 学习笔记之详细说一说封装是怎么回事](https://mp.weixin.qq.com/s/fXIKWsPqi6m2IEV--8lBsg).

   当现实世界中的事物或者实际需求转移到编程世界中去实现时,这时候就需要进行建模,建立合适的模型来反映现实的事物,为了模型的紧凑性以及更好的复用性.编程世界的前辈们总结出**封装**的概念,并在此基础上进一步衍生出一系列的编程风格,其中就包括面向对象中的**继承**概念.

> 关于继承的概念这里同样不再赘述,有兴趣的话,可以阅读[go 学习笔记之是否支持以及如何实现继承](https://mp.weixin.qq.com/s/w9ZfTAMr-mnQ9RK_Va-mEg).

   封装和继承都是在描述同类事物模型彼此共性,正如猫和狗都是动物,运用继承的概念表示的话,猫和狗继承自动物.猫和狗不仅具备各自特殊的属性和行为,还具备一般动物的属性和行为.

   然而,并不是只有同类事物才具有相同特征.家禽鸭子是鸭子,玩具太空鸭也是鸭子,看似是同类事物实际却只有某一方面的行为相同而已,一个有生命,另一个无生命.

   针对这种情况下统一共性行为的方法也就是**接口**,是对同类事物或者不同类事物的某一方面行为的统一抽象,满足该行为规范的封装对象称之为实现了该接口.

   接口描述的是规范约束和实现的一种规则,接口定义了这种约束规范,至于如何实现这种规范,接口定义者本身并不关心.如何实现是接口实现者必须关心的,定义者和实现者两者是解耦的.

   从这点来看,接口就像是现实生活中的领导下达命令给下属,下属负责实现目标.如何实现目标,领导并不关心,正所谓条条大路通罗马,手底下的人自然是八仙过海各显神通.

![go-oop-interface-type-all-roads-lead-to-rome.jpeg](../images/go-oop-interface-type-all-roads-lead-to-rome.jpeg)

## 领导关心结果,下属关心实现

   作为领导负责制定各种战略目标,总揽全局关心结果,作为下属负责添砖加瓦实现具体细节关心过程,这种职责分离的模式就是编程语言中接口定义者和接口实现者的关系,一方负责定义行为约束,另一方负责实现这种行为规范.

   **如果站在领导者的角度上看问题**,自然是希望下属**规规矩矩**按时完成自己布置的任务,千万不要出现任何差池,为此甚至会出台一系列的行为准则,签到打卡等形式依次树立领导威望来换取下属的恪尽职责.

   为了达到这个目标,领导者首先要在下属中树立足够高的威信,做到人人信服自己,这样手底下的人才能和自己统一战线一致对外,团结在一起好做事.否则的话,不满嫉妒等负面情绪就会在团队中蔓延,逐渐侵蚀削弱团队战斗力,不攻自破.

![go-oop-interface-type-team-cooperation.jpeg](../images/go-oop-interface-type-team-cooperation.jpeg)

   一般而言,这种威信的树立要么靠的是能力上技高一筹实力碾压,要么是知人善任天下贤才皆为我所用,还可以狐假虎威绿叶衬红花思想上奴役统治.

   不管是什么方式,领导者在这场游戏中占据绝对领导地位,只要上层接口发号施令,下层实现都要随之更改.如果你是领导,相信你也会喜欢这种形式的,毕竟谁心里没有控制欲,更何况是绝对的权力!

   **如果站在下层实现者的角度思考问题**,显然在这场上下级关系中实现者扮演弱势角色,长期忍受不公平的待遇要么崩溃,要么揭竿而起!

   `Go` 语言对于接口的定义者和接口的实现者的关系处理问题上,选择了揭竿而起,实现了不同于其他传统编程规范的另外一种风格规范.

   这种规范常被视为是鸭子类型 `duck typing` --- "当看到一只鸟走起来像鸭子,游泳起来像鸭子,叫起来也像鸭子,那么这只鸟就可以被称为鸭子."

   在这种规范中并不关心结构体对象是什么类型或者说到底是不是鸭子,唯一关心的只是行为.只要满足特定行为的结构体类型就是鸭子类型,哪怕这种鸭子可能只是一种玩具也行!所以,在这种接口定义者和实现者的关系中,实现者可以不必向接口特意声明实现,只要最终行为上确实实现了接口中定义的行为规范,那么就称为该结构体实现了接口.

   **如果仅仅考虑接口定义者和实现者的关系**,基于这种关系很容易进行下一步推断,要么实现者一定要声明实现接口,随时向领导汇报工作进度,要么一定不声明接口,只要保证最终能够完成任务即可.除此之外,很明显还存在另外一种可能性,那就是实现者可以选择报告工作也可以选择不报告.

   那么,这种似是而非的关系是否有存在的意义呢,又该如何表示呢以及有没有现成编程语言基于此思路实现呢?

   按照基本语义进行理解推测: 实现者需要报告给接口的方法一定是万分紧急十分重要的规范,正所谓大是大非面前不能有任何个人情感,一旦实现者无法实现,那么便不可饶恕,零容忍!

   如果实现者不报告给接口,则表示这种规范是可选规范,如果满足的话,自然是好的.如果有特殊情况一时没能实现也不算是致命的问题,这类规范是可选规范,属于锦上添花的操作.

   所以要描述这种**可有可无**的接口定义者和实现者的关系,显而易见的是,理应由**接口定义者**来指明接口的优先级,不能由实现者定义.否则的话,你认为爱国是必选的,他认为是可选的,那么接口的存在还有什么意义?既然如此,接口方法在声明时就应该声明该接口方法是必选的还是可选的,这样实现者实现该接口时才能有理可循,对于必选实现的接口只要没实现就不算是真正的接口实现者,而可选的接口允许实现者可以暂时不实现.

   由于个人知识经验所限,暂不可知有没有现成的编程语言支持这种妥协状态,接口方法既可以声明必选的也可以声明可选的.个人觉得这种方式还是比较友好的,还是有存在的价值的.

> 如果你知道有什么编程语言刚好是这种思路实现了接口规范,还望不吝赐教,可以留言评论相互学习下.

## 理论指导实践,实践中出真知

   虽然猜测中的第三种规范是介于**必须上报和必须不上报**之间的妥协状态,但是由于接口声明时有可选和必选之分,这种区分需要有接口定义者进行指定,因此在接口和实现者的关系中还是**接口定义者**占据主导地位.

   当接口定义者占据主导地位时,现成的最佳编程实践告诉我们先定义接口再写实现类,也就是**先有规范再写实现**,所以实际编程中给我们的指导就是先抽象出共同行为,定义出接口规范,再去写不同的实现类去实现该接口,当使用接口时就可以不区分具体的实现类直接调用接口本身了.

   如果有一句话来描述这种行为的话,那就是**理论指导实践,先写接口再写实现**.

   同样的,我们还知道另外一句话,这就是**实践出真知**,这种思路刚好也是比较符合现实的,先写所谓的实现类,当这种实现类写的比较多的时候,就如继承那样,自然会发现彼此之间的关联性,再抽象成接口也是水到渠成的事情,不必在编程刚开始就费时费力去抽象定义接口等高级功能特性.

   通过上篇文章关于 `Go` 语言的接口的设计思想我们知道 `Go` 语言采用的就是后一种: 实践中出真知.
接口实现者对于接口的实现是隐式的,也就是说某一种结构体很有可能有意无意实现了某种接口,真的是有心插花花不开,无心插柳柳成荫.

![go-oop-interface-type-miracle-by-chance.jpeg](../images/go-oop-interface-type-miracle-by-chance.jpeg)

## 应如何区分有没有无心插柳

   `Go` 语言这种似是而非若有还无的朦胧暧昧既给我们带来了方便,同时也给我们留下了些许烦恼,假如需要知道结构体类型到底是不是接口的实现者时,反而有些费事了.

   值得庆幸的是,现代 `IDE` 一般都比较智能,这种接口语法虽然比较灵活但还是有规律可寻的,所以一般 `IDE` 也是可以智能推测出接口和实现的关系的,并不用我们肉眼去仔细辨别.

![go-oop-interface-type-ide-instruction.png](../images/go-oop-interface-type-ide-instruction.png)

> `Programmer` 接口的左侧有个向下的箭头,而 `GoProgrammer` 结构体类型左侧有个向上箭头.此时鼠标点击箭头可以相互跳转,这就是 `IDE` 提供的可视化效果.

   如果真的需要在程序中辨别接口和实现类的关系,那么只能借助系统级别的方法来判断了,准备环境如下:

首先先定义程序员的第一课 `Hello World` 的接口:

```go
type Programmer interface {
    WriteHelloWord() string
}
```

然后按照不同的编程语言实现该接口,为了更加通用性表示 `WriteHelloWord` 的输出结果,这里将输出结果 `string` 定义成别名形式以此表示输出的是代码 `Code`.

```go
type Code string
```

按照 `Code` 别名重新整理接口定义,如下:

```go
type Programmer interface {
    WriteHelloWord() Code
}
```

接下来我们用 `Go` 语言写第一个程序,而 `Go` 实现接口的方式是隐式的,并不需要关键字强制声明.

```go
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWord() Code {
    return "fmt.Println(\"Hello World!\")"
}
```

然后,选择 `Java` 程序员作为对比,其他面向对象编程语言类似,这里不再赘述.

```go
type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWord() Code {
    return "System.out.Println(\"Hello World!\")"
}
```

当用户需要程序员写 `WriteHelloWord` 程序时,此时 `Go` 程序员和 `Java` 程序员准备各显身手,比较简单,这里重点是看一下**接口变量**的类型和值.

```go
func writeFirstProgram(p Programmer) {
    fmt.Printf("%[1]T %[1]v %v\n", p, p.WriteHelloWord())
}
```

按照接口的语义,我们可以将 `Go` 程序员和 `Java` 程序员全部扔给 `writeFirstProgram` 方法中,此时接口的类型是**具体实现类的类型**,接口的值也是**实现类的数据**.

当然,不论是 `Go` 还是 `Java` 都可以写出 `WriteHelloWord` .

```go
func TestPolymorphism(t *testing.T) {
    gp := new(GoProgrammer)
    jp := new(JavaProgrammer)

    // *polymorphism.GoProgrammer &{} fmt.Println("Hello World!")
    writeFirstProgram(gp)
    // *polymorphism.JavaProgrammer &{} System.out.Println("Hello World!")
    writeFirstProgram(jp)
}
```

上述例子很简单,我们自然也是可以一眼看出接口和实现类的关系,并且 `IDE` 也为我们提供非常直观的效果,在比较复杂的结构体中这种**可视化效果**尤为重要.

![go-oop-interface-type-programer.png](../images/go-oop-interface-type-programer.png)

   如果你非要和我较真,说你正在用的 `IDE` 无法可视化直接看出某个类型是否满足某接口,又该怎么办?

**我的建议是**,那就换成和我一样的 `IDE` 不就好了吗!

   哈哈,这只不过是我的一厢情愿罢了,有些人是不愿意改变的,不会随随便便就换一个 `IDE`,那我就告诉你另外一个方法来检测类型和接口的关系.

赵本山说,**没事你就走两步**?

![go-oop-interface-type-try-to-go-walk.jpg](../images/go-oop-interface-type-try-to-go-walk.jpg)

   真的是博大精深,言简意赅!如果某个结构体类型满足特定接口,那么这个这个结构体的实例化后一定可以赋值给接口类型,如果不能则说明肯定没有实现!肉眼看不出的关系,那就拿放大镜看,编译错误则不符合,编译通过则满足.

   为了对比效果,这里再定义一个新的接口 `MyProgrammer` ,除了名称外,接口暂时和 `Programmer` 完全一样.

![go-oop-interface-type-myProgrammer-pass.png](../images/go-oop-interface-type-myProgrammer-pass.png)

`IDE` 并没有报错,左侧的可视化效果也表明 `MyProgrammer` 和 `Programmer` 虽然名称不同,但是接口方法却一模一样,`GoProgrammer` 类型不仅实现了原来的 `Programmer` 接口还顺便实现了 `MyProgrammer`.

不仅 `GoProgrammer` 是这样,`JavaProgrammer` 也是如此,有意无意实现了新的接口,这也就是 `Go` 的接口设计不同于传统声明式接口设计的地方.

![go-oop-interface-type-myProgrammer-goProgrammer.png](../images/go-oop-interface-type-myProgrammer-goProgrammer.png)

   现在我们改变一下 `MyProgrammer` 接口中的 `WriteHelloWord` 方法,返回类型由别名 `Code` 更改成原类型 `string`,再试一下实际效果如何.

由于 `Go` 是强类型语言,即使是别名和原类型也不是相同的,正如类型之间的转换都是强制的,没有隐式类型转换那样.

因此,可以预测的是,`WriteHelloWord` 接口方法前后不一致,是没有类型结构体满足新的接口方法的,此时编译器应该会报错.

![go-oop-interface-type-myProgrammer-fail.png](../images/go-oop-interface-type-myProgrammer-fail.png)

   事实胜于雄辩,无论是 `GoProgrammer` 还是 `JavaProgrammer` 都没有实现 `MyProgrammer` ,因此是不能赋值给类型 `MyProgrammer` ,编译器确实报错了!

并不是所有长得像的都是兄弟,也不是长得不像的就不是兄弟.

```go
type Equaler interface {
    Equal(Equaler) bool
}
```

`Equaler` 接口定义了 `Equal` 方法,不同于传统的多态,`Go` 的类型检查更为严格,并不支持多态特性.

```go
type T int

func (t T) Equal(u T) bool { return t == u }
```

如果单单看 `Equal(u T) bool` 方法声明,放到其他主流的编程语言中这种情况可能是正确的,但是多态特性并不适合 `Go` 语言.

![go-oop-interface-type-equal-fail.png](../images/go-oop-interface-type-equal-fail.png)

> 不仅仅 `IDE` 没有左侧**可视化的箭头效果**,硬生生的将类型声明成接口类型也会报错,说明的确没有实现接口.

   透过现象看本质,`T.Equal` 的参数类型是`T` ,而不是字面上所需的类型`Equaler`,所以并没有实现 `Equaler` 接口中规定的 `Equal` 方法.

   是不是很意外?

![go-oop-interface-type-surprise.png](../images/go-oop-interface-type-surprise.png)

  如果你已经看到了这里,相信你现在不仅基本理解了面向对象的三大特性,还知道了 `GO` 设计的是多么与众不同!

  这种与众不同之处,不仅仅体现在面向对象中的类型和接口中,最基础的语法细节上无一不体现出设计者的匠心独运,正是这种创新也促进我们重新思考面向对象的本质,真的需要循规蹈矩按照现有的思路去设计新语言吗?

  `Go` 语言的语法精简,设计简单优雅,抛弃了某些看起来比较高级但实际使用过程中可能会比较令人困惑的部分,对于这部分的舍弃,确实在一定程度上简化了整体的设计.

   但是另一方面,如果仍然需要这种被丢弃的编程习惯时,只能由开发者手动实现,从这点看就不太方便了,所以只能尽可能靠近设计者的意图,写出真正的 `Go` 程序.

   控制权的转移意味着**开发者**承担了更多的责任,比如类型转换中没有显式类型转换和隐式类型转换之分,`Go` 仅仅支持**显式类型转换**,不会自动帮你进行隐式转换,也没有为了兼顾隐式类型的转换而引入的基本类型的**包装类型**,也就没有**自动拆箱和自动装箱**等复杂概念.

   所以如果要实现 `Equal` 接口方法,那么就应该开发者自己保证严格实现,这里只需要稍微修改下就能真正实现该方法.

```go
type T2 int

func (t T2) Equal(u Equaler) bool { return t == u.(T2) }
```

`Equal(Equaler) bool` 接口方法中的参数中要求 `Equaler` 接口,因此 `Equal(u Equaler) bool` 方法才是真正实现了接口方法.

![go-oop-interface-type-equal-pass.png](../images/go-oop-interface-type-equal-pass.png)

只有方法名称和签名完全一致才是实现了接口,否则看似实现实则是其他编程语言的逻辑,放到`Go` 语言中并没有实现接口.

## 如何保证实现者是特定类型

   但是不知道你是否发现,这种形式实现的接口方法和我们熟悉的面向接口编程还是有所不同,任何满足接口 `Equaler` 方法的类型都可以被传入到 `T2.Equal` 的参数,而我们的编译器却不会在编译时给出提示.

```go
type T3 int

func (t T3) Equal(u Equaler) bool { return t == u.(T3) }
```

仿造 `T2` 实现 `T3` 类型,同样也实现了 `Equaler` 接口所要求的 `Equal` 方法.

`T2` 和 `T3` 明显是不同的类型,编译期间 `T3` 是可以传给 `T2` 的,反之亦然, `T2` 也可以传给 `T3` .

![go-oop-interface-type-equal-error-pass.png](../images/go-oop-interface-type-equal-error-pass.png)

   编译正常而运行出错意味着后期捕捉问题的难度加大了,个人比较习惯于编译期间报错而不是运行报错,`Go` 语言就是编译型语言为什么造成了编译期间无法捕捉错误而只能放到运行期间了?

![go-oop-interface-type-equal-error-panic.png](../images/go-oop-interface-type-equal-error-panic.png)

   由此可见,`t == u.(T3)` 可能会抛出异常,异常机制也是编程语言通用的一种自我保护机制,`Go` 语言应该也有一套机制,后续再研究异常机制,暂时不涉及.

   不过我们在这里确实看到了 `u.(T3)` 判断类型的局限性,想要确保程序良好运行,应该研究一下接口变量到底是什么以及如何判断类型和接口的关系.

   编译期间的判断关系可以通过 ide 的智能提示也可以将类型声明给接口看看是否编译错误,但这些都是编译期间的判断,无法解决当前运行期间的错误.

```go
func TestEqualType(t *testing.T) {
    var t2 Equaler = new(T2)
    var t3 Equaler = new(T3)

    t.Logf("%[1]T %[1]v\n",t2)
    t.Logf("%[1]T %[1]v\n",t3)
    t.Logf("%[1]T %[1]v %v\n",t2,t2.Equal(t3))
}
```

> `%T %V` 打印出接口变量的类型和值,从输出结果上看 `*polymorphism.T2 0xc0000921d0`,我们得知接口变量的类型其实就是实现了该接口的结构体类型,接口变量的值就是该结构体的值.

`t2` 和 `t3` 接口变量的类型因此是不同的,运行时也就自然报错了.

说完现象找原因: `Go` 语言的接口并没有保证实现接口的类型具有多态性,仅仅是约束了统一的行为规范,`t2` 和 `t3` 都满足了 `Equal` 这种规范,所以对于接口的设计效果来说,已经达到目标了.

   但是这种接口设计的理念和我们所熟悉的其他编程语言的多态性是不同的,`Go` 并没有多态正如没有继承特性一样.

```go
func TestInterfaceTypeDeduce(t *testing.T) {
    var t2 Equaler = new(T2)
    var t3 Equaler = new(T3)

    t.Logf("%[1]T %[1]v %[2]T %[2]v\n",t2,t2.(*T2))
    t.Logf("%[1]T %[1]v %[2]T %[2]v\n",t3,t3.(*T3))
}
```

![go-oop-interface-type-equal-type-deduce.png](../images/go-oop-interface-type-equal-type-deduce.png)

   当 `t2.(*T2)` 或 `t3.(*T3)` 时,均正常工作,一旦 `t2.(*T3)` 则会抛出异常,因此需要特殊处理下这种情况.

   根据实验结果得知,`t2.(*T2)` 的类型和值恰巧就是接口变量的类型和值,如果结构体类型不能转换成指定接口的话,则可能抛出异常.

   因此,猜测这种形式的效果上类似于强制类型转换,将接口变量 `t2` 强制转换成结构体类型,动不动就报错或者说必须指定接口变量和结构体类型的前提,有点像其他编程语言的断言机制.

   单独研究一下这种断言机制,按照 `Go` 语言函数设计的思想,这种可能会抛出异常的写法并不是设计者的问题,而是我们使用者的责任,属于使用不当,没有检查能否转换成功.

```go
v2,ok2 := t2.(*T2)
```

   从实际运行的结果中可以看出,接口变量 `t2` 经过断言为 `*T2` 结构体类型后得到的变量和接口变量 `t2` 应该是一样的,因为他俩的类型和值完全一样.

   当这种转换失败时,`ok` 的值是 `false` ,此时得到的转换结果就是 `nil` .

![go-oop-interface-type-type-deduce.png](../images/go-oop-interface-type-type-deduce.png)

## 老子口中的无为而治空接口

   接口既然是实现规范的方式,按照以往的编程经验给我们的最佳实践,我们知道接口最好尽可能的细化,最好一个接口中只有一个接口方法,足够细分接口即减轻了实现者的负担也方便复杂接口的组合使用.

   有意思的是,`Go` 的接口还可以存在**没有任何接口方法**的空接口,这种特殊的接口叫做**空接口**,无为而治,没有任何规范约束,这不就是老子口中的**顺其自然,无为而治**吗?

```go
type EmptyInterface interface {
}
```

   道家的思想主要靠领悟,有点哲学的味道,这一点不像理科知识那样严谨,可以根据已知按照一定的逻辑推测出未知,甚至预言出超时代的新理论也不是没有可能的.

   然而,道家说**一生二,二生三,三生万物**,这句话看似十分富有哲理性但是实际却很难操作,只讲了开头和结尾,并没有讲解如何生万物,忽略了过程,全靠个人领悟,这就很难讲解了.

![go-oop-interface-type-dao-empty.jpg](../images/go-oop-interface-type-dao-empty.jpg)

   没有任何接口方法的空接口和一般接口之间是什么关系?

空接口是一,是接口中最基础的存在,有一个接口的是二,有二就会有三,自然就会有千千万万的接口,从而构造出接口世界观.

```go
func TestEmptyInterfaceTypeDeduce(t *testing.T) {
    var _ Programmer = new(GoProgrammer)
    var _ EmptyInterface = new(GoProgrammer)
}
```

`GoProgrammer` 结构体类型不仅实现了 `Programmer` 接口,也实现空接口,至少编译级别没有报错.

   但是,`Go` 语言的接口实现是**严格实现**,空接口没有接口,因此没有**任何结构体都没有实现空接口**,符合一贯的设计理念,并没有特殊处理成默认实现空接口.

![go-oop-interface-type-empty-interface-not-implement.png](../images/go-oop-interface-type-empty-interface-not-implement.png)

   所以我困惑了,一方面,结构体类型实例对象可以赋值给空接口变量,而结构体类型却又没法实现空接口,这不是有种**自相矛盾**的地方吗?

## 莫非是继承不足空接口来凑

   明明没有实现空接口却可以赋值给空接口,难不成是为了弥补语言设计的不足?

因为 `Go` 语言不支持继承,自然没有其他编程语言中的**基类**概念,而实际工作中有时候确实需要一种通用的封装结构,难道是继承不足,接口来凑?

所以设计出空接口这种特殊情况来弥补没有继承特性的不足?有了空接口就有了 `Go` 语言中的 `Object` 和泛型 `T` ,不知道这种理解对不对?

```go
func TestEmptyInterface(t *testing.T) {
    var _ Programmer = new(GoProgrammer)
    var _ EmptyInterface = new(GoProgrammer)
    var p EmptyInterface = new(GoProgrammer)

    v, ok := p.(GoProgrammer)
    t.Logf("%[1]T %[1]v %v\n", v, ok)
}
```

   空接口的这种特殊性值得我们花时间去研究一下,因为任何结构体类型都可以赋值给空接口,那么此时的接口变量断言出结构体变量是否也有配套的特殊之处呢?

```go
func TestEmptyInterfaceTypeDeduce(t *testing.T) {
    var gpe EmptyInterface = new(GoProgrammer)

    v, ok := gpe.(Programmer)
    t.Logf("%[1]T %[1]v %v\n", v, ok)

    v, ok = gpe.(*GoProgrammer)
    t.Logf("%[1]T %[1]v %v\n", v, ok)

    switch v := gpe.(type) {
    case int:
        t.Log("int", v)
    case string:
        t.Log("string", v)
    case Programmer:
        t.Log("Programmer", v)
    case EmptyInterface:
        t.Log("EmptyInterface", v)
    default:
        t.Log("unknown", v)
    }
}
```

   虽然接收的时候可以接收任何类型,但是实际使用过程中必须清楚知道具体类型才能调用实例化对象的方法,因而这种**断言**机制十分重要.

```go
func doSomething(p interface{}) {
    if i, ok := p.(int); ok {
        fmt.Println("int", i)
        return
    }
    if s, ok := p.(string); ok {
        fmt.Println("string", s)
        return
    }
    fmt.Println("unknown type", p)
}

func TestDoSomething(t *testing.T) {
    doSomething(10)
    doSomething("10")
    doSomething(10.0)
}
```

当然上述 `doSomething` 可以采用 `switch` 语句进行简化,如下:

```go
func doSomethingBySwitch(p interface{}) {
    switch v := p.(type) {
    case int:
        fmt.Println("int", v)
    case string:
        fmt.Println("string", v)
    default:
        fmt.Println("unknown type", v)
    }
}

func TestDoSomethingBySwitch(t *testing.T) {
    doSomethingBySwitch(10)
    doSomethingBySwitch("10")
    doSomethingBySwitch(10.0)
}
```

## 不一样的接口基本用法总结

- 类型别名

```go
type Code string
```

> `Code` 类型是原始类型 `string` 的别名,但 `Code` 和 `string` 却不是完全相等的,因为 `Go` 不存在隐式类型转换,`Go` 不认为这两种类型是一样的.

- 接口定义者

```go
type Programmer interface {
	WriteHelloWord() Code
}
```

> `Programmer` 接口定义了 `WriteHelloWord()` 的方法.

- 接口实现者

```go
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWord() Code {
	return "fmt.Println(\"Hello World!\")"
}
```

> `Go` 开发者实现了 `WriteHelloWord` 接口方法,而这个方法刚好是 `Programmer` 接口中的唯一一个接口方法,因此 `GoProgrammer` 也就是 `Programmer` 接口的实现者.

这种**基于方法推断**出实现者和定义者的形式和其他主流的编程语言有很大的不同,这里并没有显示声明结构体类型需要实现什么接口,而是说干就干,可能**一不小心就实现了某种接口**都有可能.

```go
type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWord() Code {
	return "System.out.Println(\"Hello World!\")"
}
```

> 此时,当然是我们故意实现了 `Programmer` 接口,以便接下来方便演示接口的基于用法.

- 接口的使用者

```go
func writeFirstProgram(p Programmer) {
	fmt.Printf("%[1]T %[1]v %v\n", p, p.WriteHelloWord())
}
```

> 定义了 `writeFirstProgram` 的函数,接收 `Programmer` 接口类型的参数,而接口中定义了 `WriteHelloWord` 的接口方法.

所以不管是 `GoProgrammer` 还是 `JavaProgrammer` 都可以作为参数传递给 `writeFirstProgram` 函数,这就是面向接口编程,并不在乎具体的实现者,只关心接口方法足矣.

- 面向接口编程

```go
func TestPolymorphism(t *testing.T) {
	gp := new(GoProgrammer)
	jp := new(JavaProgrammer)

	// *polymorphism.GoProgrammer &{} fmt.Println("Hello World!")
	writeFirstProgram(gp)
	// *polymorphism.JavaProgrammer &{} System.out.Println("Hello World!")
	writeFirstProgram(jp)
}
```

> 传递给 `writeFirstProgram` 函数的参数中如果是 `GoProgrammer` 则实现 `Go` 语言版本的 `Hello World!`,如果是 `JavaProgrammer` 则是 `Java` 版本的 `System.out.Println("Hello World!")`

- 看似松散实则依旧严格的接口实现规则

```go
type MyProgrammer interface {
	WriteHelloWord() string
}
```

![go-oop-interface-type-alias-not-implement.png](../images/go-oop-interface-type-alias-not-implement.png)

> `MyProgrammer` 和 `Programmer` 中的 `WriteHelloWord` 接口方法只有返回值类型不一样,虽然`Code` 类型是 `string` 类型的别名,但是 `Go` 依旧不认为两者相同,所以 `JavaProgrammer` 不能赋值给 `MyProgrammer` 接口类型.

- 接口变量肚子里是藏了啥

```go
type GoProgrammer struct {
	name string
}

type JavaProgrammer struct {
	name string
}
```

> 给接口实现者添加 `name` 属性,其余不做改变.

```go
func interfaceContent(p Programmer) {
	fmt.Printf("%[1]T %[1]v\n", p)
}

func TestInterfaceContent(t *testing.T) {
	var gp Programmer = &GoProgrammer{
		name:"Go",
	}
	var jp Programmer = &JavaProgrammer{
		name:"Java",
	}

	// *polymorphism.GoProgrammer &{Go}
	interfaceContent(gp)
	// *polymorphism.JavaProgrammer &{Java}
	interfaceContent(jp)
}
```

> 输出接口变量的类型和值,结果显示接口变量的类型就是结构体实现者的类型,接口变量的值就是实现者的值.

```go
func (g GoProgrammer) PrintName()  {
	fmt.Println(g.name)
}

func (j JavaProgrammer) PrintName()  {
	fmt.Println(j.name)
}
```

现在继续添加结构体类型的方法,可能 `PrintName` 方法有意无意实现了某种接口,不过在演示项目中肯定没有实现接口.

从实验中我们知道接口变量的类型和值都是实现者的类型和值,那么能否通过接口变量访问到实现者呢?

想要完成访问实现者的目标,首先需要知道具体实现者的类型,然后才能因地制宜访问具体实现者的方法和属性等.

- 断言判断接口变量的实现者

```go
func TestInterfaceTypeImplMethod(t *testing.T) {
	var gp Programmer = &GoProgrammer{
		name: "Go",
	}

	// *polymorphism.GoProgrammer &{Go}
	fmt.Printf("%[1]T %[1]v\n", gp)

	if v, ok := gp.(*GoProgrammer); ok {
		// Go
		v.PrintName()
	}else{
		fmt.Println("gp is not *GoProgrammer")
	}
}
```

> `v, ok := gp.(*GoProgrammer)` 将接口变量转换成结构体类型,如果转换成功意味着断言成功,则可以调用相应结构体类型实例对象的方法和属性.如果断言失败,则不可以.

- 空接口定义和使用

```go
type EmptyInterface interface {

}
```

> 任何结构体类型都可以赋值给空接口,此时空接口依旧和一般接口一样的是可以采用断言机制确定目标结构体类型.

但这并不是最常用的操作,比较常用的做法还是用来充当类似于 `Object` 或者泛型的角色,空接口可以接收任何类型的参数.

```go
func emptyInterfaceParam(p interface{}){
	fmt.Printf("%[1]T %[1]v",p)

	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case Programmer:
		fmt.Println("Programmer", v)
	case EmptyInterface:
		fmt.Println("EmptyInterface", v)
	default:
		fmt.Println("unknown", v)
	}
}

func TestEmptyInterfaceParam(t *testing.T) {
	var gp Programmer = new(GoProgrammer)
	var ge EmptyInterface = new(GoProgrammer)

	// *polymorphism.GoProgrammer &{}Programmer &{}
	emptyInterfaceParam(gp)
	
	// *polymorphism.GoProgrammer &{}Programmer &{}
	emptyInterfaceParam(ge)
}
```

好了,关于 `Go` 语言的接口部分暂时结束了,关于面向对象编程风格的探索也告一段落,接下来将开始探索 `Go` 的一等公民函数以及函数式编程.敬请期待,希望学习路上,与你同行!

![go-oop-interface-type-thank_you.png](../images/go-oop-interface-type-thank_you.png)

- [猜猜看go是不是面向对象语言？能不能面向对象编程？](https://mp.weixin.qq.com/s/l-nqzYjcre2g__BrVptR9A)
- [go 学习笔记之详细说一说封装是怎么回事](https://mp.weixin.qq.com/s/fXIKWsPqi6m2IEV--8lBsg)
- [go 学习笔记之是否支持以及如何实现继承](https://mp.weixin.qq.com/s/w9ZfTAMr-mnQ9RK_Va-mEg)
- [go 学习笔记之万万没想到宠物店竟然催生出面向接口编程?](https://mp.weixin.qq.com/s/Y2ZfdFUO2QMgWQY5jP0VaA)

> 上述列表是关于 `Go` 语言面向对象的全部系列文章,详情见微信公众号「雪之梦技术驿站」,如果本文对你有所帮助,欢迎转发分享,如有描述不当之处,请一定要留言评论告诉我,感谢~
