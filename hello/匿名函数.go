package main //

import "fmt"

func main() {
	a := 10
	str := "jack"
	// 匿名函数， 没有名字， 函数定义， 不调用
	f1 := func() {
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}
	f1()
	// 给一个函数类型 起别名
	type FuncType func() // 定义一个函数类型 无参，无返回
	// 声明变量
	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数同时调用
	func() {
		fmt.Printf("a=%d, str=%s\n", a, str)
	}() // 此处的 括号表示调用

	func() {
		// 闭包以引用的方式捕获外部变量
		a = 1024
		str = "hello go"
		// 内部调用
		fmt.Printf("内部：a=%d, str=%s\n", a, str)
	}() // 直接调用
	fmt.Printf("外部：a=%d, str=%s\n", a, str)
}
