package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	// "io/ioutil"
	iconv "github.com/djimenez/iconv-go"
	"net/http"
)

const (
	baseUrl = "http://www.biquzi.com"
)

func main() {
	url := "http://www.biquzi.com/14_14297/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	jqRes, err := goquery.NewDocumentFromReader(resp.Body)

	jqRes.Find("#list a").First().Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Text()
		tmpUrl, isOk := s.Attr("href")
		if !isOk {
			tmpUrl = ""
		}
		tmpUrl = baseUrl + tmpUrl
		title, _ = iconv.ConvertString(title, "gbk", "utf-8")
		// 存入数据库
		fmt.Printf("Review %d: %s-%s\n", i, title, tmpUrl)
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	// statusCode := resp.StatusCode
}
