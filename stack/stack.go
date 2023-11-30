package stack

import "datastructures/list"

type Stack[T any] interface {
	Push(element T)
	Peek() T
	Pop() T
	Size() int
	IsEmpty() bool
}

type DynamicStack[T any] struct {
	elements *list.SinglyLinkedList[T]
}

func NewStack[T any]() *DynamicStack[T] {
	return &DynamicStack[T]{
		elements: list.NewSinglyLinkedList[T](),
	}
}

func (s *DynamicStack[T]) Push(element T) {
	s.elements.Append(element)
}

func (s *DynamicStack[T]) Peek() T {
	return s.elements.Get(s.elements.Size - 1)
}

func (s *DynamicStack[T]) Pop() T {
	return s.elements.Pop()
}

func (s *DynamicStack[T]) Size() int {
	return s.elements.Size
}

func (s *DynamicStack[T]) IsEmpty() bool {
	return s.elements.Size == 0
}
