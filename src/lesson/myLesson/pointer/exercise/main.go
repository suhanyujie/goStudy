package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}
/**
.的运算有优先级比 *  要高！
 */



// 入口函数
func main() {
	var p1 Stu;
	p1.Name = "suhanyu"
	p1.Age = 29
	var p2 *Stu = &p1
	fmt.Printf("p1的地址是%p\n", &p1)
	fmt.Printf("p1的值是%v\n", p1)
	fmt.Printf("p2的地址是%p\n", &p2)
	fmt.Printf("指针p2的值是%p\n", p2)
}
