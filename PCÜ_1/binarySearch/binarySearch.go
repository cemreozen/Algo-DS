package main

func binary_search(data []int, value int) bool {
	var array_length int = len(data) - 1
	var starting_index int = 0
	var middle int = (starting_index + array_length) / 2
	for starting_index <= array_length {
		 if data[middle] == value {
			 return true
		 }
		 if data[middle] > value {
			 array_length = middle -1
		 }
		 if data[middle] < value {
			 starting_index = middle + 1
		 }
	 }
	 return false
 }
