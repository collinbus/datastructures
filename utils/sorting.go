package utils

func InsertionSortInts(array []int) {
	length := len(array)
	for i := 1; i < length; i++ {
		current := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > current; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = current
	}
}

func InsertionSortStrings(array []string) {
	length := len(array)
	for i := 1; i < length; i++ {
		current := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > current; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = current
	}
}
