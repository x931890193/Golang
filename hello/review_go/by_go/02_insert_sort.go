package main

import "fmt"

func main() {
	//插入排序
	arr := []int{1, 2, 5, 8, 7, 9, 6, 4, 3, 0}
	fmt.Println(InsertSort(arr))
}

func InsertSort(arr []int) []int {
	num := len(arr)
	for i := 1; i < num; i++ {
		min := arr[i]
		j := i - 1
		for j >= 0 && arr[j] >= min {
			arr[j+1] = arr[j]
			j--
			arr[j+1] = min
		}
	}

	return arr
}
