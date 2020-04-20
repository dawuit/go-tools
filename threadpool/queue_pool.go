package threadpool

type QueueThreadPool struct {
	ch chan bool
	workQueue chan func()
	isClose chan bool
}

func NewQueueThreadPool(poolSize, queueSize int32) *QueueThreadPool {
	if poolSize <= 0 {
		poolSize = 1
	}
	if queueSize <= 0 {
		queueSize = 1
	}
	return &QueueThreadPool{
		ch: make(chan bool, poolSize),
		workQueue: make(chan func(), queueSize),
		isClose: make(chan bool),
	}
}

func (q *QueueThreadPool)Execute(run func()) {
	if q.IsClose() {
		return
	}

	q.workQueue <- run

	select {
	case q.ch <- true:
		go func() {
			for{

				if q.IsClose() {
					return
				}

				select {
				case w := <- q.workQueue:
					w()
				default:
					<- q.ch
					return
				}

			}
		}()
	default:
		return
	}
}

func (q *QueueThreadPool)Close(){
	if q.IsClose() {
		return
	}
	close(q.isClose)
}

func (q *QueueThreadPool)IsClose() bool{
	select {
	case <- q.isClose:
		return true
	default:
		return false
	}
}