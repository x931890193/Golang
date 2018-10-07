package main

import (
	"C"
	"fmt"
)

//export BubbleSort
func BubbleSort(arr []int) []int {
	// 改进的冒泡排序
	num := len(arr) //:=自动匹配变量类型
	for i := 0; i < num; i++ {
		status := false
		for j:=0; j < num-i-1; j++ {
			if arr[j] > arr[j+1] {
				status = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
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
