package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generateor() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1800)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func worker(id int, ch chan int) {
	for n := range ch {
		fmt.Printf("worker %d received:%d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// 入口函数
func main() {
	var c1, c2 = generateor(), generateor()
	w := createWorker(0)
	n := 0
	var values []int
	var activeValue int
	tm := time.After(4 * time.Second)
	tick := time.Tick(time.Second)
	for i := 0; i < 10000; i++ {
		var activeWorker chan<- int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
			fmt.Println("c1")
		case n = <-c2:
			values = append(values, n)
			fmt.Println("c2")
		case activeWorker <- activeValue:
			values = values[1:];
		case <-tick:
			fmt.Println("queue len is:", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
