package list

import "testing"

func TestSinglyLinkedListAppend(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	expectedSize := 3

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestSinglyLinkedListPrepend(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
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
	linkedList := NewSinglyLinkedList[int]()
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
	linkedList := NewSinglyLinkedList[int]()
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
	linkedList := NewSinglyLinkedList[int]()
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
	linkedList := NewSinglyLinkedList[int]()
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
	linkedList := NewSinglyLinkedList[int]()
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

func TestSinglyLinkedListRemoveMiddle(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4

	linkedList.RemoveAt(2)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	element := linkedList.Get(0)
	if element != 1 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 1, element)
	}
	element = linkedList.Get(1)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(2)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	element = linkedList.Get(3)
	if element != 5 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 5, element)
	}
}

func TestSinglyLinkedListPoll(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4
	expectedElement := 1

	element := linkedList.Poll()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
	element = linkedList.Get(0)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(1)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(2)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	element = linkedList.Get(3)
	if element != 5 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 5, element)
	}
}

func TestSinglyLinkedListRemoveFirst(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4

	linkedList.RemoveAt(0)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}

	element := linkedList.Get(0)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(1)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(2)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	element = linkedList.Get(3)
	if element != 5 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 5, element)
	}
}

func TestSinglyLinkedListPop(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4
	expectedElement := 5

	element := linkedList.Pop()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
	element = linkedList.Get(0)
	if element != 1 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 1, element)
	}
	element = linkedList.Get(1)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(2)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(3)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
}

func TestSinglyLinkedListPopSingle(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	expectedSize := 0
	expectedElement := 1

	element := linkedList.Pop()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
}

func TestSinglyLinkedListPollSingle(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	expectedSize := 0
	expectedElement := 1

	element := linkedList.Pop()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
}

func TestSinglyLinkedListRemoveLast(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4

	linkedList.RemoveAt(4)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}

	element := linkedList.Get(0)
	if element != 1 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 1, element)
	}
	element = linkedList.Get(1)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(2)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(3)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
}

func TestDoubyLinkedListAppend(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	expectedSize := 3

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoublyLinkedListPrepend(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
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
	linkedList := NewDoublyLinkedList[int]()
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
	linkedList := NewDoublyLinkedList[int]()
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
	linkedList := NewDoublyLinkedList[int]()
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
	linkedList := NewDoublyLinkedList[int]()
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
	linkedList := NewDoublyLinkedList[int]()
	original := 3
	expectedElement := 4
	expectedSize := 4
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(original)

	linkedList.Insert(expectedElement, 2)
	element := linkedList.Get(2)
	next := linkedList.Get(3)

	if element != original {
		t.Fatalf("Expected element at index 2 should be %d but was %d", expectedElement, element)
	}

	if next != expectedElement {
		t.Fatalf("Expected element at index 3 should be %d but was %d", expectedElement, next)
	}

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestSinglyLinkedListPrependingMultipleItems(t *testing.T) {
	linkedList := NewSinglyLinkedList[int]()
	expectedSize := 3

	linkedList.Prepend(1)
	linkedList.Prepend(2)
	linkedList.Prepend(3)

	next := linkedList.Get(0)
	if next != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, next)
	}
	next = linkedList.Get(1)
	if next != 2 {
		t.Fatalf("Expected element at index 1 should be %d but was %d", 2, next)
	}
	next = linkedList.Get(2)
	if next != 1 {
		t.Fatalf("Expected element at index 2 should be %d but was %d", 1, next)
	}
	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
}

func TestDoublyLinkedListNodeBindings(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	elements := []int{10, 20, 30, 40, 50, 60}

	linkedList.Append(20)
	linkedList.Append(40)
	linkedList.Insert(10, 0)
	linkedList.Append(50)
	linkedList.Insert(30, 2)
	linkedList.Insert(60, 4)

	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestDoublyLinkedListRemoveMiddle(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	elements := []int{1, 2, 4, 5}
	expectedSize := 4

	linkedList.RemoveAt(2)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	element := linkedList.Get(0)
	if element != 1 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 1, element)
	}
	element = linkedList.Get(1)
	if element != 2 {
		t.Fatalf("Expected element at index 1 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(2)
	if element != 4 {
		t.Fatalf("Expected element at index 2 should be %d but was %d", 4, element)
	}
	element = linkedList.Get(3)
	if element != 5 {
		t.Fatalf("Expected element at index 3 should be %d but was %d", 5, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestDoublyLinkedListPoll(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4
	expectedElement := 1
	elements := []int{2, 3, 4, 5}

	element := linkedList.Poll()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
	element = linkedList.Get(0)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(1)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(2)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	element = linkedList.Get(3)
	if element != 5 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 5, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestDoublyLinkedListRemoveFirst(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4
	elements := []int{2, 3, 4, 5}

	linkedList.RemoveAt(0)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}

	element := linkedList.Get(0)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(1)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(2)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	element = linkedList.Get(3)
	if element != 5 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 5, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestDoublyLinkedListPop(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4
	expectedElement := 5
	elements := []int{1, 2, 3, 4}

	element := linkedList.Pop()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
	element = linkedList.Get(0)
	if element != 1 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 1, element)
	}
	element = linkedList.Get(1)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(2)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(3)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestDoublyLinkedListPopSingle(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	expectedSize := 0
	expectedElement := 1
	elements := []int{}

	element := linkedList.Pop()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestDoublyLinkedListPollSingle(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	expectedSize := 0
	expectedElement := 1
	elements := []int{}

	element := linkedList.Pop()

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}
	if element != expectedElement {
		t.Fatalf("Returned element should be %d but was %d", expectedElement, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func TestSinglyDoublyListRemoveLast(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	expectedSize := 4
	elements := []int{1, 2, 3, 4}

	linkedList.RemoveAt(4)

	if linkedList.Size != expectedSize {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.Size)
	}

	element := linkedList.Get(0)
	if element != 1 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 1, element)
	}
	element = linkedList.Get(1)
	if element != 2 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 2, element)
	}
	element = linkedList.Get(2)
	if element != 3 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 3, element)
	}
	element = linkedList.Get(3)
	if element != 4 {
		t.Fatalf("Expected element at index 0 should be %d but was %d", 4, element)
	}
	if !nodeBindingsCorrect(linkedList, elements) {
		t.Fatal("The bindings of the linked list are wrongly configured")
	}
}

func nodeBindingsCorrect(list *DoublyLinkedList[int], elements []int) bool {
	if len(elements) == 0 {
		return list.Head == nil && list.Tail == nil
	}
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
