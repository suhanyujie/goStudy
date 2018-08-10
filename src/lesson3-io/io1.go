package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	//"path/filepath"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, 10)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func sampleReadFromString() {
	p, _ := ReadFrom(strings.NewReader("hello world~~~~"), 3)
	fmt.Println(p)
}

func sampleReadFromFile() {
	fmt.Println(os.Getwd())
	file, err := os.Open("src/lesson2/client.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	p, _ := ReadFrom(file, 9)
	fmt.Println(string(p))
}

func main() {
	sampleReadFromFile()
}
