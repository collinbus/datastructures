package utils

func BinarySearchInts(array []int, element int) int {
	left := 0
	right := len(array) - 1
	for left < right {
		middle := left + ((right - left) / 2)
		if array[middle] == element {
			return middle
		} else if array[middle] > element {
			right = middle
		} else {
			left = middle
		}
	}
	return -1
}
