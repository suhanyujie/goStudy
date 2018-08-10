package main

import "fmt"

// 普通函数中返回局部变量地址也是可以给外面的程序使用的
// 结构的创建是在堆上 分配在堆上才能被执行垃圾回收（java c++）
// 1个目录只能有1个包

type treeNode struct {
	value       int
	left, right *treeNode
}

// 扩展别人的类型
type myTreeNode struct {
	node *treeNode
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("set value to nil node.Ignored.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Print(node.value)
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{value: 2}
	root.right = &treeNode{5, nil, nil}
	root.right.right = new(treeNode)
	root.right.right.setValue(46)

	//fmt.Println(root.right)
	root.traverse()
}
