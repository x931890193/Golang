package main

import "fmt"

func main() {
	// 算术运算符
	// + - * / % 取模
	// a := 1   a++;   ---> a = a + 1
	// b := 1   a--;  -->> b = b - 1
	fmt.Println("4 > 3", 4 > 3)
	//  != 不等于
	fmt.Println("4 != 3", 4 != 3)

	// 逻辑运算符  ！ 非   && 与   ||或
	fmt.Println("!(4 == 3) 结果:", !(4 == 3))
	//
	fmt.Println("true && true 结果:", true && true)
	//
	fmt.Println("true || false 结果:", true || false)

	// 关系运算符  >=  ==  <=   <    >
	// 0 < a < 8 整形和布尔类型不兼容
	a := 8 // 自动推导
	fmt.Println("a=", a, "a > 0 && a < 8", a > 0 && a < 8)

	// 赋值运算符
	// +=  -=  %=   &(取地址)   *(指针地址)
	// 优先级  ^  !  --> 7
	// 运算符优先级
	// && ||   低

}
