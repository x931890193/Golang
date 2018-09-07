package main // 必须分main包

import "fmt"

func main() {
	a := 10
	b := 20
	// defer 最后执行
	defer func(a, b int) { //匿名函数
		//		a = 10
		//		b = 30
		fmt.Println(a, b)
	}(a, b) //调用  此处优先传递参数，
	a = 111
	b = 222
	fmt.Println("第一次", a, b)
}
