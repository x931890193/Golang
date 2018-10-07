/*
Create on: 2018-10-03 下午8:47
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

var a string

func main() {
	a = "G"
	fmt.Println(a)
	func1()

}

func func1() {
	a := "a"
	fmt.Println(a)
	func2()
}

func func2() {
	fmt.Println(a)
}
