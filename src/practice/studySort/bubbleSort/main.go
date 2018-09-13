package main

import "fmt"

/**
冒泡排序：
	原理：从第一个数开始，挨个与其后所有的数对比一次，遇到比之小的，则互换，以保证用于对比的数是最小的

	详情可以参考 https://juejin.im/post/5b9492def265da0aff171b94?utm_source=gold_browser_extension
 */
// 入口函数
func main() {
	var numArr = [10]int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}

	RealBubbleSort(numArr)

	//numLength := len(numArr)
	//for i := 0; i < numLength; i++ {
	//	for j := i + 1; j < numLength; j++ {
	//		if numArr[i] > numArr[j] {
	//			//互换2个值
	//			numArr[i], numArr[j] = numArr[j], numArr[i]
	//		}
	//	}
	//}

	//fmt.Println(numArr)
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
