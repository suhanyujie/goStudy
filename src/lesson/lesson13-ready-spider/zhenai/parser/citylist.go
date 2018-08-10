package parser

import (
	"regexp"
	"lesson/lesson13-ready-spider/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			ParseFunc:engine.NilParser,
		})
		// fmt.Printf("Url is %s,CityName is %s \n",m[1],m[2])
	}
	// fmt.Printf("城市个数为：%d\n", len(matches))
	return result
}