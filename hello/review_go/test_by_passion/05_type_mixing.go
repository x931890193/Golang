/*
Create on: 2018-10-03 下午8:54
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

func main() {
	var a int  // 默认inte16   还有 int8(byte 别名) int32 int64
	var b int32
	a = 15
	b = a + a // cannot use a + a (type int) as type int32 in assignment
	b = b + 5  //  int16 也不能够被隐式转换为 int32。
	fmt.Println(a,b)
}