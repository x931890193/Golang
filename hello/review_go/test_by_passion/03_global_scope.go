/*
Create on: 2018-10-03 下午8:41
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main
import (
	. "fmt" //用点的方式导入, 使用的时候不用再加包名
)

var a = "A"

func main() {
	m()
	n()
	m()
}

func m(){
	Println(a)
}

func n(){
	a = "a"
	Println(a)
}