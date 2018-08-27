package goroutine

import (
	"fmt"
)

func TestGoroutine() {
	ch1 := make(chan string)
	go sendData(ch1)
	go getData(ch1)
}

func sendData(ch chan string) {
	ch <- "shenshuo"
	ch <- "suhanyu"
	ch <- "surui"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}

func TestGoroutine2() {
	ch1 := make(chan string)
	go func() {
		ch1<-"suhanyu"
	}()

	go func() {
		fmt.Println(<-ch1)
	}()

}
