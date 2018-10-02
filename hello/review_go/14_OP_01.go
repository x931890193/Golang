package main

// OOP编程 method
//A method is a function with an implicit first argument, called a receiver
// 语法   func(r, ReceviverType)funcname(paraments)(result)
import (
	"fmt"
	"math"
)

// 声明一个结构体
type A struct {
	width, height float32
}

// 声明一个结构体
type B struct {
	r float64
}

// 用python的语法来说 实例化一个对象, 绑定area方法
// 定义area函数 无参数 返回值类型为float32
func (a A) area() float32 {
	return a.width * a.height
}

func (b B) area() float64 {
	return b.r * b.r * math.Pi
}

func main() {

	s1 := A{10, 20}
	s2 := B{1}
	fmt.Println(s1.area())
	fmt.Println(s2.area())

}
