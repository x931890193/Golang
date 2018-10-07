/*
Create on: 2018-10-05 下午2:44
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
	new_map := make(map[int]string, len(barVal))
	values := make([]int, len(barVal))
	i := 0
	for k, v := range barVal {
		new_map[v] = k
		values[i] = v
		i++
	}
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(new_map)
	for k, v := range new_map {
		fmt.Printf("k=%d, v=%s\n", k, v)
	}

	muti_values := map[int][]string{1: {"str", "sss"}} // 多值map  one key muti value
	fmt.Println(muti_values)
}
