/*
Create on: 2018-10-04 下午4:03
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

func main() {
	generator := my_range()
	for i := 0; i < 1000; i++ {
		fmt.Println(<-generator)
	}
}

func my_range() chan int {
	var ch chan int = make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()
	return ch
}
