package main

import (
	"fmt"
	"math"
)

// byte
// rune类似char，32位
// float32 float64
// complex64 实部和虚部都是32位
//complex128 实部和虚部都是64位

func enums() {
	const (
		cpp    = 0
		java   = 1
		php    = 2
		golang = 3
	)
	const (
		b = 2 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
	)
	fmt.Println(cpp, java)
	fmt.Println(b, kb, mb, gb, tb, pb, eb)
}

// 类型强制转换
func func1() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 条件语句
func func2(score int) string {
	if class := "Unknow"; score > 80 {
		class = "A"
		fmt.Println("your score are great!")
		return class
	}

	panic("wrong score -2...")
	return "Unknow"
}

func main() {
	var res string = func2(81)
	fmt.Println(res)

	//func2(-3);
	func1()
	enums()
}
