package main

import (
	"errors"
	"fmt"
	"net"
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

func (m *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *Math) Divide(args *Args, reply *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero..")
	}
	reply.Quo = args.A / args.B
	reply.Remem = args.A % args.B
	return nil
}

func main() {
	math := new(Math)
	rpc.Register(math)
	// rpc.HandleHTTP()
	// err := http.ListenAndServe(":8080", nil)
	// 2
	tcpAdress, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
	}
	listener, err := net.ListenTCP("tcp", tcpAdress)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(40001)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Fatal Error:", err.Error())
			continue
		}
		rpc.ServeConn(conn)
	}
}
