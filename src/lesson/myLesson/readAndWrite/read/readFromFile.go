package read

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"io"
)

var filePath = "src/lesson/myLesson/readAndWrite/read/readStdin.go"

func TestReadFromFile() {
	inputFile,err := os.Open(filePath)
	if err!=nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString,err := inputReader.ReadString('\n')
		fmt.Println(inputString)
		if err == io.EOF{
			fmt.Println("读取文件结束-----")
			return;
		}
	}
}



