package main

import (
	"fmt"
	"time"
	"sync"
)

type worker struct {
	in chan int
	done func()
}

func createWorker(c chan int) {
	for i := 122; i > 110; i-- {
		c <- i
	}
	close(c)
}

func workerFunc(id int, c chan int) {
	for n := range c {
		fmt.Printf("channel %d received %d\n", id, n)
	}
}

func chanDemo() {
	c := make(chan int);
	go createWorker(c)
	workerFunc(0, c)
	time.Sleep(time.Second)
}

func doWork(id int, w worker) {
	for n:=range w.in{
		fmt.Printf("Worker %d reveive %c\n", id, n)
		w.done()
	}
}

func createWorker2(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}


func chanDemo2() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
}

// 入口函数
func main() {
	chanDemo2()
}
