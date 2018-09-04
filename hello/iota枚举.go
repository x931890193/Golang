package main // 必须的main包  单文件程序

import "fmt"

func main() {
	// iota 常量自动生成器  每隔一行， 自动累加1
	// iota 给常量赋值使用
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2

	)
	fmt.Printf("a=%d, b =%d, c=%d\n", a, b, c)
	// iota 遇到const 重置为0
	const d = iota
	fmt.Printf("d=%d\n", d)

	const (
		a1 = iota // 0  定义常量iota时  可以只写一个iota
		a2
		a3
		a4
	)
	fmt.Printf("a1=%d, a2=%d, a3=%d, a4=%d\n", a1, a2, a3, a4)

	// iota 累加 如果是同一行， 值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)

}
