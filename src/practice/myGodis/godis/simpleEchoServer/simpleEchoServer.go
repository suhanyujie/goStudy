package main

import (
	"net"
	"log"
	"fmt"
)

/**
net包实现服务端的核心部分是:

net.Listen()在给定的本地网络地址上来创建新的监听器。如果只传端口号给它，例如":61000", 那么监听器会监听所有可用的网络接口。 这相当方便，因为计算机通常至少提供两个活动接口，回环接口和最少一个真实网卡。 这个函数成功的话返回Listener。

Listener接口有一个Accept()方法用来等待请求进来。然后它接受请求，并给调用者返回新的连接。Accept()一般来说都是在循环中调用，能够同时服务多个连接。每个连接可以由一个单独的goroutine处理，正如下面代码所示的。
*/
const(
	Port=":3130"
)
// 入口函数
func main() {
	ln, err := net.Listen("tcp", Port)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("echo server start working...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go echoFunc(conn)
	}
}

func echoFunc(conn net.Conn) {
	buf := make([]byte, 1024)
	var body = `HTTP/1.1 200 OK
Content-type:application/html

who are u.....
`
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(n)
		fmt.Println(conn.LocalAddr())
		conn.Write([]byte(body))
		conn.Close()
	}
}
