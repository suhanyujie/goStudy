package binarySearchTree

import "fmt"

/**
@desc 参考自己的c代码逻辑  https://github.com/suhanyujie/DataStructure/blob/master/BSTree/BSTree.c
四种基本的遍历思想为：

前序遍历：根结点 ---> 左子树 ---> 右子树

中序遍历：左子树---> 根结点 ---> 右子树

后序遍历：左子树 ---> 右子树 ---> 根结点

层次遍历：仅仅需按层次遍历就可以（https://www.cnblogs.com/llguanli/p/7363657.html）
 */

type TreeNode struct {
	Data   interface{}
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type SearchTree struct {
	Root   *TreeNode
	Length int
}

// todo 插入一个节点
func (_this *SearchTree) InsertNode(value int) {
	newNode := &TreeNode{
		value,
		nil,
		nil,
		nil,
	}
	var currentNode *TreeNode = _this.Root
	var lastNode *TreeNode
	var dataValue int;
	//fmt.Printf("外部currentNode:%p\n", currentNode)
	for currentNode != nil {
		dataValue = currentNode.Data.(int)
		lastNode = currentNode
		//fmt.Printf("内部currentNode:%p\n", currentNode)
		if value < dataValue {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	newNode.Parent = lastNode
	if value < dataValue {
		lastNode.Left = newNode
	} else {
		lastNode.Right = newNode
	}
}

// todo 前序遍历
func (_this *SearchTree) BeforeTraverse(beginNode *TreeNode) {
	if beginNode == nil {
		return;
	}
	value := beginNode.Data.(int)
	fmt.Printf("%d\t", value);
	_this.BeforeTraverse(beginNode.Left)
	_this.BeforeTraverse(beginNode.Right)
}

func GetInitBst() *SearchTree {
	Root := GetRoot()
	Tree := &SearchTree{
		Root,
		1,
	}

	return Tree
}

// todo 获取一个根节点
func GetRoot() *TreeNode {
	Bst := &TreeNode{
		0,
		nil,
		nil,
		nil,
	}

	return Bst
}
