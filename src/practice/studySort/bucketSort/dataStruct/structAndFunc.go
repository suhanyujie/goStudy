package dataStruct

import "fmt"

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
func (_this *LinkNode) Push(node *LinkNode) *LinkNode {
	var next *LinkNode = _this
	//可以确定的是 当前的节点一定不为空
	if next.Next == nil {
		if node.Data > next.Data {
			next.Next = node
			node.Prev = next
			return next
		} else {
			node.Next = next
			next.Prev = node
			return node
		}
	}
	var rootNode *LinkNode = next;
	//新加的节点插入到2个节点之间
	for ; next.Next != nil; next = next.Next {
		if node.Data > next.Data && node.Data <= next.Next.Data {
			node.Next = next.Next
			next.Next.Prev = node
			next.Next = node
			node.Prev = next
			return rootNode
		}
	}
	//此时还是没找到，意味着，node对应的数是最大的，直接放在最后
	next.Next = node
	node.Prev = next

	return rootNode
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

func PrintLink(node *LinkNode) {
	//循环遍历节点
	for ;node!=nil ;node = node.Next  {
		fmt.Printf("%d\t", node.Data)
	}
}
