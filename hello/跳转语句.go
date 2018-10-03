package main

import (
	"fmt"
	"time"
)

// break continue 区别
func main() {
	// for 后边不接任何东西  条件永远满足 死循环
	go test()
	//time.Sleep(5000000000)
}

func test(){
	i := 0
	for {
		i++

		time.Sleep(time.Second) // 延时1秒
		fmt.Println(i)

		if i == 5000000 {
			fmt.Println(i)
			break // 跳出最近的循环
			// continue //跳过本次循环
		}

	}
}
