package main

import (
	"practice/dataStructure/binaryTree/binarySearchTree"
	"fmt"
)

// 入口函数
func main() {
     bst := binarySearchTree.GetRoot()
     bst.InsertNode(123)
     fmt.Println(bst)
}
