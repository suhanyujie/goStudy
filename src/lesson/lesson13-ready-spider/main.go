package main

import (
	"lesson/lesson13-ready-spider/engine"
	"lesson/lesson13-ready-spider/zhenai/parser"
	"lesson/lesson13-ready-spider/scheduler"
)

// 获取并打印 城市第一页的用户信息
// 入口函数
func main() {
	url := "http://www.zhenai.com/zhenghun"
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:3,
	}
	e.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCityList,
	})
}
