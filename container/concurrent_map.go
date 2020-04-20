package container

import "sync"

type ConcurrentMap struct{
	sync.Mutex
	m map[string]int
}

func NewConcurrentMap() *ConcurrentMap{
	return &ConcurrentMap{
		m: make(map[string]int),
	}
}

func (c *ConcurrentMap)put(k string, v int) {
	c.Lock()
	defer c.Unlock()
	c.m[k] = v
}

func (c *ConcurrentMap)remove(k string) {
	c.Lock()
	defer c.Unlock()
	delete(c.m, k)
}
