package datastructures

import "testing"

func TestPushOnStack(t *testing.T) {
	capacity := 5
	stack := NewStaticStack(capacity)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expectedValues := []int{1,2,3,0,0}
	for i := 0; i < capacity; i++ {
		if stack.values[i] != expectedValues[i] {
			t.Fatalf("element at position %d is not matching\n", i)
		}
	}
}
