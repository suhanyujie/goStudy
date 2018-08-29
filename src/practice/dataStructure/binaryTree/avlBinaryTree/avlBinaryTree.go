package avlBinaryTree
/**
平衡二叉树的实现
	介绍+实现  https://www.cnblogs.com/skywang12345/p/3576969.html

 */


type TreeNode struct {
	Type   int//关键字 键值
	Data   interface{}
	Left   *TreeNode
	Right  *TreeNode
	Height int //子树的深度
}

type AvlTree struct {
	Root   *TreeNode
	Length int
}

func LRotation(node *TreeNode) *TreeNode {
	prchild := node.Right
	node.Right = prchild.Left
	prchild.Left = node

	return prchild
}

func Max(a, b int) int {
	if a>b {
		return a
	} else {
		return b
	}
}

