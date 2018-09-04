package main //必须有一个main包 程序入口 有且只有一个
//  一个文件夹只能有一个main包  一个工程只能有一个主函数
import "fmt" // fmt.Printn   导入包 以后必须要使用
// 变量声明也必须用

func main() {
	// 变量声明：   赋值前必须声明
	// 1. var 变量名 类型 声明变量 必须要使用
	// 2. 只是声明没有初始化的变量， 默认值为0
	// 3. 同一个{}里面， 声明的变量必须要唯一
	// 4. 可以同时声明多个变量  var a, b int

	var a int
	a = 10 // 变量赋值
	fmt.Println(a)
	var b int = 10 // 初始化 声明 赋值
	fmt.Println(b)
	// 5. 自动推导数据类型， 必须初始化，通过初始化的值确定类型
	c := 30
	// %T 打印变量的所属类型
	fmt.Printf("c: type is %T\n", c)
	fmt.Println("Hello, Go!", a, b, c)
}
