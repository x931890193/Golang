/*
Create on: 2018-10-05 下午2:27
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
	"sort"
)

var (
	// 初始化一个map key为string value为int
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key %s, value %d\n", k, v)
	}
	keys := make([]string, len(barVal)) // 声明一个slice 长度为len(barVal)
	fmt.Printf("%T  %d", keys, len(barVal))
	i := 0
	for k, _ := range barVal { // 拿出map的key 放入slice中
		keys[i] = k
		i++
	}
	sort.Strings(keys) // 使用自带的字符串排序, 对slice排序
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys { // 舍弃第一个数(索引), 第二个数为element
		fmt.Printf("key %s, value %d\n", k, barVal[k])
	}
}
