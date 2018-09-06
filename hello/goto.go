package main

import "fmt"

func main() {
	//	break // break is not in a loop, switch, or select
	//	continue // break is not in a loop, switch, or select
	//	goto 可以再任何地方使用， 不能跨函数使用
	fmt.Println("test is tuning")
	// goto  不要乱用  无条件跳转
	goto End // goto是关键字  End是标签 标志

End: // 无goto 标签无用
	fmt.Println("测试通过！")

}
