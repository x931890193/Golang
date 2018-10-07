package main

import (
	"fmt"
)

func main() {
	switch_demo()

}

func switch_demo() {
	// 每一个case 后默认自带break
	// 如需要继续执行   使用fallthrough
	i := 10
	switch i {
	case 1:
		fmt.Printf("i%s is equal to 1", i)
	case 2:
		fmt.Printf("i%s is equal to 2", i)
	case 3:
		break
	default:
		fmt.Println("nothing match this")

	}
}
