package container

type IntQueue []int

func NewIntQueue() IntQueue {
	return make([]int, 0, 16)
}

func (q *IntQueue)Push(val int) {
	*q = append(*q, val)
}

func (q *IntQueue)Pop() int {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *IntQueue)Len() int{
	return len(*q)
}