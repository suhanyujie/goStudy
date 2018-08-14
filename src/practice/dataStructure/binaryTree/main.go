package main

import (
	"practice/dataStructure/binaryTree/binarySearchTree"
)

// 入口函数
func main() {
	bst := binarySearchTree.GetInitBst()
	bst.InsertNode(12)
	//for i := 12; i < 15; i++ {
	//	bst.InsertNode(i)
	//}
	bst.BeforeTraverse(bst.Root)
}
