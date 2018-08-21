package main

import (
	"practice/http/testSpider"
	"log"
	"github.com/astaxie/beego/logs"
)

// 入口函数
func main() {
	url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
	status,err := testSpider.GetDetail(url)
	if status!=nil {
		log.Fatal(err)
	}
	logs.Info(err)
}
