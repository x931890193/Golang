package main

import (
	"fmt"
)

// 除非被显式设置为其它值或iota，每个const分组的第一个常量被默认设置为它的0值，
// 第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是iota，
// 则它也被设置为iota

// const 分组
const (
	i = iota // i=0
	j        // j = 1
	k        // k=2
)

const x = iota // x=0

const (
	q, w, e = iota, iota, iota // q=w=e=0

)

func main() {
	fmt.Printf("i=%d,j=%d, k=%d\n", i, j, k)
	fmt.Printf("x%d\n", x)
	fmt.Printf("q=%d,w=%d,e=%d\n", q, w, e)
}
