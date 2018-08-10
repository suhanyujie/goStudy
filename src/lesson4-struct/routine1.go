package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("hello")
		c <- 1
	}()
	fmt.Println("hello2")
	<-c
	fmt.Println("hello3")
}
