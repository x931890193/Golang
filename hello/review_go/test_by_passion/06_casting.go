/*
Create on: 2018-10-03 下午8:58
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int = 32
	var m int32
	//在格式化字符串里， %d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数
	//字） ， %g 用于格式化浮点型（%f 输出浮点数， %e 输出科学计数表示法） ， %0d 用于规
	//定输出定长的整数， 其中开头的数字 0 是必须的。
	//%n.mg 用于表示数字 n 并精确到小数点后 m 位， 除了使用 g 之外， 还可以使用 e 或者 f， 例
	//如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00 。
	m = int32(n)                         //显示的转换 int16为int32
	fmt.Printf("32 bit int is: %d\n", m) // %x 格式化16进制的数
	fmt.Printf("16 bit int is: %d\n", n) // %d 格式化输出整形, %t 格式化bool
	a, _ := Unit8FromInt(-255)           // 无符号int8 转为int16
	fmt.Println(a)
}

func Unit8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of range")
}
