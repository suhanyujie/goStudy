package main

import "fmt"

/**
快速排序
	快速排序基于算法中很重要的思想是 分治.
	可以看这个哨兵移动的例子便于理解：https://blog.csdn.net/ricardo18/article/details/78867143


 */
// 入口函数
func main() {
	QuickSort()
}

func QuickSort() {
	var numArr = []int{21, 32, 19, 56, 29, 37, 16, 91, 126, 69}
	numLength := len(numArr)
	Qsort(&numArr, 0, numLength-1)
}

func Qsort(arr *[]int, p, r int) {
	if p > r {
		return;
	}
	pivotIndex := Partition(arr, p, r)
	fmt.Println(pivotIndex)
	Qsort(arr, 0, pivotIndex-1)
	Qsort(arr, pivotIndex+1, r)
	fmt.Println(arr)
}

func Partition(arr *[]int, p, r int) int {
	var dataArr *[]int;
	dataArr = arr
	pivot := (*dataArr)[p]
	//r表示数组的长度，最后一个单元就是arr[r-1]
	//对于l1的计算，第arr[p]是用作基准数的
	r1 := r - 1
	l1 := p + 1
	for r1 != l1 {
		for ; r1 >= l1 && (*arr)[r1] > pivot; {
			r1--
		}
		for ; l1 <= r1 && (*arr)[l1] < pivot; {
			l1++
		}
		if l1 < r1 {
			(*arr)[l1], (*arr)[r1] = (*arr)[r1], (*arr)[l1]
		}
		fmt.Printf("(*arr)[l1]:%d, (*arr)[r1]:%d\n", (*arr)[l1], (*arr)[r1])
	}
	(*arr)[l1], (*dataArr)[p] = (*dataArr)[p], (*arr)[l1]

	return 0
}
