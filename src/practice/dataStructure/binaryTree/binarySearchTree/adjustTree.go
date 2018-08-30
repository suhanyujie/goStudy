package binarySearchTree

func LLRotation(node *TreeNode) *TreeNode {
	prchild := node.Right
	node.Right = prchild.Left
	prchild.Left = node
	return prchild;
}
