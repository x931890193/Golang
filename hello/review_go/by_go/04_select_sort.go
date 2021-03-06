package main

import (
	"C"
	"fmt"
)

func main() {
	arr := []int{1, 2, 5, 8, 7, 4, 3, 6, 9, 0}
	fmt.Println(SelectSort(arr))
}

//export SelectSort
func SelectSort(arr []int) []int {
	num := len(arr)
	for i := 0; i < num; i++ {
		fmt.Println(arr)
		min := i
		for j := i + 1; j < num; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
