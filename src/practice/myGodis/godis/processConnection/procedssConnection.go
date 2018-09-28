package main

/**
参考  https://segmentfault.com/a/1190000014733620


 */
import (
	"net"
	"log"
	"bufio"
	"sync"
	"io"
	"strings"
)

type HandleFunc func(writer *bufio.ReadWriter)

type EndPoint struct {
	listener net.Listener
	handler  map[string]HandleFunc
	m        sync.RWMutex
}

const (
	Port = ":3132"
)

// 入口函数
func main() {
	endpoint := NewEndPoint()
	endpoint.Listen()
}

func NewEndPoint() *EndPoint {
	return &EndPoint{
		handler: map[string]HandleFunc{},
	}
}

// todo 新增HandleFunc
func (_this *EndPoint) AddHandleFunc(name string, f HandleFunc) {
	_this.m.Lock()
	_this.handler[name] = f
	_this.m.Unlock()
}

// todo
func (_this *EndPoint) Listen() error {
	var err error
	_this.listener, err = net.Listen("tcp", Port)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Listen on", _this.listener.Addr().String())
	for {
		log.Println("Accept a connection request.")
		conn, err := _this.listener.Accept()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("Handle is dealing message...")
		go _this.handleMessage(conn)
	}
	return err
}

// todo
func (_this *EndPoint) handleMessage(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	for {
		log.Println("Receive command...")
		cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			log.Println("Reached EOF - close this connection.\n ---")
			return
		case err != nil:
			log.Println("\nError reading command.Got:'"+cmd+"'\n", err)
		}
		cmd = strings.Trim(cmd,"\n")
		log.Println(cmd+"'")
		_this.m.Lock()
		handleCommand,ok:=_this.handler[cmd]
		_this.m.Unlock()

		if !ok {
			log.Println("Command'"+cmd+"' is not regiested.")
			return
		}
		handleCommand(rw)
	}
}
