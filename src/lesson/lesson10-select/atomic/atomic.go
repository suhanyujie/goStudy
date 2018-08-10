package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	a.value++
	defer a.lock.Unlock()
}
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

// 入口函数
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Second)
	fmt.Println(a.get())
}
