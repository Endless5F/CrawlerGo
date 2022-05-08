package scheduler

import "GoCodeProject/crawler/engine"

type SimpleSheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleSheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleSheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
