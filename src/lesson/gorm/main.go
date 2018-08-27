package main

import (
	"lesson/gorm/beegoOrm"
	"runtime"
	"fmt"
)

// 入口函数
func main() {
	//start.Test1()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
	beegoOrm.Test2()
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
