# 字典映射

- go 语言仅支持封装,不支持继承和多态
- go 语言没有class,只有struct


结构的定义

```go
type TreeNode struct
Left,Right *TreeNode
Value int
```

不管指针还是实例,都是 `.`

```go
func TestCreatTreeNode(t *testing.T) {
	var root treeNode

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
}
```

- 不论地址还是结构本身,一律使用 `.` 来访问成员

- 使用自定义工厂函数
- 注意返回了了局部变量的地址

> 结构创建在堆上还是栈上?java都是堆上(new垃圾回收),c++局部变量是栈立即销毁,返回时堆分配,手动释放.

不需要知道go,垃圾回收

```go
func createTreeNode(val int) *treeNode {
	return &treeNode{value: val}
}

func TestConstructFn(t *testing.T) {
	root := createTreeNode(1)
	
	t.Log(root)
}
```

接收者

```go
func (node treeNode) print() {
	fmt.Println(node.value)
}
```

遍历

```go
func TestTreeNodeTraverse(t *testing.T) {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	root.print()
}
```

```go
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
```

```go
func (node treeNode) setValue(value int) {
	node.value = value
}

func TestTreeNodeSetValue(t *testing.T){
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	root.left.setValue(4)
	root.left.print()
}
```

```go
func (node *treeNode) setValue(value int) {
	node.value = value
}

func TestTreeNodeSetValue(t *testing.T){
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	root.left.setValue(4)
	root.left.print()

	root.setValue(100)
	root.print()
}
```

- 显示定义和命名方法接收者
- nil 指针也可以调用方法

```go

func (node *treeNode) setValueWithNil(value int) {
	if node == nil{
		fmt.Println("Setting value to nil node")
		return
	}
	node.value = value
}

func TestTreeNodeSetValueWithNil(t *testing.T){
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
```


```go
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
```

值接收者和指针接收者

- 要改变内容必须使用指针接收者
- 结构过大也要考虑使用指针接收者
- 一致性: 如果有指针接收者,最好都是指针接收者

- 值接收者是go语言特有
- 值/指针接收者均可接收值/

封装

- 名字一般使用CamelCase
- 首字母大写: public
- 首字母小写: private

> 针对package来说

- 每个目录一个包
- main包包含可执行入口
- 为结构定义的方法必须放在同一个包内
- 可以是不同的文件

```go
package main

import "github.com/snowdreams1006/learn-go/oop/tree"

func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
}
```

扩展已有类型

如何扩充系统类型或者别人的类型

- 定义别名
- 使用组合

```
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
```

```go
func TestQueue_Push(t *testing.T) {
	q := Queue{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.IsEmpty())
	t.Log(q.Pop())
	t.Log(q.IsEmpty())
}

```

- gopath 环境变量的设置

- 默认~/go(unix/linux),%userprofile\go(windows)%
- 官方推荐: 所有项目和第三方库都放在同一个gopath
- 也可以将不同的项目放到不同的gopath

go get golang.org/x/tools/cmd/goimports

使用 gopm 获取无法下载的包

go get github.com/gpmgo/gopm

gopm get golang.org/x/tools/cmd/goimports

go install ./...

