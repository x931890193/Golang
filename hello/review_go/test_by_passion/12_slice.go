/*
Create on: 2018-10-05 下午5:43
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5,6}
	newslice := slice[2:3]
	n2 := newslice[1:2]
	fmt.Println(len(newslice),len(n2))
	}
