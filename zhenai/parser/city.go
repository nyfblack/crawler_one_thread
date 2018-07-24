package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func CityParser(context []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	result := re.FindAllSubmatch(context, -1)

	var request engine.ParserResult
	for _, r := range result {
		name := string(r[2])
		//解析的内容
		request.Items = append(request.Items, "User: "+name)
		request.Requests = append(request.Requests, engine.Request{
			Url: string(r[1]),
			Parser: func(context []byte) engine.ParserResult {
				return ProfileParser(context, name)
			},
		})
	}
	return request
}
