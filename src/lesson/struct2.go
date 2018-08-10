package main

import (
	"fmt"
)

type Data struct {
}

// 测试方法
func (self Data) String() string {

	return "this  is string: data"
}

func main() {

	fmt.Printf("%v\n", Data{})
}
