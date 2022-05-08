package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Shceduler   Shceduler
	WorkerCount int
}

type Shceduler interface {
	Submit(Request)
	ConfigureWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Shceduler.Submit(r)
	}
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Shceduler.ConfigureWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}
		for _, request := range result.Requests {
			e.Shceduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
