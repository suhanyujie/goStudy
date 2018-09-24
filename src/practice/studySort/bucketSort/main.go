package main

import (
	"fmt"
	"os"
	"practice/studySort/bucketSort/dataStruct"
)

/**
桶排序
	桶排序的动效  https://www.cs.usfca.edu/~galles/visualization/BucketSort.html

	JavaScript的实现  http://bubkoo.com/2014/01/15/sort-algorithm/bucket-sort/


	先预备好固定数量的桶，每个待排的数都通过一个函数计算得出对应的桶编号
	每个桶中可以放多个数，以链表的数据结构进行存储

 */

// 入口函数
func main() {
	var numArr = []int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	numLength := len(numArr)
	var link *dataStruct.LinkNode = dataStruct.GetNode(0)
	var newNode *dataStruct.LinkNode = new(dataStruct.LinkNode)
	for _,value := range numArr {
		newNode.Data = value;
		link.Push(newNode)
		break
	}
	fmt.Printf("%p,%T\n",newNode,newNode)
	//断点
	os.Exit(5)





	var max, min int;
	max = numArr[0];
	min = numArr[0];
	//找出最大、最小的数
	for i := 0; i < numLength; i++ {
		if numArr[i] > max {
			max = numArr[i]
		}
		if numArr[i] < min {
			min = numArr[i]
		}
	}

	fmt.Println(max)
	fmt.Println(min)
}

/**
数字对应的桶索引计算
 */
func getBucketIndex(num, maximum, numberCount int) int {
	return num * numberCount / (maximum + 1)
}
