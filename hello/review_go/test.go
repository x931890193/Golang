package main

import (
	"fmt"
)

// ascii 65-90 97-122
func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(string(i))
	}
	for j := range [5]int{97, 98, 99, 100, 122} {
		fmt.Println(string(j))
	}
	var a int = 65
	b := string(a)
	fmt.Println(b)

}
