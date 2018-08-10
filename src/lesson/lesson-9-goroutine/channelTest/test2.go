package main

import (
	"net/http"
	"log"
	"os/exec"
	"strings"
)

// 入口函数
// 入口函数
func main() {
     http.HandleFunc("/exec", func(w http.ResponseWriter,r *http.Request) {
     	defer func(){log.Printf("finished %v\n", r.URL)}()
     	out,err := genCmd(r).CombinedOutput()
		 if err!=nil {
			 w.WriteHeader(500)
			 w.Write([]byte(err.Error()))
			 return
		 }
		 w.Write(out)
	 })
     log.Fatal(http.ListenAndServe(":8000", nil))
}

func genCmd(r *http.Request) (cmd *exec.Cmd) {
	var args []string
	if got:=r.FormValue("args");got!="" {
		args = strings.Split(got, " ")
	}
	if c:=r.FormValue("cmd");len(args)==0 {
		cmd = exec.Command(c)
	}else{
		cmd = exec.Command(c, args...)
	}
	return
}
