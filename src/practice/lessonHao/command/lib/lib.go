package lib

import (
	"fmt"
	in "practice/lessonHao/command/lib/internal"
)

type TestT1 struct {
	a int
	s1 string
}

// todo
func (_this TestT1) String() string  {
	var str1 string;
	//触发递归调用，因为要转换成年字符串，所以会触发调用结构体本省的String方法！！！
	//str1 = fmt.Sprintf("%v", _this)
	return str1
}

func ShowLib() {
	fmt.Println("this is lib folder...")
	in.ShowInternal()

	t1 := &TestT1{
		1,
		"suhanyu",
	}
	t1.String()
}
