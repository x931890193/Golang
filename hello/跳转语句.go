package main

import (
	"fmt"
	//"time"
)

var ch chan int = make(chan int)

// break continue 区别
func main() {
	// for 后边不接任何东西  条件永远满足 死循环
	go test()
	<-ch // 从管道中取东西， 阻塞主进程
	//time.Sleep(5000000000)
}

func test() {
	i := 0
	for {
		i++

		//time.Sleep(time.Second) // 延时1秒
		fmt.Println(i)

		if i == 5000 {
			fmt.Println(i)
			ch <- 0 // 放入管道  信道
			break   // 跳出最近的循环
			// continue //跳过本次循环
		}
	}
}
