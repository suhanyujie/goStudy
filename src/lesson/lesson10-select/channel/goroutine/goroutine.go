package goroutine

import (
	"fmt"
)

// TestGoroutine3
func TestGoroutine3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch2 chan int) {
	for i := 0; ; i++ {
		ch2 <- i + 6
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Receive on ch1:%d\n", v)
		case v := <-ch2:
			fmt.Printf("Receive on ch2:%d\n", v)
		}
	}
}

// TestGoroutine
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

//TestGoroutine2
func TestGoroutine2() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "suhanyu"
	}()

	go func() {
		fmt.Println(<-ch1)
	}()

}
