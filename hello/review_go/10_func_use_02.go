package main

import (
	"fmt"
)

// type 关键字 类型名  函数(参数) 返回值
type testInt func(int) bool // 声明一个函数类型

func main() {
	my_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("my_slice=", my_slice)
	odd := filter(my_slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(my_slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func isOdd(a int) bool {
	if a%2 == 0 {
		return false
	}
	return true
}

func filter(my_slice []int, f testInt) []int {
	// 参数1 my_slice 类型 []int
	// 参数2 f        类型 testInt 函数类型
	// 返回值  类型 []int
	var result []int // 声明一个变量  类型 []int
	for _, value := range my_slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result

}
