package main

import "fmt"

/**
插入排序：
	这种排序方法 比冒泡、选择排序要稍微难理解一点
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
