package list

import "testing"

func TestSinglyLinkedListAppend(t *testing.T) {
	linkedList := NewSinglyLinkedList()
	expectedSize := 3

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestSinglyLinkedListPrepend(t *testing.T) {
	linkedList := NewSinglyLinkedList()
	expectedSize := 4
	expectedRoot := 0
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Prepend(expectedRoot)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	rootValue := linkedList.Root.Value
	if rootValue != expectedRoot {
		t.Fatalf("Root element should be %d but was %d", expectedRoot, rootValue)
	}
}

func TestSinglyLinkedListPrependWithEmptyValue(t *testing.T) {
	linkedList := NewSinglyLinkedList()
	expectedSize := 1
	expectedRoot := 5

	linkedList.Prepend(expectedRoot)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	rootValue := linkedList.Root.Value
	if rootValue != expectedRoot {
		t.Fatalf("Root element should be %d but was %d", expectedRoot, rootValue)
	}
}

func TestSinglyLinkedListGet(t *testing.T) {
	linkedList := NewSinglyLinkedList()
	expectedElement := 2
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	element := linkedList.Get(1)

	if element != expectedElement {
		t.Fatalf("Expected element should be %d but was %d", expectedElement, element)
	}
}

func TestSinglyLinkedListInsert(t *testing.T) {
	linkedList := NewSinglyLinkedList()
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

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestSinglyLinkedListInsertFront(t *testing.T) {
	linkedList := NewSinglyLinkedList()
	expectedElement := 4
	expectedSize := 1

	linkedList.Insert(expectedElement, 0)

	element := linkedList.Get(0)

	if element != expectedElement {
		t.Fatalf("Expected element at index 0 should be %d but was %d", expectedElement, element)
	}

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestSinglyLinkedListInsertBack(t *testing.T) {
	linkedList := NewSinglyLinkedList()
	original := 3
	expectedElement := 4
	expectedSize := 4
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(original)

	linkedList.Insert(expectedElement, 2)
	element := linkedList.Get(2)
	next := linkedList.Get(3)

	if element != expectedElement {
		t.Fatalf("Expected element at index 2 should be %d but was %d", expectedElement, element)
	}

	if next != original {
		t.Fatalf("Expected element at index 3 should be %d but was %d", expectedElement, next)
	}

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoubyLinkedListAppend(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	expectedSize := 3

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoublyLinkedListPrepend(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	expectedSize := 4
	expectedRoot := 0
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Prepend(expectedRoot)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	rootValue := linkedList.Head.Value
	if rootValue != expectedRoot {
		t.Fatalf("Root element should be %d but was %d", expectedRoot, rootValue)
	}
}

func TestDoublyLinkedListPrependWithEmptyValue(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	expectedSize := 1
	expectedRoot := 5

	linkedList.Prepend(expectedRoot)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	rootValue := linkedList.Head.Value
	if rootValue != expectedRoot {
		t.Fatalf("Root element should be %d but was %d", expectedRoot, rootValue)
	}
}

func TestDoublyLinkedListGet(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	expectedElement := 2
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	element := linkedList.Get(1)

	if element != expectedElement {
		t.Fatalf("Expected element should be %d but was %d", expectedElement, element)
	}
}

func TestDoublyLinkedListInsert(t *testing.T) {
	linkedList := NewDoublyLinkedList()
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

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoublyLinkedListInsertFront(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	expectedElement := 4
	expectedSize := 1

	linkedList.Insert(expectedElement, 0)

	element := linkedList.Get(0)

	if element != expectedElement {
		t.Fatalf("Expected element at index 0 should be %d but was %d", expectedElement, element)
	}

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoublyLinkedListInsertBack(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	original := 3
	expectedElement := 4
	expectedSize := 4
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(original)

	linkedList.Insert(expectedElement, 2)
	element := linkedList.Get(2)
	next := linkedList.Get(3)

	if element != expectedElement {
		t.Fatalf("Expected element at index 2 should be %d but was %d", expectedElement, element)
	}

	if next != original {
		t.Fatalf("Expected element at index 3 should be %d but was %d", expectedElement, next)
	}

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoublyLinkedListNodeBindings(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	elements := [5]int{10, 20, 30, 40, 50}

	linkedList.Append(20)
	linkedList.Append(40)
	linkedList.Insert(10, 0)
	linkedList.Append(50)
	linkedList.Insert(30, 2)

	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func nodeBindingsCorrect(list *DoublyLinkedList, elements [5]int) bool {
	current := list.Head
	if list.Head.Value != elements[0] {
		return false
	}
	if list.Tail.Value != elements[len(elements)-1] {
		return false
	}
	for i := 0; i < len(elements); i++ {
		if current.Value != elements[i] {
			return false
		}
		if i == 0 && current.Prev != nil {
			return false
		} else if i == len(elements)-1 && current.Next != nil {
			return false
		}
		if i > 0 && current.Prev.Value != elements[i-1] {
			return false
		}
		if i < len(elements)-1 && current.Next.Value != elements[i+1] {
			return false
		}
		current = current.Next
	}
	return true
}
