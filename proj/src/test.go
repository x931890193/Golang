package main

import "fmt"

func init() {
	fmt.Println("test is initial------") // 导包的时候自动执行init函数 初始化
}

func test() {
	fmt.Println("this a test")
}
