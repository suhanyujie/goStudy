package main

import (
	"lesson/structStudy/stack/stackType"
)

// 入口函数
func main() {
	stack1 := new(stackType.BaseStack)
	stack1.Push(123)
	stack1.Push(43)
	stack1.Push(56)
	stack1.String()
	stack1.Pop()
	stack1.String()
}
