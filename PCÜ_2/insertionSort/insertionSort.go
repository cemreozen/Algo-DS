package main
import ("fmt")

func insertionSort(array []int) {
	i := 1
	j := 0
	length := len(array) - 1
	for i <= length {
		element := array[i]
		j = i
		for j > 0 && array[j-1] > element {
			array[j] = array[j-1]
			j = j - 1
		}
		array[j] = element
		i = i + 1
	}
}

func main() {
	array := []int {1, 809, 32, 571}
	insertionSort(array)
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])
	fmt.Println(array[3])
}
