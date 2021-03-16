package main

import (
	"fmt"
	"errors"
	"time"
)

//todo    golang实现连接池  channel + goroutine

//worker pool的设计常用来加速处理执行较耗时的重任务，且为了避免goroutine的过度创建，需要指定工作池的大小。
//使用golang的goroutine与chan，数行代码即可实现一个简单的工作池。
//新建两个channel，一个是works chan，一个是results chan，然后调用startWorkerPool启动指定goroutine个数的工作池

//todo 简单
/*
var res chan int
func main(){
	works := make(chan int,10)
	res = make(chan int,10)

	startPool(works, 3)
	for i := 1; i <= 5; i++ {
		works <- i
	}
	close(works)
	// waiting for results
	for i := 0; i < 5; i++ {
		<-res
	}
}
func  startPool(works chan int,size int){
	//启动多协程
	for i:= 1; i<= size; i++ {
		go worker(works,i)
	}
}

func worker(works chan int,g int ){
	for work := range works{
		 do(work,g)
	}
}

func do(work int ,g int) {
	time.Sleep(time.Second)
	fmt.Printf("goroutine %d done work %d\n", g, work)
	res <- work
}
*/

//todo 复杂

type WorkerPool struct {
	PoolSize    int
	tasksSize   int
	tasksChan   chan Task
	//resultsChan chan Task
	//Results     func() []Task
}

type Task struct {
	Id  int
	Err error
	f  func() error
}

func (t *Task)Do()error{
	return t.f()
}

func NewWorkerPool(tasks []Task, size int) *WorkerPool {
	tasksChan := make(chan Task, len(tasks))
	//resultsChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)
	pool := &WorkerPool{PoolSize: size, tasksSize: len(tasks), tasksChan: tasksChan}
	//pool.Results = pool.results
	return pool
}


func (pool *WorkerPool) Start() {
	for i := 0; i < pool.PoolSize; i++ {
		go pool.worker()
	}
}


func (pool *WorkerPool) worker() {
	for task := range pool.tasksChan {
		task.Err = task.Do()
		//pool.resultsChan <- task
		res_task <- task
	}
}

/*
func (pool *WorkerPool) results() []Task {
	tasks := make([]Task, pool.tasksSize)
	for i := 0; i < pool.tasksSize; i++ {
		tasks[i] = <-pool.resultsChan
	}
	return tasks
}
*/

var res_task chan Task
func main(){

	t := time.Now()
	tasks := []Task{
		{Id: 0, f: func() error { time.Sleep(2 * time.Second); fmt.Println(0); return nil }},
		{Id: 1, f: func() error { time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
		{Id: 2, f: func() error { fmt.Println(2); return errors.New("error") }},
	}

	res_task = make(chan Task)
	pool := NewWorkerPool(tasks, 2)
	pool.Start()

	//tasks = pool.Results()
/*
	for task := range res_task {
		fmt.Printf("result of task %d is %v\n", task.Id, task.Err)
	}
*/
	for i := 0; i < len(tasks); i++ {
		task := <-res_task
		fmt.Printf("result of task %d is %v\n", task.Id, task.Err)
	}

	fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
}





