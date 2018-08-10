package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Math int
type Quotient struct {
	Quo, Remem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
	}
	serverAddr := os.Args[1]
	// client, err := rpc.DialHTTP("tcp", serverAddr+":8080")
	client, err := rpc.Dial("tcp", serverAddr+":8080")
	if err != nil {
		log.Fatal("dialing", err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("call trouble--", err)
	}
	fmt.Printf("the result is: %d\n", reply)
}
