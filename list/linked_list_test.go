package list

import "testing"

func TestLinkedListAppend(t *testing.T) {
	linkedList := NewLinkedList()
	expectedSize := 3

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	if linkedList.size != 3 {
		t.Fatalf("Size should be %d but was %d", expectedSize, linkedList.size)
	}
}
