/*
Create on: 2018-10-03 下午10:25
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
)

func main() {
	// 不不定长参数的使用
	x := min(1, 3, 2, 5) //调用函数 求最小值
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1} // 声明一个slice 并且赋初始值
	x = min(arr...)             // 函数需要一个不定长参数, 传递必须 arr...
	fmt.Printf("The minimum in the array arr is: %d\n", x)
	MyTest(arr...) // 传递一个不定长参数
}
func min(a ...int) int { // 不定长参数
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func MyTest(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
	T2(args...)
}

func T2(args ...int) {
	fmt.Println(args)
}
