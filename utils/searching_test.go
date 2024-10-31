package utils

import "testing"

func TestBinarySearchInts(t *testing.T) {
	array := []int{1, 2, 3, 5, 7, 10, 12, 15, 21}
	elementToFind := 12

	index := BinarySearchInts(array, elementToFind)

	expectedIndex := 6
	if index != expectedIndex {
		t.Fatalf("Index of element '%d' should be %d but was: %d\n", elementToFind, expectedIndex, index)
	}
}

func TestBinarySearchIntsEven(t *testing.T) {
	array := []int{1, 2, 3, 5, 7, 10, 12, 15}
	elementToFind := 10

	index := BinarySearchInts(array, elementToFind)

	expectedIndex := 5
	if index != expectedIndex {
		t.Fatalf("Index of element '%d' should be %d but was: %d\n", elementToFind, expectedIndex, index)
	}
}
