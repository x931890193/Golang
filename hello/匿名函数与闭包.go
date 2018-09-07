package main

import "fmt"

// 所谓闭包就是一个函数 捕获了和它在一个作用域的其他常量和变量，
// 当闭包被调用的时候， 不管程序在什么地方调用，闭包都能使用这些量
// 延长了 变量， 常量的生命周期
// 不关心这些变量和常量是否超出了作用域， 只要闭包还在使用他们，这些变量就会存在

func test() int {
	// 变量的生命周期， 函数被调用时才分配空间， 初始化为零
	// 函数调用完成， 变量被销毁
	var x int // 初始值为0
	x++
	return x * x
}

// 函数的返回值是一个匿名函数
func demo() func() int {
	var i int = 20
	return func() int {
		i *= 2
		return i * i

	}
}

func main() {
	fmt.Println(test())
	// 返回值为一个函数
	f := demo()
	//	result := f()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
