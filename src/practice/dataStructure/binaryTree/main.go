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

/**
在加入节点5的时候 平衡遭到破坏 并且没有通过旋转得到解决
 */
func TestAvl() {
	var numArr []int = []int{1,2,3,4,5,6,7,8}
	var tree *avlBinaryTree.TreeNode
	tree = avlBinaryTree.CreateAvlTree(25, nil, nil)
	for _,data := range numArr{
		tree = avlBinaryTree.InsertNode(data,tree)
		tree.PrevTraverse()
	}

	//avlBinaryTree.InsertNode(36,tree)
	//avlBinaryTree.InsertNode(27,tree)

	//tree.ToString()
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
