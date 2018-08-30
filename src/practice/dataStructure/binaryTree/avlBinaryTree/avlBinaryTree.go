package avlBinaryTree

/**
平衡二叉树的实现
	介绍+实现  https://www.cnblogs.com/skywang12345/p/3576969.html
	一共有4种旋转操作

 */

type TreeNode struct {
	Type   int //关键字 键值
	Data   interface{}
	Left   *TreeNode
	Right  *TreeNode
	Height int //子树的深度
}

type AvlTree struct {
	Root   *TreeNode
	Length int
}

func LLRotation(node *TreeNode) *TreeNode {
	k1 := node.Left
	node.Left = k1.Right
	k1.Right = node;
	node.Height = Max(node.Left.Height,node.Right.Height) + 1
	k1.Height = Max(k1.Left.Height,node.Height) + 1

	return k1
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
