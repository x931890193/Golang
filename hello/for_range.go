package main

import "fmt"

func main() {
	// for 循环
	// for 初始条件; 判断条件； 条件变化{   }
	// 1 + 2 + --- 100

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1 + 2 + --- 100	=", sum)

	str := "abcdef"
	// 通过 for打印字符
	for j := 0; j < len(str); j++ {
		fmt.Printf("%c\n", str[j])
	}
	// range 迭代 配合数组和切片 (enumerate)  默认返回两个值
	for _, data := range str {
		fmt.Printf("%c\n", data)
	}
	// 元素下标
	for k := range str {
		fmt.Printf("str[%d]\t%c\n", k, str[k])
	}

}
