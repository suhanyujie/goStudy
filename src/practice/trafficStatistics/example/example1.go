package main

import "fmt"

/**
* 进程：太重
* 线程：上下文切换开销太大
* 协程：轻量级线程，简洁的并发模式



 */

// 入口函数
func main() {
	//FIFO：先进先出 first in first out
	var chMessage = make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			chMessage <- i + 2
		}
		close(chMessage)
	}()
	for num1 := range chMessage {
		fmt.Printf("data is:%d\n", num1)
	}
}

