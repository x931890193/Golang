package main

import "fmt"

// switch 语句
func main() {
	//	num := 10  关键字  fallthrough

	//	var num int
	//	fmt.Printf("请选择楼层:")
	//	fmt.Scan(&num)
	// 支持一个初始化语句, 初始化语句和变量本身
	switch num := 10; num { // switch 后面写的是变量本身
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
	case 5, 6, 7, 8, 9:
		fmt.Println("xxxxxxxx")
	default:
		fmt.Printf("按下的是%d楼\n", num)
	}

	//	score := 85
	var score int
	fmt.Printf("请输出分数！")
	fmt.Scan(&score)
	switch { // 可以没有条件

	case score > 90:
		fmt.Println("优秀！")
	case score > 85:
		fmt.Println("良好！")
	case score > 60:
		fmt.Println("及格！")

	}

}
