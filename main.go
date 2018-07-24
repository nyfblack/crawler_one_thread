package main

import (
	"crawler/engine"
	"crawler/penfu/parser"
)

const zhenai = "http://www.zhenai.com/zhenghun"
const penfu = "https://www.pengfu.com/xiaohua_1.html"

func main() {
	/*engine.Run(engine.Request{
		Url:    zhenai,
		Parser: parser.CityListParser,
	})*/
	engine.Run(engine.Request{
		Url:    penfu,
		Parser: parser.JokeListParser,
	})
}
