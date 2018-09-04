// 常量
package main //  导入main包

import "fmt"

func main() {
	// 变量: 程序运行中， 可以改变  var 声明
	// 常量： 程序运行中， 不可以改变  const 声明
	const a int = 10
	//	a = 20 // 尝试修改a的内容  报错-->>>  cannot assign to a
	// 常量不允许修改
	fmt.Println("a=", a)
	const b = 10.235656 // 自动推导类型   不用 ：=   报错
	fmt.Printf("b: type is %T\n", b)
	fmt.Println("b=", b)
}
