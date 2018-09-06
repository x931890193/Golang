package main

import "fmt"

// if 语句
func main() {

	s := "hello world!"

	if s == "hello world!" {
		// 条件判断
		fmt.Println("hello world")
	}
	// if 支持1个初始化语句, 初始化语句和判断条件以分号分割
	if a := 10; a == 10 { // 前面为初始化语句， 后边为判断语句
		// 条件为真执行 单分支
		fmt.Println("a=10")
	}
	if b := 120; b == 10 {
		fmt.Println("b==10")
	} else {
		fmt.Println("a!=10")

	}
	var c int
	fmt.Scan(&c)
	fmt.Printf("c' type is %T\n", c)
	if c == 10 {
		fmt.Println("6666")
	} else if c == 20 {
		fmt.Println("20")
	} else {
		fmt.Println("no")
	}

}
