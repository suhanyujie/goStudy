package main

import "fmt"

/**
选择排序：
	每一次比较都能找到最大/最小的那个数
 */
func main() {
	var numArr = [10]int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	numLength := len(numArr)
	//z最后一个数不需要再进行比较 所以是numLength-1
	for i := 0; i < numLength-1; i++ {
		for j := i + 1; j < numLength; j++ {
			if numArr[i] > numArr[j] {
				//互换2个值
				numArr[i], numArr[j] = numArr[j], numArr[i]
			}
		}
	}

	fmt.Println(numArr)
}
