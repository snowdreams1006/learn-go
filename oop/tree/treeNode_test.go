package tree

import (
	"fmt"
	"testing"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

func TestCreatTreeNode(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	t.Log(nodes)
}

func createTreeNode(val int) *treeNode {
	return &treeNode{value: val}
}

func TestConstructFn(t *testing.T) {
	root := createTreeNode(1)

	t.Log(root)
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func TestTreeNodeTraverse(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	root.print()
}

func print(node treeNode) {
	fmt.Println(node.value)
}

func TestTreeNodeTraverseByNormalWay(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	print(root)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func TestTreeNodeSetValue(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	root.left.setValue(4)
	root.left.print()

	root.setValue(100)
	root.print()

	pRoot := &root
	pRoot.setValue(200)
	pRoot.print()
}

func (node *treeNode) setValueWithNil(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node")
		return
	}
	node.value = value
}

func TestTreeNodeSetValueWithNil(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	var pRoot *treeNode
	pRoot.setValueWithNil(200)

	pRoot = &root
	pRoot.setValueWithNil(300)
	pRoot.print()
}

func (node *treeNode) traverse() {
	if node == nil{
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func TestTreeNodetraverse(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createTreeNode(2)
	root.right.left.setValue(4)

	root.traverse()
}