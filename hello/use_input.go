package main // main 包

import "fmt" // 导入 fmt 包

func main() { // 入口

	var a int // 声明变量
	fmt.Printf("请输入变量a:")
	// 阻塞等待用户输入
	fmt.Scanf("%d", &a) // 加 &
	fmt.Println("a=", a)
	fmt.Printf("请再次输入变量a:")
	fmt.Scan(&a) // 自动匹配格式
	fmt.Println("a=", a)
}
