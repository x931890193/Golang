package main

import "fmt"
import "os" // os 包

// 获取shell 参数
func main() {
	// 接受用户传递的参数. 以字符串的方式传递
	list := os.Args
	n := len(list) // 参数的个数 和python 类似  sys.argv 第一个参数是文件本身
	fmt.Println(n)
	for i := 0; i < n; i++ { // for 循环
		fmt.Printf("list[%d]=%s\n", i, list[i])
	}

	for num, data := range list { // 迭代实现
		fmt.Printf("list[%d]=%s\n", num, data)
	}

}
