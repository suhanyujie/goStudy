package client

import (
	"net/rpc"
	"log"
	"lesson/goRpc/one/service"
	"fmt"
)

// 入口函数
func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing error:", err)
	}
	args := &service.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err!= nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith:%d * %d = %d", args.A, args.B, reply)
}
