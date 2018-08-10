package main

import "fmt"

func calcTriangle(a, b int) int {
	return a + b
}

// 入口函数
func main() {
     sum := calcTriangle(1,3)
     fmt.Println(sum)
}
