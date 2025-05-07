package main

import "fmt"

 func partition(array []int, low, high int) ([]int, int) {
	pivot := array[low]
	
	i := low

	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]
	return array, i
}

func quickSort(array []int, lo int, hi int) []int {
	if (hi > lo) {
		var p int
		array, p = partition(array, lo, hi)
		quickSort(array, lo, p-1)
		quickSort(array, p+1, hi)
	}
	return array
}

func main() {
	array := []int {3, 2, 1, 4, 9}

	fmt.Println(quickSort(array, 0, len(array) -1))
}
	
