package main // 单文件执行 包含main 包

import "fmt"

// go 函数 可以返回多个值

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	// a := 10
	// b := 20
	// c := 30

	a, b := 10, 20
	// 交换两个变量 传统方法
	fmt.Print("传统方法\n")
	fmt.Printf("交换前：a=%d, b=%d\n", a, b)
	var tmp int
	tmp = a
	a = b
	b = tmp

	fmt.Printf("交换后：a=%d, b=%d\n", a, b)

	c, d := 20, 30
	fmt.Print("go 和 python的方法\n")
	fmt.Printf("交换前：c=%d, d=%d\n", c, d)
	c, d = d, c
	fmt.Printf("交换后：c=%d, d=%d\n", c, d)
	i, j := 10, 20 // 多重赋值

	// 匿名变量 (_)， 丢弃数据不处理   _ 匿名变量  配合函数返回值使用才有优势
	tmp, _ = i, j
	fmt.Println("tmp=", tmp)
	var x, y, z int
	x, y, _ = test() // retrun 1, 2, 3
	fmt.Printf("x= %d, y=%d, z=%d\n", x, y, z)
	fmt.Printf("x= %d, y=%d\t%s\n", x, y, "用_ 匿名变量丢弃了z")
}
