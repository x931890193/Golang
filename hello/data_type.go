package main // 包含main包

import "fmt"

func main() {
	// bool 长度 1  false true  初始值（零值） false
	// byte 字节型 英文字符 单引号  0
	// rune 字符  中文 单引号
	// int uint
	// float32 float64           0.0
	// complex64 complex128
	// unintger 整形
	// string 字符串 双引号

	// 1. 声明变量
	var a bool
	fmt.Println(a) // 没有初始化时 零值为 false
	a = true
	fmt.Println(a)

	// 2. 自动推导类型
	//	var b = false
	b := false
	fmt.Println(b)

	// 3. 字符类型： ASCII码 对应  65-90  97-122
	var ch byte
	ch = 97
	fmt.Println(ch)                // 得不到想要的结果
	fmt.Printf("%c\n%d\n", ch, ch) // 格式化打印 字符是单引号

	// 4.大小写转换
	fmt.Printf("大写:%d, 小写：%d\n", 'A', 'a')
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)

	fmt.Printf("hello go\n") // 格式化输出 不换行  '\n' 转义字符 换行  不会输出到屏幕， 有特殊用途

	// 5. 浮点型
	var f1 float64 // 声明变量
	f1 = 3.14
	fmt.Printf("f1=%f\n", f1)
	fmt.Println("f1=", f1)
	fmt.Printf("f1: type is %T\n", f1) // f1 : type is float64
	// float64 存储小数 比 float32更精确 但是占用空间大

	// 6. 字符串类型

	var str1 string
	str1 = "hello world!"
	fmt.Println("str1=", str1)
	fmt.Printf("str1: type is %T\n", str1) // %T 格式化输出类型

	str2 := "hello go!" // 自动推导类型
	fmt.Printf("str2: type is %T, 字符串的长度为:%d\n", str2, len(str2))

	// 7. 字符和字符串的区别
	var cha byte
	var str string

	cha = 'a'   // 字符单引号， 字符往往都只有一个字符， 转义字符除外
	str = "abc" // 字符串 双引号， 字符串有一个或多个字符串组成   字符串隐藏一个结束符 ”\0“
	fmt.Println("cha=", cha)
	fmt.Println("str=", str)
	// 字符串索引 取值
	fmt.Println(str[0], str[0], str[1], str[2])

	// 8. complex
	var com complex128
	com = 2.1 + 2.14i
	fmt.Println("com=", com)
	fmt.Printf("com: type is %T\n", com)

}
