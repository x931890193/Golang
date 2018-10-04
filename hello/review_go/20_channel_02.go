/*
Create on: 2018-10-04 下午3:30
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

// 多个goroutine的情况
var ch chan int

func main() {
	count := 1001
	ch = make(chan int)
	for i := 1; i < count; i++ {
		go Foo(i)
	}
	for i := 1; i < count; i++ {
		fmt.Println(<-ch)
	}

}

func Foo(id int) {
	//fmt.Println(id)
	ch <- id
}
