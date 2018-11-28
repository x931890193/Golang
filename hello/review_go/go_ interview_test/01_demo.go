package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("i got you")
		}
	}()
	panic("触发异常") //

}
