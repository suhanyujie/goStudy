package string

import "fmt"

var ChineseUnitMap map[string]string

func Chinese2Int() {
	ChineseUnitMap := map[string]string{
		"零": "0","一": "1","二": "2","三": "3", "四": "4", "五": "5", "六": "6", "七": "7", "八": "8", "九": "9",
		"壹": "1", "贰": "2", "叁": "3", "肆": "4", "伍": "5", "陆": "6", "柒": "7", "捌": "8", "玖": "9",
		"两": "2",
		"仟": "千", "佰": "百", "拾": "十",
		"万万": "亿",
	}

	fmt.Println(ChineseUnitMap)
}
