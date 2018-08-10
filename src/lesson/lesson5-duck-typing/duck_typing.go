package main

import (
	"fmt"
	"lesson/lesson5-duck-typing/mock"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://laravel.suhanyu.top")
}


func main() {

	var r Retriever
	r = mock.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}

	fmt.Printf("%T--%v", r,r)
	//fmt.Println(download(r))
	fmt.Println("123")
}
