package main

import "fmt"

var x int = 10

func test() {
	a := 10
	fmt.Println(a)

}

func demo() (result int) {

	fmt.Println("函数内部，使用全局变量", x)
	result = x
	return x
}

func main() {
	// 定义在{}中的变量称为局部变量  {} 内部
	// 局部变量的作用域只在大括号内有效
	// 执行到定义变量的那句话，才分配空间，离开作用域自动销毁
	//	fmt.Println(a) //undefined: a

	if status := 3; status == 3 {
		fmt.Println("status=", status)

	}
	//	status = 4 //undefined: status

	// 全局变量 定义在函数外部的变量
	fmt.Println("a=", x)

	// 不同作用域变量名同名的处理
	// 使用变量的原则 就近使用
	var x float64 = 3.14 // 局部变量
	fmt.Println("x=", x)
	fmt.Printf("x type is %T\n", x)

	{ // 局部变量
		var x byte = 'a'
		fmt.Printf("x=%c\n", x)
		fmt.Printf("x type is %T\n", x)
	}
	fmt.Println("x=", demo())
}
