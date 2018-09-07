package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// 延迟调用， 在主函数结束之前调用
	defer fmt.Println("bye world!") // 最后再调用

}
