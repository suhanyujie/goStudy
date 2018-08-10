package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

// 入口函数
func main() {
	ch := make(chan string)
	fmt.Printf("who are u...\n")
	go sendData(ch)
	if v, ok := <-ch; ok {
		fmt.Printf("val:%s", v)
	}
	defer close(ch)
}
