package main // 必须有main 包

import "fmt"

// 函数名  大写为公有   小写为私有   不支持默认参数
//func FuncName(参数 类型， 参数 类型) ( 返回值 类型， 返回值， 类型){

//	函数体

//}

// 1.无参数， 无返回的函数的定义
func MyFunc() {
	a := 100
	fmt.Println("a=", a)
}

// 2.有参数， 没有返回值的函数的定义
// 参数类型一样时   a, b int, c, d, float128
func MyFunc2(a int, b float32) { //形参 参数传递只能由实参传给形参
	//	a = 300
	fmt.Println(a, b)

}

func MyFunc3(a, b int, c, d string, e, f float64) {
	fmt.Println(a, b, c, d, e, f)
}

// 3. 不定长参数列表  ...type 不定长参数类型  注意：+++++++不定长参数只能放在形参中的最后一个参数+++++++
func MyFunc4(args ...int) {
	fmt.Printf("len(args)=%d\n", len(args))
	for x := 0; x < len(args); x++ {
		fmt.Printf("args[%d]=%d\n", x, args[x])
	}
	fmt.Println("------------")
	for data, _ := range args { // 迭代的使用 返回两个值  第一个为下标 第二个为下标对应的数
		fmt.Printf("args[%d]=%d\n", data, args[data])
	}

}

// 4. 无参数， 有返回值的函数定义
//    4.1 (只有一个返回值的时候)   可以省略圆括号
// 有返回值的函数 需要 return中断函数
func MyFunc5() int { // 一个返回值 不用给变量名
	return 1024
}

// 给返回值 起一个变量名 go 推荐写法
func MyFunc6() (result int) {
	result = 2048
	return result // 常用写法
}

// 4.2 函数有多个返回值

func MyFunc7() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return a, b, c

}

// 5, 有参数， 有返回值的函数

func MyFunc8(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return max, min

}

func main() { // 有且只有一个入口
	MyFunc()           // 函数定义的位置没有先后
	MyFunc2(300, 3.14) // 调用函数， 传递的参数叫实参
	MyFunc3(1, 2, "ABC", "DEF", 3.14, 1.732)
	MyFunc4(12, 45)
	ret := MyFunc5()
	fmt.Println("函数返回值为:", ret)
	fmt.Println("函数返回值为:", MyFunc6())
	a, b, c := MyFunc7()
	fmt.Printf("函数返回值为:%d,%d,%d\n", a, b, c)
	max, min := MyFunc8(10, 100)
	fmt.Printf("max=%d, min=%d\n", max, min)
}
