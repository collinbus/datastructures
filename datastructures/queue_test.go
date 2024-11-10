package datastructures

import "testing"

func TestEnqueue3Items(t *testing.T) {
	queue := NewQueue()
	expectedQ := []int{ 1, 2, 3 }
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	current := queue.front
	for i := 0 ; i < len(expectedQ) ; i++ {	
		if current.Value != expectedQ[i] {
			t.Fatalf("Item at %d is %d, should be %d\n", i, current.Value, expectedQ[i])
		}
		current = current.Next
	}
}

func TestEnqueue5ItemsDequeue3(t *testing.T) {
	queue := NewQueue()
	expectedQ := []int{ 4, 5 }
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	current := queue.front
	for i := 0 ; i < len(expectedQ) ; i++ {	
		if current.Value != expectedQ[i] {
			t.Fatalf("Item at %d is %d, should be %d\n", i, current.Value, expectedQ[i])
		}
		current = current.Next
	}
}
