package utils

func BinarySearchInts(array []int, element int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (right + left) / 2
		if array[middle] == element {
			return middle
		} else if array[middle] < element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
