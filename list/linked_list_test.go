package list

import "testing"

func TestLinkedListAppend(t *testing.T) {
	linkedList := NewLinkedList()
	expectedSize := 3

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	if linkedList.size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.size)
	}
}

func TestLinkedListPrepend(t *testing.T) {
	linkedList := NewLinkedList()
	expectedSize := 4
	expectedRoot := 0
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Prepend(expectedRoot)

	if linkedList.size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.size)
	}
	rootValue := linkedList.root.value
	if rootValue != expectedRoot {
		t.Fatalf("Root element should be %d but was %d", expectedRoot, rootValue)
	}
}

func TestLinkedListPrependWithEmptyValue(t *testing.T) {
	linkedList := NewLinkedList()
	expectedSize := 1
	expectedRoot := 5

	linkedList.Prepend(expectedRoot)

	if linkedList.size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.size)
	}
	rootValue := linkedList.root.value
	if rootValue != expectedRoot {
		t.Fatalf("Root element should be %d but was %d", expectedRoot, rootValue)
	}
}

func TestLinkedListGet(t *testing.T) {
	linkedList := NewLinkedList()
	expectedElement := 2
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	element := linkedList.Get(1)

	if element != expectedElement {
		t.Fatalf("Expected element should be %d but was %d", expectedElement, element)
	}
}

func TestLinkedListInsert(t *testing.T) {
	linkedList := NewLinkedList()
	original := 2
	expectedElement := 4
	expectedSize := 4
	linkedList.Append(1)
	linkedList.Append(original)
	linkedList.Append(3)

	linkedList.Insert(expectedElement, 1)
	element := linkedList.Get(1)
	next := linkedList.Get(2)

	if element != expectedElement {
		t.Fatalf("Expected element at index 1 should be %d but was %d", expectedElement, element)
	}

	if next != original {
		t.Fatalf("Expected element at index 2 should be %d but was %d", expectedElement, next)
	}

	if linkedList.size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.size)
	}
}

func TestLinkedListInsertFront(t *testing.T) {
	linkedList := NewLinkedList()
	expectedElement := 4
	expectedSize := 1

	linkedList.Insert(expectedElement, 0)

	element := linkedList.Get(0)

	if element != expectedElement {
		t.Fatalf("Expected element at index 0 should be %d but was %d", expectedElement, element)
	}

	if linkedList.size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.size)
	}
}
