package engine

import (
	"log"
)

type ConcurrentEngin struct {
	Scheduler   Schedul
	WorkerCount int
}

type Schedul interface {
	Submit(Request)
	ConfigerWorkerInChan(chan Request)
}

func (c ConcurrentEngin) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	c.Scheduler.ConfigerWorkerInChan(in)
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	for {
		count := 0
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d : %v", count, item)
			count++
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		request := <-in
		ps, err := worker(request)
		if err != nil {
			return
		}
		out <- ps
	}()
}
