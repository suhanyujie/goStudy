package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 30; j++ {
				fmt.Printf("hello from "+"goroutine %d\n", i)
				// 主动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
