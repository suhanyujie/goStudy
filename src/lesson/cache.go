package main

import (
	"fmt"
)

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", (x + 0/x)) // panics if x == 0
}
