package _struct

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"unsafe"
)

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

func (e Employee) toString() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToString(t *testing.T) {
	e := Employee{"0", "Bob", 20}

	t.Log(e.toString())
}

func (e *Employee) toStringPointer() string {
	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToStringPointer(t *testing.T) {
	e := &Employee{"0", "Bob", 20}

	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	t.Log(e.toStringPointer())
}

func (e Employee) toStringValue() string {
	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestToStringValue(t *testing.T) {
	e := Employee{"0", "Bob", 20}

	fmt.Printf("Name address is %x\n", unsafe.Pointer(&e.Name))

	t.Log(e.toStringValue())
}

type point struct {
	i, j int
}

var dir []point = []point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func TestPoint(t *testing.T) {
	if file, err := os.Open("container/maza/maze.in"); err != nil {
		panic(err)
	} else {
		t.Log(file)
	}
}

type MyDynamicArray struct {
	ptr *[10]int
	len int
	cap int
}

func NewMyDynamicArray() *MyDynamicArray {
	var myDynamicArray MyDynamicArray

	myDynamicArray.len = 0
	myDynamicArray.cap = 10
	var arr [10]int
	myDynamicArray.ptr = &arr

	return &myDynamicArray
}

func (myArr *MyDynamicArray) AddLast(value int) {
	myArr.Add(myArr.len, value)
}

func (myArr *MyDynamicArray) AddFirst(value int) {
	myArr.Add(0, value)
}

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

func (myArr *MyDynamicArray) GetLen() int {
	return myArr.len
}

func (myArr *MyDynamicArray) GetCap() int {
	return myArr.cap
}

func (myArr *MyDynamicArray) IsEmpty() bool {
	return myArr.len == 0
}

func (myArr *MyDynamicArray) Print() string {
	res := fmt.Sprintf("Array: len = %d , cap = %d\n", myArr.len, myArr.cap)
	res = res + "["

	for i := 0; i < myArr.len; i++ {
		res = res + strconv.Itoa((*myArr.ptr)[i])
		if i != myArr.len-1 {
			res = res + ", "
		}
	}

	res = res + "]"
	return res
}

func TestMyDynamicArray(t *testing.T) {
	myDynamicArray := NewMyDynamicArray()

	for i := 0; i < myDynamicArray.cap; i++ {
		myDynamicArray.AddLast(i)
	}

	t.Logf("array = %v,len = %d,cap = %d", *myDynamicArray.ptr, myDynamicArray.len, myDynamicArray.cap)

	fmt.Println(myDynamicArray.Print())

	myDynamicArray.Set(6,666)

	fmt.Println(myDynamicArray.Print())

	fmt.Println(myDynamicArray.Get(6))
	fmt.Println(myDynamicArray.IsEmpty())
}
