package main

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

func TestInsertLastLinkedList(t *testing.T) {
	items := []int{1, 2, 4, 5}
	expectedItems := []int{1, 2, 4, 5, 6}
	list := NewLinkedList(items)

	list.Insert(6, 4)

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

func TestDeleteLastItemLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	list := NewLinkedList(items)
	expectedItems := []int{1, 2, 3, 4}

	list.RemoveAt(4)

	current := list.Head
	for i := 0; i < len(expectedItems); i++ {
		if current != nil && current.Value != expectedItems[i] {
			t.Fatalf("Item at index %d is not matching\n", i)
		}
		current = current.Next
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
