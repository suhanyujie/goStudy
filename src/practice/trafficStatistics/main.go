package main

import (
	"flag"
	"fmt"
)

// 入口函数
func main() {
	filePath := flag.String("filePath", "/usr/local/Cellar/nginx/1.12.2_1/logs", "param sample:--paramName=123123")
	limit := flag.String("limit", "100", "param sample:--paramName=123123")
	flag.Parse()
	fmt.Println(*filePath)
	fmt.Println(*limit)
}

func WriteLog() {
	sample := ``;
	fmt.Println(sample)
}
