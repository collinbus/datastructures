package datastructures

import (
	"testing"
	"fmt"
)

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

func TestPopOnStack(t *testing.T) {
	capacity := 5
	stack := NewStaticStack(capacity)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	poppedValue := stack.Pop()

	
	expectedPoppedValue := 3
	expectedValues := []int{1,2,0,0,0}
	
	fmt.Println(stack.values)
	for i := 0; i < capacity; i++ {
		if stack.values[i] != expectedValues[i] {
			t.Fatalf("element at position %d is not matching\n", i)
		}
	}
	if expectedPoppedValue != poppedValue {
		t.Fatalf("Expected popped value %d but got %d\n", expectedPoppedValue, poppedValue)
	}
}
