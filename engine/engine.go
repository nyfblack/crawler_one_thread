package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	//任务队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetcher Url : %s", r.Url)
		context, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher context is error : %s", err)
			continue
		}

		ps := r.Parser(context)
		requests = append(requests, ps.Requests...)
		for _, item := range ps.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
