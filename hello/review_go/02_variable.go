package main

import (
	"fmt"
)

// 所有声明的变量,import的包都需要使用, 不然报错

// 类型为int 4个字节, 未初始化
//var a, b, c int //初始值为0
//多种定义方法
var (
	a int     = 20
	b float32 = 3.2
	c string  = "go go go"
)

// 定义变量并初始化
var d float32 = 3.14

var e, f, g string = "hello", "world", "byebye"

var x, y, z = 1, 2.0, 3 // 定义多个变量, 同时赋初始值

func main() {
	// 简短声明。不过它有一个限制，它只能用在函数内部
	i, j, k := "hello", 2, 3.14 // 自动推导类型
	_, o := 1, "world"          // "_" 丢弃该变量
	fmt.Println(a, b, c, d, e, f, g, i, j, k, x, y, z, o)

}
