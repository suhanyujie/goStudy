package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"regexp"
	//"reflect"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"bufio"
	"errors"
	//"practice/storage"
	"practice/storage"
	"reflect"
	"time"
	"os"
)

var (
	titlePattern = regexp.MustCompile(`<dd><a href="/[_\d]+/[\d]+.html">([^<]+)</a></dd>`)
	urlPattern = regexp.MustCompile(`<dd><a href="(/[_\d]+/[\d]+.html)">[^<]+</a></dd>`)

	baseUrl = "http://www.biquge.com.tw"
)

// todo
func Get(url string) (content string, err error) {
	errors.New("111----test")
	res, err := http.Get(url);
	if err != nil {
		log.Fatal("request status code error...");
	}
	defer res.Body.Close()
	bodyReader := bufio.NewReader(res.Body)
	utf8Reader := transform.NewReader(bodyReader, simplifiedchinese.GBK.NewDecoder())
	bytesContent, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		errors.New("transfrom error")
	}
	newStringContent := string(bytesContent[:]);
	//fmt.Println(newStringContent)

	return newStringContent, nil
}

// 入口函数
func main() {
	url := "http://www.biquge.com.tw/18_18820/"
	content, err := Get(url)
	if err != nil {
		log.Fatal(err)
	}

	matches := titlePattern.FindAllStringSubmatch(content, -1)
	for i := 0; i < len(matches); i++ {
		//fmt.Println(matches[i][1])
	}
	matches = urlPattern.FindAllStringSubmatch(content, -1)
	for i := 0; i < len(matches); i++ {
		//fmt.Println(matches[i][1])
	}
	sql := fmt.Sprintf(`insert into go_spider_list(website,fiction_id,title,url,status,add_time) values (1,1,"test","testUrl",1,"%s")`, time.Now().Format("2006-01-02 15:04:05"))
	storage := new(storage.StorageObj)
	storage.Connect()
	sql = "select * from go_spider_list where id=1"
	value := storage.GetRow(sql)
	fmt.Printf("value is %v\n", value)
	os.Exit(0)


	_ = storage.Insert(sql)

	//fmt.Println(insertId)
	//fmt.Println(storage.GetLastSql())
	fmt.Printf("content is %v\n", len(matches))
	fmt.Printf("type is %s\n", reflect.TypeOf(matches))
}
