package main

import (
	"regexp"
	"fmt"
)

const text = `My name is ccmous@51zhaoyou.com
ad  avc@51zhaoyou.com
dsrdfd@51zhaoyou.com`

// 入口函数
func main() {
	re := regexp.MustCompile(`([a-z]+)@[a-z0-9]+.[a-zA-Z0-9]+`)
	match := re.FindAllStringSubmatch(text, -1)
	for _,m:= range match{
		fmt.Println(m)
	}
	//fmt.Println(match);
}


