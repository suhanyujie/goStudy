package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

// 入口函数
func main() {
     conn,err:=net.Dial("tcp","127.0.0.1:3130")
	if err!=nil {
		log.Println(err)
	}
     fmt.Fprintf(conn,"GET / HTTP/1.0\r\n")
     status,err:=bufio.NewReader(conn).ReadString('\n')
     fmt.Println(status)
}
