package main

import (
	"fmt"
	"github.com/snowdreams1006/learn-go/oop/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	//myTreeNode.node.Left.postOrder()
	myLeft := myTreeNode{myNode.node.Left}
	myLeft.postOrder()

	myRight := myTreeNode{myNode.node.Right}
	myRight.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	fmt.Println()
}
