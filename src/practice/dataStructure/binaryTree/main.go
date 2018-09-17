package main

import (
	"practice/dataStructure/binaryTree/binarySearchTree"
	"math/rand"
	"fmt"
	"practice/dataStructure/binaryTree/avlBinaryTree"
)

// 入口函数
func main() {
	TestAvl()
}

func TestAvl() {
	tree := avlBinaryTree.CreateAvlTree(25, nil, nil)
	avlBinaryTree.InsertNode(12,tree)
	avlBinaryTree.InsertNode(16,tree)
	avlBinaryTree.InsertNode(36,tree)
	avlBinaryTree.InsertNode(27,tree)

	tree.ToString()
}

func TestBinaryTree() {
	//二叉搜索树的初始化
	bst := binarySearchTree.GetInitBst()
	for i := 0; i < 24; i++ {
		randVal := rand.Intn(300)
		bst.InsertNode(randVal)
	}
	fmt.Println("------------------")
	bst.MiddleTraverse(bst.Root)
	searchValue := 74
	//二叉搜索树的搜索
	findedNode := binarySearchTree.SearchNode(bst, searchValue)
	if findedNode == nil {
		fmt.Printf("\n没有找到值为%d的节点\n", searchValue)
	} else {
		fmt.Printf("\n已经找到值为%d的节点，其地址是：%p\n", searchValue, findedNode)
	}
}
