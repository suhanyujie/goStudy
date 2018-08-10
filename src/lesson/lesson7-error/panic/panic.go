package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred", err)
		} else {
			panic(err)
		}
	}()
	panic(errors.New("this is an error"))
}

// 入口函数
func main() {
	tryRecover()
}
