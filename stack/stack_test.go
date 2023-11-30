package stack

import "testing"

func TestPush(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	expectedTop := 3

	result := stack.Peek()

	if result != expectedTop {
		t.Fatalf("expected top should be %d but was %d", expectedTop, result)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	poppedElement := stack.Pop()
	expectedTop := 2
	expectedSize := 2
	result := stack.Peek()

	if result != expectedTop {
		t.Fatalf("expected top should be %d but was %d", expectedTop, result)
	}
	if poppedElement != 3 {
		t.Fatalf("expected popped element should be %d but was %d", 3, poppedElement)
	}
	size := stack.Size()
	if size != expectedSize {
		t.Fatalf("expected size %d but was %d", expectedSize, size)
	}
}

func TestIsEmpty(t *testing.T) {
	stack := NewStack[int]()

	isEmpty := stack.IsEmpty()

	if !isEmpty {
		t.Fatal("stack is not empty but should be empty")
	}
}
