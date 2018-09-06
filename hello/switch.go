package main

import "fmt"

// switch 语句
func main() {
	//	num := 10  关键字  fallthrough

	var num int
	fmt.Printf("请选择楼层:")
	fmt.Scan(&num)
	switch num { // switch 后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼", num)
		//		break
	case 2:
		fmt.Println("按下的是2楼", num)
		//		break
		fallthrough // 满足上边条件 无条件的往下执行
	case 3:
		fmt.Println("按下的是3楼", num)
		//		break
	case 4:
		fmt.Println("按下的是4楼", num)
		//		break
	default:
		fmt.Printf("按下的是%d楼\n", num)
	}
}
