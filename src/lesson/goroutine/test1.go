package main

import (
	"log"
	"time"
	"fmt"
)

func doSomething(id int) {
	log.Println("before do job")
	time.Sleep(2 * time.Second)
	log.Println("after job...")
}

func main() {
	//doSomething(1)
	//doSomething(2)
	//doSomething(3)
	ch := make(chan int)
	go func() {
		ch <- 120
	}()
	fmt.Println(<- ch)

	go doSomething(1)
	go doSomething(2)
	go doSomething(3)
	time.Sleep(4 * time.Second)
}
