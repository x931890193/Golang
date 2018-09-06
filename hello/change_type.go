package main // 不要忘了  main 包

import "fmt"

//  数据类型的转换

func main() {
	// 不能转换的类型，叫做不兼容类型
	var flag bool
	flag = true
	//	fmt.Printf("flag=%d\n", flag) // %d 匹配整形
	fmt.Printf("flag=%t\n", flag)
	// bool类型不能转换为Int
	// 整形也不能转换为bool
	// flag=bool(1)
	fmt.Println("flag = ", flag)

	var ch byte
	ch = 'a' // 字符类型本质就是整形
	var t int
	t = int(ch)
	fmt.Println("t=", t)

}
