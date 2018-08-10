package server

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
	"lesson/goRpc/one/service"
)

// 入口函数
func main() {
	arich := new(service.Arich)
	rpc.Register(arich)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)
}
