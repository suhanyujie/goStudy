package stackType

import "fmt"

type BaseStack struct {
	CurrentLen int
	stack [4]interface{}
}

func (_this *BaseStack) Push(value interface{}) {
	if _this.CurrentLen < 0 {
		_this.CurrentLen = 0
	}
	_this.stack[_this.CurrentLen] = value
	_this.CurrentLen++;
}

// todo
func (_this *BaseStack) String()  {
	fmt.Printf("CurrentLen:%v,\n", _this.CurrentLen)
	fmt.Printf("stack:%v,\n", _this.stack)
}

// 从数组末尾去除一个数据
func (_this *BaseStack) Pop() interface{}  {
	len := _this.CurrentLen
	if len<=0 {
		return nil
	}
	popValue := _this.stack[len-1];
	_this.stack[len-1] = nil;
	_this.CurrentLen--;
	return popValue
}

