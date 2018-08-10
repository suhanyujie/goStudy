package main

import "fmt"

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	var r1, r2 int = div(10, 3)
	fmt.Println(r1, r2)
}
