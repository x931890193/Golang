package main //main 包

// 传统写法
//import "fmt" //导入包
//import "os"  //

// 圆括号形式
import (
	//	"time"
	"fmt"
	"os"
)

//以点的形式导入
import . "fmt" //
import . "os"  //  使用的时候不用包名字

// 给包起别名
import io "fmt"

// 忽略某个包
import _ "fmt" // 不用此包  为了使用他的init函数

func main() {

	fmt.Println("测试！")
	fmt.Println(os.Args)
}
