/*
Create on: 2018-10-04 下午7:15
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var f uint8
	a := 2
	b := 2.0
	c := "2"
	d := 't' // int32 -->>tune
	e := true
	f = 255
	fmt.Printf("The %d size: %d, type is %T\n", a, unsafe.Sizeof(a), a)
	fmt.Printf("The %1.1f size: %d, type is %T\n", b, unsafe.Sizeof(b), b)
	fmt.Printf("The %s size: %d, type is %T\n", c, unsafe.Sizeof(c), c)
	fmt.Printf("The %c size: %d, type is %T\n", d, unsafe.Sizeof(d), d)
	fmt.Printf("The %t size: %d, type is %T\n", e, unsafe.Sizeof(e), e)
	fmt.Printf("The %b size: %d, type is %T\n", f, unsafe.Sizeof(f), f)
}
