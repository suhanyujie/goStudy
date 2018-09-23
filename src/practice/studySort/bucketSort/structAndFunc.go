package main

type LinkNode struct {
	Data int
	Prev *LinkNode
	Next *LinkNode
}

// todo 生成一个新的链表节点
func (_this *LinkNode) NewLink(value int)*LinkNode  {
	return &LinkNode{value,nil,nil}
}
