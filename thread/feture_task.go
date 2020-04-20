package thread

type FetureTask struct {
	ch chan interface{}
}

func StartTask(run func() interface{}) *FetureTask{
	fetureTask := FetureTask{
		ch : make(chan interface{}, 1),
	}
	go func() {
		fetureTask.ch <- run()
	}()
	return &fetureTask
}

func (f *FetureTask)Get() interface{} {
	return <-f.ch
}
