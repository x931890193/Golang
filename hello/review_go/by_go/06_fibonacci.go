/*
Create on: 2018-10-03 下午8:20
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

func main() {
	// fibnocci 数列 递归实现
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Println(result)
	}

}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}

}
