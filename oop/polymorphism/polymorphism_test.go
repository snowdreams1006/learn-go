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

	t.Logf("%[1]T %[1]v\n",t2)
	t.Logf("%[1]T %[1]v\n",t3)
	t.Logf("%[1]T %[1]v %v\n",t2,t2.Equal(t3))
}

func TestInterfaceTypeDeduce(t *testing.T) {
	var t2 Equaler = new(T2)
	var t3 Equaler = new(T3)

	t.Logf("%[1]T %[1]v %[2]T %[2]v\n",t2,t2.(*T2))
	t.Logf("%[1]T %[1]v %[2]T %[2]v\n",t3,t3.(*T3))
}