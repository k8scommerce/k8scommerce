package workerpool

type WorkerPool interface {
	StartWorkers()
	AddTask(task func())
}

func NewWorkerPool(maxWorkers int) WorkerPool {
	return &workerPool{
		maxWorkers:   maxWorkers,
		queuedTasksC: make(chan func()),
	}
}

type workerPool struct {
	maxWorkers   int
	queuedTasksC chan func()
}

func (wp *workerPool) StartWorkers() {
	wp.startWorkers()
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTasksC <- task
}

func (wp *workerPool) GetTotalQueuedTasks() int {
	return len(wp.queuedTasksC)
}

func (wp *workerPool) startWorkers() {
	for i := 0; i < wp.maxWorkers; i++ {
		go func(workerId int) {
			for task := range wp.queuedTasksC {
				// logx.Infof("starting worker #%d", workerId)
				task()
			}
		}(i + 1)
	}
}
