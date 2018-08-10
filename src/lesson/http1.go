package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func main() {
	// http.HandleFunc("/", myHandler{})

	if err := http.ListenAndServe(":8088", &myHandler{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("serve is start in port:8088..")
	}

}

func (this *myHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	rw.Write([]byte("index"))
}

func handlerFunc1(rw http.ResponseWriter, req *http.Request) {
	rw.Write(([]byte)("This is handler func.."))
}
