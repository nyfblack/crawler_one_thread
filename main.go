package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

const zhenai = "http://www.zhenai.com/zhenghun"
const penfu = "https://www.pengfu.com/xiaohua_1.html"

func main() {
	engine.ConcurrentEngin{
		Scheduler:   scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}.Run(engine.Request{
		Url:    zhenai,
		Parser: parser.CityListParser,
	})
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:    penfu,
		Parser: parser.JokeListParser,
	})*/

}
