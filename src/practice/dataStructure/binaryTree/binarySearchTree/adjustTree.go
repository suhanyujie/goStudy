package binarySearchTree

func LRotation(node *TreeNode) *TreeNode {
	prchild := node.Right
	node.Right = prchild.Left
	prchild.Left = node


}
