package main

/**
插入排序：

 */

// 入口函数
func main() {
	var numArr = [10]int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	numLength := len(numArr)
	var newNumArr [...]int
	for i := 0; i < numLength; i++ {
		for j := i + 1; j < numLength; j++ {

		}
		if numArr[i] > numArr[i+1] {
			numArr[i], numArr[i+1] = numArr[i+1], numArr[i]
		}
	}
}
