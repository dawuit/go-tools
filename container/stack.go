package container

type IntStack []int

func NewIntStack() IntStack {
	return make([]int, 0, 16)
}

func (s *IntStack)Push(val int) {
	*s = append(*s, val)
}

func (s *IntStack)Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}