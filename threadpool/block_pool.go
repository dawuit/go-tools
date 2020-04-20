package threadpool

type Task func(func())

//没有队列的阻塞协程池。如果协程池已满则调用的协程会阻塞
func NewBlockThreadPool(size int64) Task {
	ch := make(chan bool, size)
	task := func(run func()) {
		ch <- true
		go func() {
			run()
			<- ch
		}()
	}
	return task
}