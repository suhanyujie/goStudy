package main

import "fmt"

var A string;

// 入口函数
func main() {
	fmt.Printf("%p\n", &A)
	//A = "suhanyu"
	A := "suhanyu"
	fmt.Printf("%p\n", &A)
}
