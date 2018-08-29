package tick

import (
	"time"
	"fmt"
	"errors"
)

//TestTick2
func TestTick2() {
	ch := make(chan error, 1)
	//模拟耗时调用
	go func() { time.Sleep(3 * time.Second); ch <- errors.New("has return....") }()
	select {
	case resp := <-ch:
		fmt.Printf("resp1 call,result is:%s\n", resp)
	case <-time.After(time.Second * 5)://定时器信号
		fmt.Println("time After....")
		break;
	}
}

//TestTick1
func TestTick1() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick:....")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("---------------.")
			time.Sleep(5e7)
		}
	}
}
