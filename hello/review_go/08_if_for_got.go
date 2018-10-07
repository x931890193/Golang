package main

import (
	"fmt"
)

func main() {

	func_goto_test()
	func_if_test()
	func_for_test()
}

func func_goto_test() {
	i := 0
here:
	fmt.Println(i)
	i++
	if i >= 1000 {
		return
	}
	goto here
	// 死循环
}

func func_if_test() {
	var a int
	fmt.Println("请输入:")
	fmt.Scan(&a)
	if a > 10 {
		fmt.Println("a is greater than 10")
	} else {
		fmt.Println("a is less than 10")
	}
}

func func_for_test() {
	var sum int
	for i := 0; i < 1000; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)
	for sum < 999999 {
		sum += sum
	}
	fmt.Println(sum)
	for index := 0; index < 1000; index++ {
		if index == 500 {
			break
		}
		fmt.Println(index)
	}
}
