package main

import (
	"fmt"
	"strings"
)

var (
	s      string = "hello world! fantastic world"
	n1, n2 int
	b1, b2 bool
)

func main() {
	// func Count(s, seq string)int
	n1 = strings.Count(s, "l")
	n2 = strings.Count(s, "o")
	fmt.Printf("%s中l的个数为:%d\n", s, n1)
	fmt.Printf("%s中o的个数为:%d\n", s, n2)
	// func Contains(s,substr string)bool
	b1 = strings.Contains(s, "world")
	fmt.Printf("%s是否包含world: ", s)
	fmt.Println(b1)
	// func ContainsAny(s, chars string)bool
	// s 字符串中是否包含chars中的任意一个字符串 双引号
	b2 = strings.ContainsAny(s, "abcde")
	fmt.Printf("%s是否包含abcde中的任意字符: ", s)
	fmt.Println(b2)
	// func ContainsRune(s tring, r rune)bool
	b1 = strings.ContainsRune(s, 'a')
	fmt.Printf("%s是否包含a字符", s)
	fmt.Println(b1)
	// index    func Index(s, seq string)int rabin-karp 算法实现
	// seq 子串第一次出现的位置 没有为-1, 如果seq为空 返回为0
	n1 = strings.Index(s, "a")
	fmt.Printf("%s中a的索引为:%d\n", s, n1)
	// LastIndex 返回子串最后一次出现的位置
	// 没有为-1 如果为空 返回s的长度

}
