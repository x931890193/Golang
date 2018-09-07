package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

//没有函数名字， FuncType 它是一个函数类型
type FuncType func(int, int) int

func main() {
	// 函数类型  函数指针
	// 函数也是一种数据类型, 可以使用type 起别名
	fmt.Println("a+b=", Add(1, 2))
	var fTest, fTest2 FuncType // 声明一个函数类型的变量， 变量名fTest
	fTest = Add                // 函数名赋值给变量 可以通过变量名调用函数
	var x, y int
	fmt.Print("请输入第一个数字:")
	fmt.Scan(&x)
	fmt.Print("请输入第二个数字:")
	fmt.Scan(&y)
	fmt.Printf("%d+%d=%d\n", x, y, fTest(x, y))

	fTest2 = Minus // 函数命名
	var i, j int
	fmt.Print("请输入第一个数字:")
	fmt.Scan(&i)
	fmt.Print("请输入第二个数字:")
	fmt.Scan(&j)
	fmt.Printf("%d+%d=%d\n", i, j, fTest2(i, j))

	// 变量调用 体现了一种多态的思想

}
