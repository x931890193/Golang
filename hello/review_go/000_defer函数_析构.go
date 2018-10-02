package main // 必须分main包

import "fmt"
// defer 从下开始往上开始执行  程序执行从上往下执行参数传递
func main() {
	a := 10
	b := 20
	// defer 最后执行
	defer func(a, b int) {
				a = 11
				b = 22
		fmt.Println(1111111)
		fmt.Println(a, b)
	}(a, b) //调用  此处优先传递参数， 从上到下开始执行 此处先传递参数
	a = 111
	b = 222
	fmt.Println("第一次", a, b)
	defer func(a, b int) {
		a = 22
		b = 33
		fmt.Println(2222222)
		fmt.Println(a, b)
	}(a, b)


}
