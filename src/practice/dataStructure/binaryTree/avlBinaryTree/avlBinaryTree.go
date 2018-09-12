package avlBinaryTree

/**
平衡二叉树的实现
	介绍+实现  https://www.cnblogs.com/skywang12345/p/3576969.html
	一共有4种旋转操作

实现步骤：
	1.方法1：创建树				100%
	2.方法2：获取树的高度			100%
	3.方法3：比较两个值的大小		100%
	4.方法4.1 左左旋
	5.方法4.2 左右旋
	6.方法4.3 右左旋
	7.方法4.4 右右旋

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

//1.创建树
func CreateAvlTree(data interface{},left *TreeNode,right *TreeNode) *TreeNode {
	newNode := new(TreeNode)
	newNode.Left = left
	newNode.Right = right
	newNode.Data = data
	newNode.Height = 1

	return newNode
}

//2.获取树的高度
/**
树的高度为最大层次。即空的二叉树的高度是0，非空树的高度等于它的最大层次(根的层次为1，根的子节点为第2层，依次类推)。
 */
func Height(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	return tree.Height
}

//3.比较2个interface{}类型的数据大小，返回大的那个值
func MaxOfI(data1,data2 interface{}) interface{} {
	num1 := data1.(int)
	num2 := data2.(int)
	if num1>num2 {
		return data1
	} else {
		return data2
	}
}

//4.1 左旋
/**

 */
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
