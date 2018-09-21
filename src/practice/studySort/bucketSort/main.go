package main

import "fmt"

/**
桶排序
	桶排序的动效  https://www.cs.usfca.edu/~galles/visualization/BucketSort.html

	JavaScript的实现  http://bubkoo.com/2014/01/15/sort-algorithm/bucket-sort/

 */

// 入口函数
func main() {
	var numArr = []int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
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

	fmt.Println(max)
	fmt.Println(min)
}

/**
数字对应的桶索引计算
 */
func getBucketIndex(num, maximum, numberCount int) int {
	return num * numberCount / (maximum + 1)
}
