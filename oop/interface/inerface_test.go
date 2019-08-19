package _interface

import (
	"fmt"
	"testing"
)

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
