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
	fmt.Printf("%T %v\n", p, p.WriteHelloWord())
}

func TestPolymorphism(t *testing.T) {
	gp := new(GoProgrammer)
	jp := new(JavaProgrammer)

	writeFirstProgram(gp)
	writeFirstProgram(jp)
}
