package testSpider

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	//"os"
	//"practice/http/file"
	"errors"
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"sync"
	"time"
	"lesson/gorm/beegoOrm"
	"lesson/gorm/beegoOrm/models"
	string2 "practice/http/libs/string"
	"strings"
	"strconv"
	"math/rand"
)

var BaseUrl string;

/**
处理任务数据
 */
func DealTask(ch1 chan beegoOrm.FictionOneOfList, wg sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case oneTask := <-ch1:
			//产生一个随机的秒数延迟 避免被加入黑名单
			ms := rand.Intn(500) + 1000
			time.Sleep(time.Millisecond * time.Duration(ms))
			fmt.Println("已经sleep完毕")
			fmt.Println(oneTask.Title)
			status,detailContent, returnErr := GetDetail(oneTask.Url)
			if status != nil {
				log.Fatal(returnErr)
			}
			//将数据存放到数据库中
			detail := &models.NovelContent{
				ListId:uint(oneTask.InsertId),
				Chapter:oneTask.Chapter,
				Title:oneTask.Title,
				Content:detailContent,
				WorkerId:0,
				ErrFlag:1,
			}
			insertId, _ := beegoOrm.InsertDetail((*detail))
			fmt.Printf("详细内容爬取完成：%d\n", insertId)
		}
	}
	wg.Done()
}

/**
获取小说列表页
 */
func GetList(chTask chan beegoOrm.FictionOneOfList,url string) (status interface{}, data []beegoOrm.FictionOneOfList, returnVarerror error) {
	result, err := GetHttpResponse(url)
	if err != nil {
		return 30019, nil, err
	}
	//解析html内容
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		return 30024, nil, err
	}
	var listArr []beegoOrm.FictionOneOfList
	dom.Find("#list dl dd a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		//如果不包含 第、章等字样，则跳过
		if !strings.Contains(title,"第") || !strings.Contains(title,"章") {
			return
		}
		url, _ := s.Attr("href")
		oneList := &beegoOrm.FictionOneOfList{
			title,
			BaseUrl + "/" + url,
			0,
			string2.Chinese2Int(title),
		}
		//查询这个章节的是否已经存在于数据库中
		where := map[string]string{
			"novel_id":"11",
			"chapter":strconv.Itoa(oneList.Chapter),
		}
		existList,err := beegoOrm.QueryList(where)
		log.Println(existList)
		if err!=nil {
			log.Println(err)
		}
		if existList.Id>0 {
			fmt.Println("["+oneList.Title+"]跳过这一次插入...")
			return
		}
		insertId, msg := beegoOrm.InsertList(*oneList)
		oneList.InsertId = insertId
		chTask<-(*oneList)
		//listArr = append(listArr, *oneList)
		log.Printf("[%d]%s\n", insertId, msg)
		time.Sleep(time.Millisecond * 100)
	});
	return nil, listArr, errors.New("任务完成")
}

/**
@desc 获取小说的详情
 */
func GetDetail(url string) (status interface{}, content string, returnErr error) {
	result, err := GetHttpResponse(url)
	if err != nil {
		return 30021,"", err
	}
	//解析html内容
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		return 30026,"", err
	}
	var detailContent string
	dom.Find(".box_con #content").Each(func(i int, s *goquery.Selection) {
		detailContent = s.Text()
	});
	//fmt.Println(detailContent)
	//dir, _ := os.Getwd()
	//newFileName := dir + "/src/practice/http/testSpider/cache/fiction.inc"
	//将内容写入文件中
	//fileContent := []byte(detailContent)
	//err = file.SsWriteFile(newFileName, fileContent)
	//if err != nil {
	//	return 30036,"", err
	//}

	return nil,detailContent, errors.New("任务完成...")
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
	//fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New(fmt.Sprintf("该请求的状态码为：%d\n", response.StatusCode))
	}
	utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	return ioutil.ReadAll(utf8Content)
}
