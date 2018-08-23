package future

import (
	"time"
	"fmt"
)

/**
Go-简洁的并发
http://www.yankay.com/go-clear-concurreny/

 */

type Query struct {
	sql chan string
	result chan string
}

//执行query
func ExecQuery(q Query) {
	go func() {
		sql :=  <-q.sql
		q.result <- "get:"+sql
	}()
}

func Test1() {
	q := Query{make(chan string,1), make(chan string,1)}
	ExecQuery(q)
	time.Sleep(time.Second * 4)
	q.sql <- "select * from oil_galaxy_order limit 10"
	fmt.Println(<-q.result)
	fmt.Println("func end----")
}
