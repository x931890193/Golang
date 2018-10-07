/*
Create on: 2018-10-05 下午4:58
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	fmt.Println(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(content) // 读出来是 byte类型  全是数字
	text := string(content) // 转为string类型
	fmt.Println(len(text))
	count := CountWords(text)
	fmt.Printf("File %s has %d words\n", strings.Split(filename, "/")[2], count)
	for k, v := range strings.Fields(text) { // k 索引, v 是element
		fmt.Println(k, v)
	}
}

func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
