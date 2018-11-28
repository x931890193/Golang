package main

import (
	"fmt"
	//    "file"
)

func main() {
	var a, b, c int
	fmt.Println("请输入a:")
	fmt.Scan(&a)
	fmt.Println("请输入b:")
	fmt.Scan(&b)
	fmt.Println("请输入c:")
	fmt.Scan(&c)
	fmt.Printf("max(a,b)=%d\n", max(a, b))
	fmt.Printf("min(a,b,c)=%d\n", min(a, b, c))
	myfunc(11, 22, 33, 44, 55, 66)
	a := 1
	fmt.Printf("原始值:a=%d\n", a)
	fmt.Printf("传值返回:a=%d\n", myfunc2(a))
	fmt.Printf("传值以后:a=%d\n", a)
	fmt.Printf("传指针返回:a=%d\n", myfunc3(&a))
	fmt.Printf("传指针以后:a=%d,原始值改变\n", a)
	// defer 的使用 最后调用, 清理现场
	fmt.Println(my_defer())
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else if b < a {
		if b < c {
			return b
		} else {
			return c
		}
	} else {
		if c < a {
			return c
		} else {
			return a
		}
	}
}

func myfunc(args ...int) {
	// 不定长参数  变参 传递
	for index, n := range args {
		fmt.Println(index, n)
	}
}

func myfunc2(a int) int {
	//传值, 原值不变
	// 传递给函数的是原数据的copy
	a += 1
	return a
}

func myfunc3(a *int) int {
	// 传指针,修改原始值
	*a += 1
	/*
	   1.传指针使得多个函数能操作同一个对象。
	   2.传指针比较轻量级(8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
	   3.Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针*/
	return *a
}

func my_defer() bool {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	//    file.Open("./test.txt")
	//    defer file.Close()
	//    if failureX{
	//        return false
	//    }
	//    if failureY{
	//        return false
	//    }
	return true
}
