package main

import (
	"fmt"
)

func BubbleSort(arr []int) []int {
	// 改进的冒泡排序
	num := len(arr) //:=自动匹配变量类型
	for i := 0; i < num; i++ {
		status := false
		for j := i + 1; j < num; j++ {
			if arr[i] > arr[j] {
				status = true
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		if status == false {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{100, 200, 55, 9, 88, 77, 66, 55, 44, 1, 2, 3, 65, 4, 7, 89, 6, 3, 32, 1, 4, 0, 5, 8, 7}
	fmt.Printf("排序前:%v\n", arr)
	NewArr := BubbleSort(arr)
	fmt.Printf("排序后:%v", NewArr)
}
