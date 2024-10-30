package utils

import (
	"fmt"
	"sort"
	"testing"
)

func TestInsertionSortInts(t *testing.T) {
	array := []int{46, 28, 53, 1, 15, 97, 34, 28}

	InsertionSortInts(array)
	isSorted := sort.SliceIsSorted(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	if !isSorted {
		fmt.Println(array)
		t.Fatal("Array is not sorted")
	}
}

func TestInsertionSortStrings(t *testing.T) {
	array := []string{"whut", "thut", "can", "i", "you", "feel", "ah", "lol"}

	InsertionSortStrings(array)
	isSorted := sort.SliceIsSorted(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	if !isSorted {
		fmt.Println(array)
		t.Fatal("Array is not sorted")
	}
}
