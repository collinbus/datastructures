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
