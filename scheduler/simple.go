package scheduler

import "crawler/engine"

type SimpleScheduler struct{}

var in chan engine.Request

func (SimpleScheduler) ConfigerWorkerInChan(ic chan engine.Request) {
	in = ic
}

func (SimpleScheduler) Submit(request engine.Request) {
	go func() { in <- request }()
}
