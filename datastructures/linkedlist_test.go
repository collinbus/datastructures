package datastructures

import (
	"testing"
)

func TestInitLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	list := NewLinkedList(items)

	current := list.Head
	for i := 0; i < len(items); i++ {
		if current != nil && current.Value != items[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
}

func TestInsertLinkedList(t *testing.T) {
	items := []int{1, 2, 4, 5}
	expectedItems := []int{1, 2, 3, 4, 5}
	list := NewLinkedList(items)

	list.Insert(3, 2)

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
}

func TestInsertFrontLinkedList(t *testing.T) {
	items := []int{1, 2, 4, 5}
	expectedItems := []int{3, 1, 2, 4, 5}
	list := NewLinkedList(items)

	list.Insert(3, 0)

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
}

func TestPushList(t *testing.T) {
	items := []int{1, 2, 4, 5}
	expectedItems := []int{1, 2, 4, 5, 6}
	list := NewLinkedList(items)

	list.Push(6)

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
}

func TestDeleteLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	list := NewLinkedList(items)
	expectedItems := []int{1, 2, 4, 5}

	list.RemoveAt(2)

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
}

func TestDeleteFirstItemLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	list := NewLinkedList(items)
	expectedItems := []int{2, 3, 4, 5}

	list.RemoveAt(0)

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
}

func TestPopLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	list := NewLinkedList(items)
	expectedItems := []int{1, 2, 3, 4}
	expectedPoppedValue := 5

	poppedValue := list.Pop()

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
	}
	if poppedValue != expectedPoppedValue {
		t.Fatalf("Expected popped value to be %d but was %d\n", expectedPoppedValue, poppedValue)
	}
}

func TestNextAndPrevValues(t *testing.T) {
	items := []int{3, 6, 9}
	list := NewLinkedList(items)

	current := list.Head
	if current.Prev != nil && current.Next.Value != 6 {
		t.Fatal("first node has invalid pointers")
	}
	current = current.Next
	if current.Prev.Value != 3 && current.Next.Value != 9 {
		t.Fatal("second node has invalid pointers")
	}
	current = current.Next
	if current.Prev.Value != 6 && current.Next != nil {
		t.Fatal("third node has invalid pointers")
	}
}

func TestNextAndPrevValuesAfterManipulation(t *testing.T) {
	items := []int{3, 6, 9}
	list := NewLinkedList(items)

	list.Insert(4, 1)
	list.RemoveAt(2)

	current := list.Head
	if current.Prev != nil && current.Next.Value != 4 {
		t.Fatal("first node has invalid pointers")
	}
	current = current.Next
	if current.Prev.Value != 3 && current.Next.Value != 9 {
		t.Fatal("second node has invalid pointers")
	}
	current = current.Next
	if current.Prev.Value != 4 && current.Next != nil {
		t.Fatal("third node has invalid pointers")
	}
}

func TestHeadAndTail(t *testing.T) {
	items := []int{3, 6, 9}
	
	list := NewLinkedList(items)
	list.Insert(1,0)
	list.RemoveAt(3)
	
	if list.Head.Value != 1 && list.Head.Prev != nil && list.Head.Next.Value != 3 {
		t.Fatal("Head is invalid")
	}
	if list.Tail.Value != 6 && list.Tail.Prev.Value != 6 && list.Tail.Next != nil {
		t.Fatal("Tail is invalid")
	}
}

