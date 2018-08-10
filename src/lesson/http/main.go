package main

import (
	"net/http"
	"log"
	"fmt"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	log.Print("req [%v]", *r)
	fmt.Fprintln(w, "hello world")
}

// 入口函数
func main()  {
	http.Handle("/", sayHello)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error",err)
	}
}
