package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "who are you ? I do not know.how about you"
	res := CountWord(str1)
	fmt.Println(res)
}

func CountWord(s string) map[string]int {
	vmap := make(map[string]int)
	strs := strings.Fields(s)
	length := len(strs)
	for i := 0; i < length; i++ {
		vmap[strs[i]] = vmap[strs[i]] + 1
	}
	return vmap
}
