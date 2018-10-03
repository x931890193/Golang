/*
Create on: 2018-10-03 下午8:37
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

var a = "A"
func main() {
	m()
	n()
	m()
}

func m(){
	fmt.Println(a)
}
func n(){
	a := "a" //local variable
	fmt.Println(a)
}