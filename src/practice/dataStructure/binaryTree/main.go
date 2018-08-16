package main

import (
	"practice/dataStructure/binaryTree/binarySearchTree"
	"fmt"
)

// 入口函数
func main() {
	bst := binarySearchTree.GetInitBst()
	for i := 12; i < 321; i++ {
		bst.InsertNode(i)
	}
	//bst.BeforeTraverse(bst.Root)
	bst.AfterTraverse(bst.Root)
	fmt.Println("------------------")
	bst.MiddleTraverse(bst.Root)
}

