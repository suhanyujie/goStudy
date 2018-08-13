package binarySearchTree

type TreeNode struct {
	Data interface{}
	Left *TreeNode
	Right *TreeNode
}

type SearchTree struct {
	Root *TreeNode
	Length int
}

// todo 插入一个节点  【未完待续...】
func (_this *SearchTree) InsertNode(value int) {
	newNode := &TreeNode{
		value,
		nil,
		_this.Root,
	}
	compareNode := _this.Root
	for true {
		dataValue := compareNode.Data.(int)
		if value < dataValue {
			if compareNode.Left == nil {
				compareNode.Left = newNode
				break;
			}
			compareNode = compareNode.Left
		}else{
			if compareNode.Right == nil {
				compareNode.Right = newNode
				break;
			}
			compareNode = compareNode.Right
		}
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
func GetRoot () *TreeNode {
	Bst := &TreeNode{
		0,
		nil,
		nil,
	}

	return Bst
}


