package binarySearchTree

/**
@desc 参考自己的c代码逻辑  https://github.com/suhanyujie/DataStructure/blob/master/BSTree/BSTree.c
 */

type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
	Parent *TreeNode
}

type SearchTree struct {
	Root   *TreeNode
	Length int
}

// todo 插入一个节点  【未完待续...】
func (_this *SearchTree) InsertNode(value int) {
	newNode := &TreeNode{
		value,
		nil,
		_this.Root,
		nil,
	}
	currentNode := _this.Root
	var lastNode *TreeNode
	var dataValue int;
	for currentNode != nil {
		dataValue = currentNode.Data.(int)
		lastNode = currentNode
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
