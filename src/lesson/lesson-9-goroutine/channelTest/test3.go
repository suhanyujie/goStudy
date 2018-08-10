package main

import (
	"log"
	"runtime"
	"time"
)

// 入口函数
func main() {
     logGoNum()
     var ch chan int
     go func(ch chan int) {
     	<-ch
     }(ch)

     for range time.Tick(1*time.Second){
     	logGoNum()
	 }
}

func logGoNum() {
	log.Printf("goroutine number:%d\n", runtime.NumGoroutine())
}
