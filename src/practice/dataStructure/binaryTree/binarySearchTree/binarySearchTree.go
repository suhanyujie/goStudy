package binarySearchTree

import "fmt"

/**
@desc 参考自己的c代码逻辑  https://github.com/suhanyujie/DataStructure/blob/master/BSTree/BSTree.c
四种基本的遍历思想为：
	前序遍历：根结点 ---> 左子树 ---> 右子树
	中序遍历：左子树---> 根结点 ---> 右子树
	后序遍历：左子树 ---> 右子树 ---> 根结点
	层次遍历：仅仅需按层次遍历就可以（https://www.cnblogs.com/llguanli/p/7363657.html）

树的高度
	高度的定义为：从结点x向下到某个叶结点最长简单路径中边的条数
	在树中，树的高度也是从下往上数。最底层高度为1，
树的深度
	从根节点往下，列如上图中：根节点深度为1
最深的叶结点的深度就是树的深度，树根的高度就是树的高度。这样树的高度和深度是相等的
1.https://blog.csdn.net/Fanpei_moukoy/article/details/23828603
2.https://blog.csdn.net/bwh12398/article/details/78011819

平衡二叉树的实现
	https://blog.csdn.net/Jinhua_Wei/article/details/79595507
	介绍+实现  https://www.cnblogs.com/skywang12345/p/3576969.html

 */

type TreeNode struct {
	Data   interface{}
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
	Height int//树的深度
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
		1,
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

// todo 中序遍历
func (_this *SearchTree) MiddleTraverse(beginNode *TreeNode)  {
	if beginNode == nil {
		return;
	}
	_this.MiddleTraverse(beginNode.Left)
	value := beginNode.Data.(int)
	fmt.Printf("%d\t", value);
	_this.MiddleTraverse(beginNode.Right)
}

// todo 后序遍历
func (_this *SearchTree) AfterTraverse(beginNode *TreeNode)  {
	if beginNode == nil {
		return;
	}
	_this.AfterTraverse(beginNode.Right)
	_this.AfterTraverse(beginNode.Left)
	value := beginNode.Data.(int)
	fmt.Printf("%d\t", value);
}

func GetInitBst() *SearchTree {
	Root := GetRoot()
	Tree := &SearchTree{
		Root,
		1,
	}

	return Tree
}

//todo 在树中搜索一个确定的值的节点
func SearchNode(Tree *SearchTree, searchValue int) *TreeNode {
	currentNode := Tree.Root
	for currentNode !=nil {
		currentValue := currentNode.Data.(int)
		if searchValue==currentValue {
			break;
		}
		if searchValue>currentValue {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}

	return currentNode
}

// todo 获取一个根节点
func GetRoot() *TreeNode {
	Bst := &TreeNode{
		0,
		nil,
		nil,
		nil,
		1,
	}

	return Bst
}

func GetHeight(_this *TreeNode) int {
	if _this == nil {
		return 0
	}
	return _this.Height;
}
