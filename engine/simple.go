package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds ...Request) {
	//任务队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ps, err := worker(r)
		if err != nil {
			continue
		}
		count := 0
		requests = append(requests, ps.Requests...)
		for _, item := range ps.Items {
			log.Printf("Got item #%d : %v", count, item)
			count++
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("fetcher Url : %s", r.Url)
	context, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher context is error : %s", err)
		return ParserResult{}, err
	}

	return r.Parser(context), nil
}
