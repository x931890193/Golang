package main // 不要忘了  main 包

import "fmt"

//  数据类型的转换

func main() {

	var flag bool
	flag = true
	fmt.Printf("flag=%d\n", flag) // %d 匹配整形
	fmt.Printf("flag=%t\n", flag)

	fmt.Printf("flag = ")
}
