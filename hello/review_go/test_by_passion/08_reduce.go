/*
Create on: 2018-10-03 下午10:51
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import "fmt"

func main() {
	//ReducePrint(100)
	res := ReduceFactorial(13)
	fmt.Println(res)
}

func ReducePrint(n int) {
	if n < 1 {
		return
	} else {
		fmt.Println(n)
		ReducePrint(n - 1)
	}

}

func ReduceFactorial(n int) (res int) {
	if n == 1 {
		return 1
	} else {
		return n * ReduceFactorial(n-1)
	}

}
