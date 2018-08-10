package main

import (
	"os"
	"bufio"
	"errors"
	"fmt"
)


func writeFile(filename string) {
	file,err := os.Create(filename)
	if err!=nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

}

// 入口函数
func main() {
	fmt.Println("who are u...");
	errors.New("this is customer error...")
}
