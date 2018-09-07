package main

import "fmt"

type FuncType func(int, int) int

//回调函数， 函数的一个参数是一个函数类型， 这个函数就是回调函数
// 计算器， 四则运算
// 多态， 多种形态， 调用同一个接口，可以有不同的表现,四则运算
// 先有想法， 后面再实现功能
//var fTest FuncType

func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Printf("Calc\t")
	result = fTest(a, b) // fTest 函数
	return result
}

func main() {
	var a, b int
	var c string
	for {
		fmt.Printf("数字1：")
		fmt.Scan(&a)

		fmt.Printf("操作方法:")
		fmt.Scanf("%s", &c)

		fmt.Printf("数字2：")
		fmt.Scan(&b)
		//		fmt.Printf("%s", c)
		if c == "+" {
			result := Calc(a, b, Add)
			fmt.Println(result)

		} else if c == "-" {
			result := Calc(a, b, Minus)
			fmt.Println(result)

		} else if c == "*" {
			result := Calc(a, b, Mult)
			fmt.Println(result)

		} else if c == "/" {
			result := Calc(a, b, Div)
			fmt.Println(result)

		} else {
			fmt.Println("输入有误！")

		}

	}

}
