package main

import "fmt"

func test() (sum int) {
	for i := 0; i <= 100; i++ {

		sum += i
	}
	return sum

}

// 以下展示递归的两种用法

func test2(i int) int { //demo 1

	if i == 1 {
		return 1
	} else {
		return i + test2(i-1)
	}
}

func test3(i int) int { //demo2
	if i == 100 {
		return 100
	} else {
		return i + test3(i+1)
	}

}

func main() {

	sum := test() //调用函数实现加法 传统写法
	fmt.Println("sum=", sum)

	sum2 := test2(100)
	fmt.Println("sum=", sum2)

	sum3 := test3(0)
	fmt.Println("sum=", sum3)
}
