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
	numLength := len(*arr)
	pivotIndex := Partition(arr, 0, numLength)
	fmt.Println(pivotIndex)
	Qsort(arr, 0, pivotIndex)
	Qsort(arr, pivotIndex, numLength)
	fmt.Println(arr)
}

func Partition(arr *[]int, p, r int) int {
	var dataArr *[]int;
	dataArr = arr
	pivot := (*dataArr)[p]
	r1 := r - 1
	l1 := p + 1
	for r1 >= l1 {
		for ; r1 >= l1 && (*arr)[r1] > pivot; {
			r1--
		}
		for ; l1 <= r1 && (*arr)[l1] < pivot; {
			if l1 == r1 {
				(*arr)[l1] = pivot
				return l1
			}
			l1++
		}
		if l1 == r1 {
			(*arr)[l1] = pivot
			return l1
		}

		if l1 < r1 {
			(*arr)[l1], (*arr)[r1] = (*arr)[r1], (*arr)[l1]
		}
		fmt.Printf("(*arr)[l1]:%d, (*arr)[r1]:%d\n",(*arr)[l1], (*arr)[r1])
		if l1>r1 {
			break
		}
	}

	return 0
}
