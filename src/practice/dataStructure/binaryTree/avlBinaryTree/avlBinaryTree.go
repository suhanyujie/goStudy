package avlBinaryTree

import "fmt"

/**
平衡二叉树的实现
	介绍+实现  https://www.cnblogs.com/skywang12345/p/3576969.html
	一共有4种旋转操作
	更加详细的8种旋转操作

实现步骤：
	1.方法1：创建树				100%
	2.方法2：获取树的高度			100%
	3.方法3：比较两个值的大小		100%
	4.方法4.1 左左旋
	5.方法4.2 左右旋
	6.方法4.3 右左旋
	7.方法4.4 右右旋

注意事项：
	每新增一个节点a，则a的height是0，因为插入操作后，a会成为叶子节点
	关于高度，有的地方规定"空二叉树的高度是-1"，而本文采用维基百科上的定义：树的高度为最大层次。即空的二叉树的高度是0，非空树的高度等于它的最大层次(根的层次为1，根的子节点为第2层，依次类推)


数据结构和算法目录：https://www.cnblogs.com/skywang12345/p/3603935.html
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
func CreateAvlTree(data interface{}, left *TreeNode, right *TreeNode) *TreeNode {
	newNode := new(TreeNode)
	newNode.Left = left
	newNode.Right = right
	newNode.Data = data
	newNode.Height = 1

	return newNode
}

/**
1.1新增节点
 */
func InsertNode(val int, tree *TreeNode) *TreeNode {
	if tree == nil {
		tree = &TreeNode{
			1,
			val,
			nil,
			nil,
			1,
		}
		return tree
	}
	var originRoot *TreeNode = tree;
	//通过递归 最终找到待插入的叶子节点
	if val > tree.Data.(int) {
		tree.Right = InsertNode(val, tree.Right)
		//检查是否失衡
		if Height(tree.Right)-Height(tree.Left) >= 2 {
			if val < tree.Right.Data.(int) {
				originRoot = RRRotation(originRoot)
				fmt.Println("发生了右右旋转")
			} else {
				originRoot = RLRotation(originRoot)
				fmt.Println("发生了右左旋转")
			}
		}
	} else if (val < tree.Data.(int)) {
		tree.Left = InsertNode(val, tree.Left)
		if Height(tree.Left)-Height(tree.Right) >= 2 {
			if val > tree.Left.Data.(int) {
				originRoot = LRRotation(originRoot)
				fmt.Println("发生了左右旋转")
			} else {
				originRoot = LLRotation(originRoot)
				fmt.Println("发生了左左旋转")
			}
		}
	} else {
		fmt.Println("失败，不允许添加相同值的节点！")
	}
	originRoot.Height = Max(Height(originRoot.Left), Height(originRoot.Right)) + 1

	return originRoot
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
func MaxOfI(data1, data2 interface{}) interface{} {
	num1 := data1.(int)
	num2 := data2.(int)
	if num1 > num2 {
		return data1
	} else {
		return data2
	}
}

/**
4.1 左左旋 已左测节点为轴心旋转
	左左旋
 */
func LLRotation(node *TreeNode) *TreeNode {
	var nodeLeftHeight int = 0
	var nodeRightHeight int = 0
	var k1 *TreeNode = node.Left
	node.Left = k1.Right
	k1.Right = node
	if k1.Right != nil {
		nodeRightHeight = k1.Right.Height
	}
	if k1.Left != nil {
		nodeLeftHeight = k1.Left.Height
	}
	k1.Height = 1 + Max(nodeLeftHeight, nodeRightHeight)
	nodeLeftHeight = 0;
	nodeRightHeight = 0
	if node.Right != nil {
		nodeRightHeight = node.Right.Height
	}
	if node.Left != nil {
		nodeLeftHeight = node.Left.Height
	}
	node.Height = 1 + Max(nodeLeftHeight, nodeRightHeight)

	return k1
}

//4.2右右旋转
func RRRotation(node *TreeNode) *TreeNode {
	var nodeLeftHeight int = 0
	var nodeRightHeight int = 0
	k1 := node
	var k2 *TreeNode = k1.Right
	k1.Right = k2.Left
	k2.Left = k1
	if k1.Right != nil {
		nodeRightHeight = k1.Right.Height
	}
	if k1.Left != nil {
		nodeLeftHeight = k1.Left.Height
	}
	k1.Height = Max(nodeLeftHeight, nodeRightHeight) + 1
	nodeLeftHeight = 0;
	nodeRightHeight = 0
	if k2.Right != nil {
		nodeRightHeight = k2.Right.Height
	}
	if k2.Left != nil {
		nodeLeftHeight = k2.Left.Height
	}
	k2.Height = Max(nodeLeftHeight, nodeRightHeight) + 1

	return k2
}

/**
4.3左右旋转
 */
func LRRotation(node *TreeNode) *TreeNode {
	var k1 *TreeNode = RRRotation(node.Left)
	node.Left = k1
	var k2 *TreeNode = LLRotation(node)
	return k2
}

/**
4.4右左旋转
 */
func RLRotation(node *TreeNode) *TreeNode {
	var k1 *TreeNode = LLRotation(node.Right)
	node.Right = k1
	var k2 *TreeNode = RRRotation(node)
	return k2
}

/**
获取两个数中 值大的那个
 */
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
前序遍历

 */
func (_this *TreeNode) PrevTraverse() {
	fmt.Printf("%d\t", _this.Data.(int))
	if _this.Left != nil {
		ShowNode(_this.Left)
	}
	if _this.Right != nil {
		ShowNode(_this.Right)
	}
}

/**
中序遍历

 */
func (_this *TreeNode) MiddleTraverse() {
	if _this.Left != nil {
		ShowNode(_this.Left)
	}
	fmt.Printf("%d\t", _this.Data.(int))
	if _this.Right != nil {
		ShowNode(_this.Right)
	}
}

/**
后序遍历

 */
func (_this *TreeNode) AfterTraverse() {
	if _this.Left != nil {
		ShowNode(_this.Left)
	}
	if _this.Right != nil {
		ShowNode(_this.Right)
	}
	fmt.Printf("%d\t", _this.Data.(int))
}

// todo 后续遍历打印节点值
func (_this *TreeNode) ToString() {
	ShowNode(_this)
}

// todo 后序遍历
func ShowNode(node *TreeNode) {
	if node.Left != nil {
		ShowNode(node.Left)
	}
	if node.Right != nil {
		ShowNode(node.Right)
	}
	fmt.Printf("%d\t", node.Data.(int))
}
