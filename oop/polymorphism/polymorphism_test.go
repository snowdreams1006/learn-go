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
	var p Programmer = new(GoProgrammer)
	//var _ MyProgrammer = new(JavaProgrammer)
}


