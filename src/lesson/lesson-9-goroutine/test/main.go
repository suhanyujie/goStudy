package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 入口函数
func main() {
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

func sayHello() {
	fmt.Println("this is no rain...")
}
