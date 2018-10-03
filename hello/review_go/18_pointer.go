/*
Create on: 2018-10-03 下午9:47
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var il = 5
	fmt.Printf("An integer: %d, it is location in memory: %p\n", il, &il)
	var intp *int // 指针定义  默认为nil  一个指针变量通常缩写为 ptr.
	intp = &il
	fmt.Printf("The memeory at %p is An integer: %d\n", intp, il)
	StringPointer()
	VoidPointer()
	UseRuntime()
}

func StringPointer() {
	s := "good bye"  // &s 取出地址
	var p *string = &s     // 定义一个指针变量 字符串类型
	*p = "hello world" // 去指针变量的值 然后修改
	fmt.Println(*p)
	fmt.Println(s)
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}

func VoidPointer(){
	var a int = 10
	var p *int=nil // 定义一个空指针
	//*p = 6666 // 空指针无法反向引用
	p = &a
	*p = 60 // 修改a原始地址的值
	fmt.Println(p, a)
}

func UseRuntime(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Version())
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU(), "CPU个数")
	fmt.Println(runtime.MemProfileRate)
}