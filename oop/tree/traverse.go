package tree

import "fmt"

func (node *Node) Traverse() {
	//if node == nil{
	//	return
	//}

	//node.Left.Traverse()
	//node.Print()
	//node.Right.Traverse()

	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
