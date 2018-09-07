package main

import "fmt"

func DivZero(x int) {
	a := 10
	fmt.Println(x)
	b := a / x
	fmt.Println(b)
}

func main() {
	fmt.Println("hello world")

	// 延迟调用， 在主函数结束之前调用
	defer fmt.Println("bye world!") // 最后再调用

	// defer 的执行顺序  类似与栈  先写的后执行 从上到下 无论如何都会执行
	defer fmt.Println("11111111")
	defer fmt.Println("22222222")
	// 调用一个零除
	defer DivZero(0)

	defer fmt.Println("33333333")

	// defer从下到上执行  最下的先执行 多个defer的执行顺序 不管程序有没有错s
}
