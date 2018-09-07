package main

import (
	"fmt"
	"sub"
)

func main() {
	test() // 调用test.go 下边的test函数

	result := sub.Add(12, 20) // 在sub文件夹下
	// 调用不同文件夹下的包的函数   包名. 私有方法   首字母大写  共有方法
	fmt.Println(result)
}

// 分文件编程， 必须放在src目录下
// 设置 GOPATH环境变量，
// 同一个目录的 package 必须一致
// go env 查看go的环境变量
// 同一个目录， 调用别的文件函数， 直接调用， 无需包名引用
