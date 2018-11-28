package main

import (
	"errors"
	"fmt"
)

//type FuncType func()  //自定义一个函数类型

func main() {
	fmt.Println("1,Boolean类型, true or false")
	fmt.Println("2.int, unit类型, 长度相同,具体长度根据编译器实现", "int8, int16,int32", "rune==int32", "byte==uint8")
	// go 强类型语言, 不同类型的数据不能进行运算
	fmt.Println("3.浮点类型, float32, float64")
	fmt.Println("4.complex, 复数 ai+bj")
	fmt.Println("5.字符串类型,字符串不可变, 不能切片赋值, 双引号(单行字符串), 或者反引号(多行字符串) ``, byte使用单引号''")
	fmt.Println("6.error 类型 错误类型, 专门的一个处理处理错误的包")
	test()
	change_string()
	error_type()
	var t []byte = []byte{'a'}
	fmt.Println(t)
	g := []byte("dsadasdas")
	fmt.Println(g)
}

func test() {
	no, yes, maybe := "no", "yes", "maybe" // 自动推导类型
	china := "welcome to china"
	french := "Bonjour" // 常规赋值
	fmt.Println(no, yes, maybe, china, french)
}

func change_string() {
	var (
		s2 []byte
		s3 string
	)

	s1 := "hello"
	s2 = []byte(s1) // 把字符串转成[]byte类型, byte类型的[]
	s2[0] = 'w'
	s3 = string(s2)
	fmt.Println(s3)
	s4 := s1 + " " + s3
	fmt.Println(s4)
}

func error_type() {
	err := errors.New("This is a error, when you see this")
	if err != nil {
		fmt.Println(err)
	}

}
