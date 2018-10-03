package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(BinarySearch(arr, 5))
}

func BinarySearch(arr []int, key int) int {
	low := 0
	high := len(arr)-1
	for arr[low] <= arr[high] {
		if key > low {
			low = key
		} else if key < high {
			high = key
		} else {
			return key
		}
	}
	return -1

}
