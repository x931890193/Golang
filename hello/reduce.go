package main

import "fmt"

// 递归函数的使用
func MyReduce(a int) {
	// 函数调用自身
	fmt.Println("a=", a)

	if a <= 0 {
		fmt.Print("函数终止！")
		return
	}
	MyReduce(a - 1)

}

func main() {
	MyReduce(1000)
}
