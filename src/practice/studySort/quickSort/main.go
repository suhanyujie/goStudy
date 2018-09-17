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
	Qsort(numArr, 0, numLength-1)
}

func Qsort(arr []int, p, r int) {
	numLength := len(arr)
	pivotIndex := Partition(arr,0, numLength-1)
	fmt.Println(pivotIndex)
}

func Partition(arr []int, p, r int) int {
	pivot := arr[p]
	r1 := r - 1
	l1 := p + 1
	for r1>=l1 {
		for ; r1 >= l1; r1-- {
			if arr[r1] < pivot {
				break;
			}
		}
		for ;l1<=r1 ;l1++  {
			if arr[l1]>pivot {
				break;
			}
		}
		if l1 == r1 {
			return l1
			//arr[p],arr[l1] = arr[l1],arr[p]
		}
		arr[l1],arr[r1] = arr[r1],arr[l1]
	}
	return 0
}
