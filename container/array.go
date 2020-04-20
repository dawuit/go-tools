package container

import "errors"

type IntArray []int

func NewIntArray() IntArray {
	return make([]int, 0, 0)
}

func (a *IntArray)AddFirst(val int) {
	*a = append(*a, 0)
	copy((*a)[1:], *a)
	(*a)[0] = val
}

func (a *IntArray)AddLast(val int) {
	*a = append(*a, val)
}

func (a *IntArray)RemoveFirst() int {
	res := (*a)[0]
	*a = (*a)[1:]
	return res
}

func (a *IntArray)RemoveLast() int {
	res := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return res
}

func (a *IntArray)Insert(val, index int) error {
	if index < 0 || index > len(*a) {
		return errors.New("out of range")
	}
	if len(*a) == index {
		a.AddLast(val)
		return nil
	}
	*a = append(*a, val)
	copy((*a)[index+1:], (*a)[index:])
	return nil
}

func (a *IntArray)Remove(index int) (int, error) {
	if index < 0 || index >= len(*a) {
		return 0, errors.New("out of range")
	}
	if len(*a)-1 == index {
		return a.RemoveLast(), nil
	}
	res := (*a)[index]
	copy((*a)[index:], (*a)[index+1:])
	return res, nil
}

func (a *IntArray)Len() int {
	return len(*a)
}