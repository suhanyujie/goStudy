package main

import (
	"fmt"
	"time"
	"container/list"
)

func test1(num1 ...int) {
	fmt.Println(num1[1])
}

type Link struct {
	Ele list.Element

}

func generateList() {
	l := list.New()
	l.PushBack(12)
	l.PushBack(42)
	for i:=3;i<12;i++ {
		l.PushFront(i)
		//l.InsertBefore(i, l.Back())
	}
	l.InsertAfter(123, l.Back());
	fmt.Printf("链表的第一个元素值为：%d\n", l.Back().Value)
}

// 入口函数
func main() {
	start := time.Now()
	generateList();
	mapA := map[int]func()int {
		1:func()int{return 1},
		2:func()int {return 2},
		3:func()int {return 3},
	}
	for key1,_ := range mapA{
		value,isPresent := mapA[key1]
		if isPresent == true {
			fmt.Println("exist");
		}else{
			fmt.Println("not exist");
		}
		fmt.Println(value);
	}

	weeks := map[string]string{
		"Monday":"周1",
		"Thuesday":"周2",
		"Wednsday":"周3",
		"Thursday":"周4",
		"Friday":"周5",
	}
	for index,value := range weeks {
		if index=="Thuesday" || index=="Wednsday" {
			fmt.Println(index, value);
			continue;
		}
		fmt.Println(value);
	}




	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("cost time is %s\n", delta)

}
