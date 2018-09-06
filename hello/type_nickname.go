package main // main

import "fmt"

func main() { // entry

	type bigint int64                 // int64's nickname
	var a bigint                      // equal  int64
	fmt.Printf("a's type is %T\n", a) //a's type is main.bigint

	type (
		long int64
		char byte
	)
	var b long = 11
	var ch char = 'a'
	fmt.Printf("b's type is %T\n", b)
	fmt.Printf("ch's type is %T\n", ch)

}
