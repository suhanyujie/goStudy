package main

import "fmt"

/**
插入排序：
	这种排序方法 比冒泡、选择排序要稍微难理解一点
	插入排序如果很难理解，可以查看这个页面中的动图。https://blog.csdn.net/YuYunTan/article/details/52026857

第1次
	32先和21比较，32大，不用处理，继续和下标为i - 1的数比较，但此时i-1小于0，不满足
第2次
	19和32比较，19小，将32和19位置的值互换，此时再拿19和21比较，19小，将19和21的值互换



 */

// 入口函数
func main() {
	var numArr = [10]int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	numLength := len(numArr)
	var j,compNumber int;
	for i := 1; i < numLength; i++ {
		compNumber = numArr[i]
		for j = i - 1; j >= 0 && compNumber < numArr[j];  {
			numArr[j+1], numArr[j] = numArr[j], numArr[j+1]
			j--
		}
		numArr[j+1] = compNumber
	}

	fmt.Println(numArr)
}
