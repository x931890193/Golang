package main

import "fmt"

func main() {

	var b int
	fmt.Scan(&b)
	//	b := 10

	if b == 10 {
		fmt.Println("a=10")
	}

	if b > 10 {
		fmt.Println("a>10")
	}

	if b < 10 {
		fmt.Println("a<10")
	}

}
