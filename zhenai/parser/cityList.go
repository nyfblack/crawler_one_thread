package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//城市解析器：解析城市列表页面的内容
func CityListParser(context []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	result := re.FindAllSubmatch(context, -1)

	var request engine.ParserResult
	for _, r := range result {
		//解析的内容
		request.Items = append(request.Items, "City: "+string(r[2]))
		request.Requests = append(request.Requests, engine.Request{
			Url:    string(r[1]),
			Parser: CityParser,
		})
	}
	return request
}
