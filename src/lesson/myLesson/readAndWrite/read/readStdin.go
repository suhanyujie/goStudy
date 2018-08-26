package read

import (
	"bufio"
	"os"
	"fmt"
	"log"
)

func TestReadFromCmd() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入字母：")
	input,err := inputReader.ReadString('\n')
	if err!=nil {
		log.Fatal(err)
	}

	fmt.Printf("your string is: %s\n", input)
}
