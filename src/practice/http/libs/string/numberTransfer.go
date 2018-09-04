package string

import (
	"strings"
)

/**
算法的实现逻辑可以参考： https://segmentfault.com/a/1190000004881457
设计思想：
	将中文数学转换成阿拉伯数字。
	将中文权位转换成10的位数。
	对每个权位依次转换成位数并求和。
	零直接忽略即可。


 */
var chNumChar = map[string]int{
	"零": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9,
	"壹": 1, "贰": 2, "叁": 3, "肆": 4, "伍": 5, "陆": 6, "柒": 7, "捌": 8, "玖": 9,
	"两": 2,
}

type valueObj struct {
	value   int
	secUnit bool
}

var chNumMapValue = map[string]valueObj{
	"十": valueObj{10, false},
	"百": valueObj{100, false},
	"千": valueObj{1000, false},
	"万": valueObj{10000, true},
	"亿": valueObj{100000000, true},
}

func Chinese2Int(originStr string) int {
	//在原字符串中，将中文数字转换为阿拉伯数字
	var dealStr string;
	dealStr = originStr;
	dealStrArr := strings.Split(dealStr, "")
	var number, sectionNum, rtn int;
	for index, val := range dealStrArr {
		oneNum, isOk := chNumChar[val]
		//中文转数字是否匹配上，没匹配上说明是单位或者其他文字
		if isOk {
			//如果是最后一位，是各位，只需加上它
			if (index - 1) == len(dealStrArr) {
				sectionNum += oneNum
			} else {
				number = oneNum
			}
		} else {
			_,isOk := chNumMapValue[val]
			if !isOk {
				continue
			}
			var unit = chNumMapValue[val].value
			var secUnit = chNumMapValue[val].secUnit
			if secUnit {
				sectionNum = sectionNum + number*unit
				rtn += sectionNum
				sectionNum = 0
			} else {
				sectionNum += number * unit
			}
		}
	}

	return rtn + sectionNum
}
