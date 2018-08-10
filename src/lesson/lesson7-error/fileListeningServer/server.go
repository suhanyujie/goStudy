package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
)



// 入口函数
func main() {
	fmt.Println("server running in port 7000...")
	http.HandleFunc("/list/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/list/"):]
		fmt.Println(path)
		file, err := os.Open("go语言中的range.md")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		w.Write(all)
	})
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		panic(err)
	}
}
