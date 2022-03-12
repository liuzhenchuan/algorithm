package goroutinePool

import "fmt"

// Task 定义一个任务
type Task struct {
	f func() error
}

// NewTask 创建一个任务
func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

// Execute 执行任务
func (t *Task) Execute() {
	t.f()
}

// Pool 定义协程池
type Pool struct {
	// 对外接受task的入口
	EntryChannel chan *Task
	WorkerNum    int
	JobsChannel  chan *Task
}

// CreatePool 创建协程池
func CreatePool(workerNum int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		WorkerNum:    workerNum,
		JobsChannel:  make(chan *Task),
	}
	return &p
}

// Worker 协程池创建一个worker开始工作
func (p *Pool) Worker(workerId int) {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("worker ID : ", workerId, "执行完毕")
	}
}

// Run 启动协程池
func (p *Pool) Run() {
	for i := 0; i < p.WorkerNum; i++ {
		go p.Worker(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	close(p.JobsChannel)
}
