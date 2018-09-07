package string

import (
	"strings"
	"regexp"
)

/**
算法的实现逻辑可以参考： https://segmentfault.com/a/1190000004881457
设计思想：
	将中文数学转换成阿拉伯数字。
	将中文权位转换成10的位数。
	对每个权位依次转换成位数并求和。
	零直接忽略即可。

解决的问题：
1.一旦字符串中 含有非数字 非数学单位的文字 解析会出问题 已解决
2.字符串为十，解析出问题。直接判断第0个位置如果是权重单位，则特殊处理 已解决
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

const ChineseNumPattern = `([零一二三四五六七八九壹贰叁肆伍陆柒捌玖两十百千万亿]+)`

/**
中文数字转阿拉伯数字
第二百三十六->236

1.先将字符串通过空字符串分隔开，成为数组，针对单个文字进行处理
2.如果是数字，则转为数字
3.如果是"数字位"，则将之前的数字乘以权重 与后续的section进行相加
 */
func Chinese2Int(originStr string) int {
	//在原字符串中，将中文数字转换为阿拉伯数字
	var dealStr string;
	dealStr = originStr;
	//将字符串中跟数字无关的字符去除
	re := regexp.MustCompile(ChineseNumPattern)
	result := re.FindAllStringSubmatch(dealStr, 1)
	dealStr = result[0][1]
	dealStrArr := strings.Split(dealStr, "")
	var number, sectionNum, rtn int;
	for index, val := range dealStrArr {
		oneNum, isOk := chNumChar[val]
		//中文转数字是否匹配上，没匹配上说明是单位或者其他文字
		if isOk {
			number = oneNum
			//如果是最后一位，是各位，只需加上它
			if (index) == (len(dealStrArr) - 1) {
				sectionNum += oneNum
			}
		} else {
			_, isOk := chNumMapValue[val]
			if !isOk {
				continue
			}
			var unit = chNumMapValue[val].value
			var secUnit = chNumMapValue[val].secUnit
			if secUnit {
				sectionNum = (sectionNum + number) * unit
				rtn += sectionNum
				sectionNum = 0
			} else {
				//如果第一个文字就是数学单位，则直接使用1*对应的数学单位
				if index == 0 {
					sectionNum += 1 * unit
				} else {
					sectionNum += number * unit
				}
			}
			number = 0
		}
	}

	return rtn + sectionNum
}
