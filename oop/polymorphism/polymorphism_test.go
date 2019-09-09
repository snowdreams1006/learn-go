package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWord() Code
}

type GoProgrammer struct {
	name string
}

func (g *GoProgrammer) WriteHelloWord() Code {
	return "fmt.Println(\"Hello World!\")"
}

func (g GoProgrammer) PrintName()  {
	fmt.Println(g.name)
}

type JavaProgrammer struct {
	name string
}

func (j *JavaProgrammer) WriteHelloWord() Code {
	return "System.out.Println(\"Hello World!\")"
}

func (j JavaProgrammer) PrintName()  {
	fmt.Println(j.name)
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%[1]T %[1]v %v\n", p, p.WriteHelloWord())
}

func TestPolymorphism(t *testing.T) {
	gp := new(GoProgrammer)
	jp := new(JavaProgrammer)

	// *polymorphism.GoProgrammer &{} fmt.Println("Hello World!")
	writeFirstProgram(gp)
	// *polymorphism.JavaProgrammer &{} System.out.Println("Hello World!")
	writeFirstProgram(jp)
}

type MyProgrammer interface {
	WriteHelloWord() string
}

func TestInterfaceType(t *testing.T) {
	var _ Programmer = new(GoProgrammer)
	//var _ MyProgrammer = new(JavaProgrammer)
}

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

type Equaler interface {
	Equal(Equaler) bool
}

type T int

func (t T) Equal(u T) bool { return t == u }

func TestEqualFail(t *testing.T) {
	//var _ Equaler = T{}
}

type T2 int

func (t T2) Equal(u Equaler) bool { return t == u.(T2) }

func TestEqualPass(t *testing.T) {
	var _ Equaler = (*T2)(nil)
}

type T3 int

func (t T3) Equal(u Equaler) bool { return t == u.(T3) }

func TestEqualType(t *testing.T) {
	var t2 Equaler = new(T2)
	var t3 Equaler = new(T3)

	t.Logf("%[1]T %[1]v\n", t2)
	t.Logf("%[1]T %[1]v\n", t3)
	t.Logf("%[1]T %[1]v %v\n", t2, t2.Equal(t3))
}

func TestInterfaceTypeDeduce(t *testing.T) {
	var t2 Equaler = new(T2)
	var t3 Equaler = new(T3)

	t.Logf("%[1]T %[1]v %[2]T %[2]v\n", t2, t2.(*T2))
	t.Logf("%[1]T %[1]v %[2]T %[2]v\n", t3, t3.(*T3))

	v2, ok2 := t2.(*T2)
	t.Logf("%[1]T %[1]v %v\n", v2, ok2)

	v3, ok3 := t2.(*T3)
	t.Logf("%[1]T %[1]v %v\n", v3, ok3)
}

type EmptyInterface interface {
}

func TestEmptyInterface(t *testing.T) {
	var _ Programmer = new(GoProgrammer)
	var _ EmptyInterface = new(GoProgrammer)
	var p EmptyInterface = new(GoProgrammer)

	v, ok := p.(GoProgrammer)
	t.Logf("%[1]T %[1]v %v\n", v, ok)
}

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