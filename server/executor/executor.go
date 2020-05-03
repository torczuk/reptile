package executor

import (
	"log"
	"time"
)

type fn func()

type Executor struct {
	task     fn
	duration time.Duration
	done     chan bool
}

func NewExecutor(task fn, duration time.Duration) *Executor {
	return &Executor{task, duration, make(chan bool, 1)}
}

func (e *Executor) Start() {
	log.Printf("starting executor ...")
	ticker := time.NewTicker(e.duration)
	defer ticker.Stop()
	for {
		select {
		case <-e.done:
			log.Printf("stoping executor ...")
			return
		case <-ticker.C:
			log.Printf("execute task ...")
			e.task()
		}
	}

}

func (e *Executor) Stop() {
	e.done <- true
	close(e.done)
}
