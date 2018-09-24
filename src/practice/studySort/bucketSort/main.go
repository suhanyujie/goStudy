package main

import (
	"practice/studySort/bucketSort/dataStruct"
)

/**
桶排序
	桶排序的动效  https://www.cs.usfca.edu/~galles/visualization/BucketSort.html

	JavaScript的实现  http://bubkoo.com/2014/01/15/sort-algorithm/bucket-sort/


	先预备好固定数量的桶，每个待排的数都通过一个函数计算得出对应的桶编号
	每个桶中可以放多个数，以链表的数据结构进行存储

 */

const BucketNum = 10

// 入口函数
func main() {
	var numArr = []int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	BucketSort(numArr)
}

// 输入一组数 对他们进行排序 并输出
func BucketSort(numArr []int) {
	numLength := len(numArr)
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
	//预先准备好一定数量的桶
	var buckets [BucketNum]*dataStruct.LinkNode
	for i := 0; i < BucketNum; i++ {
		buckets[i] = dataStruct.GetNode(0)
	}
	for _, value := range numArr {
		bucketIndex := getBucketIndex(value, max, numLength)
		var newNode *dataStruct.LinkNode = new(dataStruct.LinkNode)
		newNode.Data = value
		buckets[bucketIndex] = buckets[bucketIndex].Push(newNode)
	}
	//打印输出
	dataStruct.PrintLink(buckets[1])
}

/**
数字对应的桶索引计算
 */
func getBucketIndex(num, maximum, numberCount int) int {
	return num * numberCount / (maximum + 1)
}
