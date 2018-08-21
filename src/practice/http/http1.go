package main

import (
	"practice/http/testSpider"
	"log"
	"github.com/astaxie/beego/logs"
)

// 入口函数
func main() {
	listUrl := "https://www.biduo.cc/biquge/17_17308"
	//url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
	//status,err := testSpider.GetDetail(url)
	testSpider.GetBaseUrl(listUrl)
	status ,err := testSpider.GetList(listUrl)
	if status!=nil {
		log.Fatal(err)
	}
	logs.Info(err)
}
