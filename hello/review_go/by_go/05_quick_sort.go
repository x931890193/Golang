package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,5,8,7,4,3,6,9,0,12,13,45,78,89,56,23,11,12,23,56,89,79,46,13,00,11,22,11,22,33,66,88,77,44,44,11,10,26}
	//fmt.Println(len(arr))
	fmt.Println(QuickSort(arr),len(QuickSort(arr)))
}

func QuickSort(arr []int)[]int{
	if len(arr) ==0 {
		return []int{}
	}
	qmiddle := arr[0]
	qleft := QuickSort(QuickLeft(arr, qmiddle))
	qright := QuickSort(QuickRight(arr,qmiddle))
	return append(append(qleft,qmiddle),qright...) //list = append(slice, data)   list = append(slice1, slice2...) !
	}

func QuickLeft(arr []int, qmiddle int)[]int{
	leftarr := []int{}
	num := len(arr)
	for i:= 1; i < num; i++{
		if arr[i] <= qmiddle{
			leftarr = append(leftarr,arr[i])
		}
	}
	return leftarr
}

func QuickRight(arr []int, qmiddle int)[]int{
	rightarr := []int{}
	num := len(arr)
	for j:= 1; j < num; j++{
		if arr[j] > qmiddle{
			rightarr = append(rightarr,arr[j])
		}
	}
	return rightarr
}