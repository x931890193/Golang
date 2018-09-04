package main // 导入main包 单文件

import "fmt"

func main() {
	// 不同类型变量的声明（定义）
	//	var a int
	//	var b float64
	var (
		a int
		b float64
	)

	a, b = 10, 3.14

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	const (
		c int     = 30
		d float64 = 3.1415926
	)

	const (
		e = 30 // 常量定义自动推导类型
		f = 3.1415926
	)

	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("c=", e)
	fmt.Println("d=", d)

}
