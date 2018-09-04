package main

import (
	string2 "practice/http/libs/string"
	"fmt"
)

// 入口函数
func main() {
	chapterName := "第一千二百零六"
	num := string2.Chinese2Int(chapterName)
	fmt.Println(num)

	/*
	listUrl := "https://www.biduo.cc/biquge/17_17308"
	//url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
	//status,err := testSpider.GetDetail(url)
	testSpider.GetBaseUrl(listUrl)
	status,data,err := testSpider.GetList(listUrl)
	if status!=nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data)
	logs.Info(err)*/
}
