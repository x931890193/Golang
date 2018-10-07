/*
Create on: 2018-10-04 下午4:17
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/

package main

import (
	"C"
)

// go build -buildmode=c-shared -o 输出文件名.so  需要编译的文件名.go
// 指定哪些函数内被外部调用 指明函数test能被外部调用
//export test
func test() int {
	// 计算0-100000的和
	var sum int
	for i := 0; i <= 10000; i++ {
		sum += i
	}
	return sum
}

//export addstr
func addstr(a, b *C.char) *C.char {

	merge := C.GoString(a) + C.GoString(b)
	return C.CString(merge)

}

func main() {

}
