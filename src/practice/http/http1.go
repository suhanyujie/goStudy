package main

import (
	string2 "practice/http/libs/string"
	"fmt"
	"practice/http/testSpider"
	"log"
	"github.com/astaxie/beego/logs"
	"sync"
)


var chTask = make(chan testSpider.FictionOneOfList)

// 入口函数
func main() {
	var wg sync.WaitGroup
	chapterName := "第一千二百零六"
	num := string2.Chinese2Int(chapterName)
	fmt.Println(num)

	/*   */
	listUrl := "https://www.biduo.cc/biquge/17_17308"
	//url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
	//status,err := testSpider.GetDetail(url)
	testSpider.GetBaseUrl(listUrl)
	status,data,err := testSpider.GetList(listUrl)
	if status!=nil {
		log.Fatal(err)
	}
	logs.Info(err)
	//将一个任务数据存入数据库
	go testSpider.DealTask(chTask,wg)
	// 将每一个任务数据放入数据库，
	// 然后放入任务队列中，获取具体的详情页内容
	wg.Add(1)
	go func() {
		var countTimes = 0;
		for index,oneTask := range data{
			chTask<-oneTask
			fmt.Println(index,oneTask)
			if countTimes >= 3 {
				break;
			}
			countTimes++;
		}
		wg.Done()
	}()
	wg.Wait()
}
