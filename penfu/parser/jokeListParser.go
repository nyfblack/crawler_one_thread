package parser

import (
	"crawler/engine"
	"regexp"
)

const listRe = `<a href="(https://www.pengfu.com/[\da-zA-Z_.]+)"[^>]*>([^<]+)</a>`

func JokeListParser(context []byte) engine.ParserResult {
	re := regexp.MustCompile(listRe)
	match := re.FindAllSubmatch(context, -1)

	var result engine.ParserResult
	for _, m := range match {
		if len(m) >= 3 {
			title := string(m[2])
			result.Items = append(result.Items, title)
			result.Requests = append(result.Requests, engine.Request{
				Url: string(m[1]),
				Parser: func(context []byte) engine.ParserResult {
					return JokeParser(context, title)
				},
			})
		}
	}
	return result
}
