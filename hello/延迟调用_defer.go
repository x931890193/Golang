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
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("i got you")
		}
	}() // 捕获恐慌 哈哈哈
	// 延迟调用， 在主函数结束之前调用
	defer fmt.Println("bye world!") // 最后再调用

	// defer 的执行顺序  类似与栈  先写的后执行 从上到下 无论如何都会执行
	defer fmt.Println("11111111")
	defer fmt.Println("22222222")
	// 调用一个零除
	defer DivZero(0) // 产生panic 上边的recover 捕获程序不崩溃

	defer fmt.Println("33333333")

	// defer从下到上执行  最下的先执行 多个defer的执行顺序 不管程序有没有错s
}
