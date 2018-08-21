package main

import (
	"practice/http/testSpider"
	"github.com/astaxie/beego/logs"
	"practice/http/file"
		"log"
	"github.com/PuerkitoBio/goquery"
	"bytes"
			)

// 入口函数
func main() {
	url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
	result, err := testSpider.GetHttpResponse(url)
	if err != nil {
		logs.Error(err)
	}
	//解析html内容
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		logs.Error(err)
	}
	var detailContent string
	dom.Find(".box_con #content").Each(func(i int,s *goquery.Selection) {
		detailContent = s.Text()
	});
	//将内容写入文件中
	fileContent := []byte(detailContent)
	err = file.SsWriteFile("fiction.inc", fileContent)
	if err != nil {
		log.Fatal(err)
	}
	logs.Info("完成任务.....")
}
