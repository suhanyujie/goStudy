package main

import (
	string2 "practice/http/libs/string"
)

// 入口函数
func main() {
	chapterName := "第二百三十六"
	string2.Chinese2Int(chapterName)

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
