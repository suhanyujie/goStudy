package testSpider

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"os"
	"practice/http/file"
	"errors"
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
)

var BaseUrl string;

type FictionOneOfList struct {
	title string
	url   string
}

/**
获取小说列表页
 */
func GetList(url string) (status interface{}, returnVarerror error) {
	result, err := GetHttpResponse(url)
	if err != nil {
		return 30019, err
	}
	//解析html内容
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		return 30024, err
	}
	var detailContent []FictionOneOfList
	dom.Find("#list dl dd a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		url, _ := s.Attr("href")
		oneList := FictionOneOfList{
			title,
			BaseUrl + "/" + url,
		}
		detailContent = append(detailContent, oneList)
	});
	fmt.Println(detailContent)

	return nil, errors.New("任务完成")
}

/**
@desc 获取小说的详情
 */
func GetDetail(url string) (status interface{}, returnErr error) {
	result, err := GetHttpResponse(url)
	if err != nil {
		return 30021, err
	}
	//解析html内容
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		return 30026, err
	}
	var detailContent string
	dom.Find(".box_con #content").Each(func(i int, s *goquery.Selection) {
		detailContent = s.Text()
	});
	dir, _ := os.Getwd()
	newFileName := dir + "/src/practice/http/testSpider/cache/fiction.inc"
	//将内容写入文件中
	fileContent := []byte(detailContent)
	err = file.SsWriteFile(newFileName, fileContent)
	if err != nil {
		return 30036, err
	}

	return nil, errors.New("任务完成...")
}

/**
获取url公共的前缀部分
 */
func GetBaseUrl(baseUrl string) {
	BaseUrl = baseUrl
}

/**
根据url请求对应的详细并返回
 */
func GetHttpResponse(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(err.Error() + "-------[30015]")
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_13_6) AppleWebKit/537.36 (KHTML,Like Gecko)Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New(err.Error() + "-------[30021]")
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New(fmt.Sprintf("改请求的状态码为：%d\n", response.StatusCode))
	}
	utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	return ioutil.ReadAll(utf8Content)
}