/*
Create on: 2018-10-03 下午7:39
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i:= 0; i<10;i++{
		fmt.Printf("%d\n",rand.Int()) // 随机整数 10个
	}
	fmt.Println("------")
	for i:=0; i<10; i++{
		x := rand.Intn(8) // 随机生成0-7 的整数 [0 8)
		fmt.Printf("%d\n",x)
	}
	fmt.Println("-------")
	for i:=0; i<10; i++{
		timens := int64(time.Now().Nanosecond())
		rand.Seed(timens)
		fmt.Printf("%2.3f\n", 100*rand.Float32()) // 格式化小数, 小数点 前后数字指明位数
	}
}
