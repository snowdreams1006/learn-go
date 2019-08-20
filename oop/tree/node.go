package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func Print(node Node) {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) SetValueWithNil(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node")
		return
	}
	node.Value = value
}
