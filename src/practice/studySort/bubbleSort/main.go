package main

import "fmt"

/**
冒泡排序：
	原理：相邻的两个数进行比较，一旦遇到小的，则进行交换。以此不断的往后递推

	详情可以参考 https://juejin.im/post/5b9492def265da0aff171b94?utm_source=gold_browser_extension
 */
// 入口函数
func main() {
	var numArr = [10]int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	RealBubbleSort(numArr)
}

/**
真正的冒泡排序实现。。。。
 */
func RealBubbleSort(numArr [10]int) {
	numLength := len(numArr)
	for i := 0; i < numLength; i++ {
		for j := 0; j < (numLength - i - 1); j++ {
			if numArr[j] > numArr[j+1] {
				//互换2个值
				numArr[j], numArr[j+1] = numArr[j+1], numArr[j]
			}
		}
		fmt.Println(numArr)
	}

	fmt.Println("最终结果如下：")
	fmt.Println(numArr)
}
