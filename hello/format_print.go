package main // 需要一个main包

import "fmt"

func main() {
	//格式化输出

	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	fmt.Printf("%T,%T,%T,%T\ns", a, b, c, d)           // 查看类型
	fmt.Printf("a=%d, b=%s, c=%c, d=%f\n", a, b, c, d) // 格式化输出
	fmt.Printf("a=%v, b=%v, c=%v, d=%v\n", a, b, c, d) // 万能输出

}
