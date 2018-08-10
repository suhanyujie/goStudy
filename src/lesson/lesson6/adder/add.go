package adder

import (
	"fmt"
	"io"
	"strings"
	"bufio"
)

// ad adder
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
// an data type
type intGen func() int

// 斐波那契函数
func fabonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// todo 从reader中读取一些具体内容
func (_this intGen) Read(p []byte) (n int,err error) {
	next := _this()
	if next > 10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 入口函数
func main() {
	//a := adder()
	a := fabonacci()
	for i := 0; i < 10; i++ {
		//fmt.Println(a())
		// fmt.Printf("0+...+%d = %d\n",i,a(i))
	}
	printFileContents(a)
}
