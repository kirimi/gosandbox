package main

type Task func(int) string

type WorkersPoll struct {
	maxWorkers int
	task       Task
	in         chan int
	out        chan string
	workers    chan interface{}
}

func NewPool(maxWorkers int, task Task) *WorkersPoll {
	return &WorkersPoll{
		maxWorkers: maxWorkers,
		task:       task,
		in:         make(chan int),
		out:        make(chan string),
		workers:    make(chan interface{}, maxWorkers),
	}
}

func (wp *WorkersPoll) doJob(value int) {

forloop:
	for {
		select {
		case wp.workers <- 1:
			break forloop
		}
	}

	go func(value int) {
		result := wp.task(value)
		wp.out <- result
		<-wp.workers
	}(value)
}
