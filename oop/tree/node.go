package tree

import (
	"fmt"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func createTreeNode(val int) *TreeNode {
	return &TreeNode{value: val}
}

func (node TreeNode) Print() {
	fmt.Println(node.value)
}

func Print(node TreeNode) {
	fmt.Println(node.value)
}

func (node *TreeNode) SetValue(value int) {
	node.value = value
}

func (node *TreeNode) SetValueWithNil(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node")
		return
	}
	node.value = value
}

func (node *TreeNode) Traverse() {
	if node == nil{
		return
	}

	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}