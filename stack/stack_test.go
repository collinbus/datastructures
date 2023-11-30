package stack

import "testing"

func TestPush(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	expectedTop := 3

	result := stack.Peek()

	if result != expectedTop {
		t.Fatalf("expected top should be %d but was %d", expectedTop, result)
	}
}
