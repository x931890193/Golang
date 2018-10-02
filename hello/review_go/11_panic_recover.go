package main

import (
	"fmt"
	"os"
)

// 恐慌
// Panic 是一个内建函数, 使终端原有的控制流程, 进入一个令人恐慌的流程
// 当函数调用 Panic , 函数的执行被中断, 但是函数的延迟函数会执行
// panic  类型python中的异常 except  需要处理, 只能在defer中处理  使用reover()
// panic  抛异常, 在defer中 用recover捕获
func main() {
	fmt.Println("this is show how to use panic")
	my_panic()
	ThorwPanic(my_panic)
}

func my_panic() {
	// 我的恐慌!
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(" i got you")
		}
	}()
	var user = os.Getenv("USER")
	if user != "root" {
		panic("tht user envirnoment not set")
	}

}

func ThorwPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Exception has been caught.")
			b = true
		}
	}()
	f()
	panic(1)
	return b
}
