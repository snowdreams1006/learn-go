# 数组切片

字符串

- string 是数据类型,不是引用或指针类型
- string 是只读的 byte slice,len 函数可以返回它所包含的 byte 数
- string 的 byte 数组可以存放任何数组

```go
func TestString(t *testing.T) {
    var s string

    // s = ,len(s) =  0
    t.Logf("s = %v,len(s) =  %d", s, len(s))

    s = "hello"

    // s = hello,len(s) =  5
    t.Logf("s = %v,len(s) =  %d", s, len(s))

    s = "\xE4\xB8\xA5"

    // s = 严,len(s) =  3
    t.Logf("s = %v,len(s) =  %d", s, len(s))

    s = "\xE444\xB888\xA555"

    // s = �44�88�55,len(s) =  9
    t.Logf("s = %v,len(s) =  %d", s, len(s))
}
```

unicode utf8

- unicode 是一种字符集(code point)
- utf8 是 unicode 的存储实现(转换为字节序列的规则)

```go
func TestStringCodePoint(t *testing.T) {
    var s string

    // s = ,len(s) =  0
    t.Logf("s = %v,len(s) =  %d", s, len(s))

    s = "hello"

    // s = hello,len(s) =  5
    t.Logf("s = %v,len(s) =  %d", s, len(s))

    s = "中"

    // s = 中,len(s) =  3
    t.Logf("s = %v,len(s) =  %d", s, len(s))

    //s = e4b8ad,len(s) =  3
    t.Logf("s = %x,len(s) =  %d", s, len(s))

    ch := []rune(s)

    // s = [4e2d],len(s) =  1
    t.Logf("ch = %x,len(ch) =  %d", ch, len(ch))
}
```

```go
func TestString2Rune(t *testing.T) {
    s := "中华人民共和国"

    for _, c := range s {
        t.Logf("%[1]c %[1]d %[1]x", c)
    }
}
```

常用字符串函数

- strings 包

```go
func TestStrings(t *testing.T) {
    s := "A,B,C"

    parts := strings.Split(s,",")
    for _, part := range parts {
        t.Log(part)
    }

    newStr := strings.Join(parts,"-")

    t.Log(newStr)
}
```

- strconv 包

```go
func TestStrconv(t *testing.T) {

    s := strconv.Itoa(10)

    // string 10
    t.Logf("%[1]T %[1]v", s)

    // int 10
    if i, err := strconv.Atoi("10"); err == nil {
        t.Logf("%[1]T %[1]v", i)
    }
}
```

函数是一等公民

- 它可以返回多个值
- 所有参数都是值传递: `map` ,`slice`,`channel` 会有引用传递的错觉
- 函数可以作为变量的值
- 函数可以作为参数和返回值

```go
func returnMultiValues() (int, int) {
    return rand.Intn(10), rand.Intn(10)
}

func TestReturnMultiValues(t *testing.T) {
    a, b := returnMultiValues()
    
    t.Log(a, b)
}
```

```go
func timeSpend(inner func(op int) int) func(op int) int {
    return func(n int) int {
        start := time.Now()

        ret := inner(n)

        fmt.Println("time spend : ", time.Since(start).Seconds())

        return ret
    }
}

func slowFunc(op int) int {

    time.Sleep(time.Second * 1)

    return op
}

func TestSlowFuncTimeSpend(t *testing.T){

    slowFuncTimeSpend := timeSpend(slowFunc)

    t.Log(slowFuncTimeSpend(10))
}
```

```go
func TestDefer(t *testing.T) {
    defer func() {
        t.Log("Clear resource")
    }()

    t.Log("Started")

    panic("Fatal error")
}
```

[https://golang.google.cn/doc/faq#Is_Go_an_object-oriented_language](https://golang.google.cn/doc/faq#Is_Go_an_object-oriented_language)

封装数据和行为

```go

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

```go
func (e Employee) toString() string {
    return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToString(t *testing.T) {
    e := Employee{"0", "Bob", 20}

    t.Log(e.toString())
}
```

```go

func (e *Employee) toStringPointer() string {
    return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToStringPointer(t *testing.T) {
    e := &Employee{"0", "Bob", 20}

    t.Log(e.toStringPointer())
}
```

```go
func (e *Employee) toStringPointer() string {
    fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

    return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToStringPointer(t *testing.T) {
    e := Employee{"0", "Bob", 20}

    fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

    t.Log(e.toStringPointer())
}
```

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

```go
type Programmer interface {
    WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
    return "fmt.Println(\"Hello World\")"
}

func TestGoProgrammer(t *testing.T) {
    var p Programmer
    p = new(GoProgrammer)
    
    t.Log(p.WriteHelloWorld())
}
```

- 接口为非入侵,实现不依赖于接口定义
- 所以接口的定义可以包含在接口的使用者包内

```go
type Pet struct {
}

func (p *Pet) Speak() {
    fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
    p.Speak()
    fmt.Println(" ", host)
}

type Dog struct {
    p *Pet
}

func (d *Dog) Speak() {
    d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
    d.p.SpeakTo(host)
}

func TestDog(t *testing.T){
    d := new(Dog)
    d.SpeakTo("Chao")
}
```

```go
type Code string

type Programmer interface {
    WriteHelloWord() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWord() Code {
    return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWord() Code {
    return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
    fmt.Printf("%T %v\n", p, p.WriteHelloWord())
}

func TestPolymorphism(t *testing.T) {
    gp := new(GoProgrammer)
    jp := new(JavaProgrammer)

    writeFirstProgram(gp)
    writeFirstProgram(jp)
}
```

空接口与断言

- 空接口可以表示任何类型
- 通过断言来将空接口转换为指定类型

`v,ok := p.(int)`

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
    fmt.Println("unknown type")
}

func TestDoSomething(t *testing.T) {
    doSomething(10)
    doSomething("10")
    doSomething(10.0)
}
```

接口最佳实践

- 倾向于使用小的接口定义,很多接口只包含一个方法
- 较大的接口定义可以由小的接口组合而成
- 只依赖于必要功能的最小接口

>>>>>>> a38c4ca339dd38f436a45102878e92135d4b6d3b
