package dataStruct

type LinkNode struct {
	Data int
	Prev *LinkNode
	Next *LinkNode
}

// todo 生成一个新的链表节点
func (_this *LinkNode) NewLink(value int) *LinkNode {
	return &LinkNode{value, nil, nil}
}

// todo 将一个节点加入链表的末尾
func (_this *LinkNode) Push(node *LinkNode) {
	var next *LinkNode = _this
	for next.Next != nil {
		next = next.Next
	}
	next.Next = node
	node.Prev = next
}

// todo 将最后一个节点取出
func (_this *LinkNode) Pop() *LinkNode {
	next := _this
	for next.Next != nil {
		next = next.Next
	}
	popedNode := next
	next = nil
	return popedNode
}

func GetNode(value int) *LinkNode {
	return &LinkNode{value, nil, nil}
}
