package datastructures

type Stack interface {
	Push(element int)
}

type StaticStack struct {
	values []int
	size   int
}

func NewStaticStack(capacity int) *StaticStack {
	values := make([]int, capacity)
	return &StaticStack{values: values}
}

func (s *StaticStack) Push(element int) {
	s.values[s.size] = element
	s.size++
}
