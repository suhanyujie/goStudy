package anonymousFields

import (
	"fmt"
)

type inners struct {
	in1 int
	in2 int
}

type outers struct {
	b int
	c float32
	int    //anonymous fields
	inners //anonymous fields
}

// todo
func (_this *inners) ChangeInner1(num1 int) {
	_this.in1 = num1
}

// todo
func (_this *inners) String() {
	fmt.Printf("the value is {in1:%d,in2:%d}\n", _this.in1,_this.in2)
}

func Test1() {
	inner1 := &inners{0, 0}
	inner1.ChangeInner1(123123)
	inner1.String();

	//outer := new(outers)
	//outer.c = 3.323
	//outer.b = 123
	//outer.int = 12312
	//outer.inners.in1 = 12
	//outer.inners.in2 = 13
	//fmt.Printf("the values is %v\n", outer)
}
