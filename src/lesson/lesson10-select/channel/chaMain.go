package main

import (
	"lesson/lesson10-select/channel/goroutine"
	"time"
)

/**
1.可以在运算循环中周期的使用 runtime.Gosched()：这会让出处理器，允许运行其他协程；它并不会使当前协程挂起，所以它会自动恢复执行。使用 Gosched() 可以使计算均匀分布，使通信不至于迟迟得不到响应。
2.GOMAXPROCS 变量。这会告诉运行时有多少个协程同时执行
3.默认情况下，通信是同步且无缓冲的




 */
// 入口函数
func main() {
	//channelFirst.Test1()
	goroutine.TestGoroutine2()
	time.Sleep(time.Second*2)
}
