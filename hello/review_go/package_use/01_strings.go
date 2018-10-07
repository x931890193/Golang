package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	s      string = "hello world! fantastic world"
	n1, n2 int
	b1, b2 bool
)

func main() {
	// 字符串统计个数
	// func Count(s, seq string)int
	n1 = strings.Count(s, "l")
	n2 = strings.Count(s, "o")
	fmt.Printf("%s中l的个数为:%d\n", s, n1)
	fmt.Printf("%s中o的个数为:%d\n", s, n2)
	// 字符串是否包含子串
	// func Contains(s,substr string)bool
	b1 = strings.Contains(s, "world")
	fmt.Printf("%s是否包含world: ", s)
	fmt.Println(b1)
	// func ContainsAny(s, chars string)bool
	// s 字符串中是否包含chars中的任意一个字符串 双引号
	b2 = strings.ContainsAny(s, "abcde")
	fmt.Printf("%s是否包含abcde中的任意字符: ", s)
	fmt.Println(b2)
	// 是否包含某个字符 rune
	// func ContainsRune(s tring, r rune)bool
	b1 = strings.ContainsRune(s, 'a')
	fmt.Printf("%s是否包含a字符", s)
	fmt.Println(b1)
	// 索引位置
	// index    func Index(s, seq string)int rabin-karp 算法实现
	// seq 子串第一次出现的位置 没有为-1, 如果seq为空 返回为0
	n1 = strings.Index(s, "a")
	fmt.Printf("%s中a的索引为:%d\n", s, n1)
	// LastIndex 返回子串最后一次出现的位置
	// 没有为-1 如果为空 返回s的长度
	// 字符串以结尾 开始
	// strings.HasPrefix(s, prefix string)bool
	// strings.HasSuffix(s, suffix string)bool
	// 字符串替换
	// strings.Replace(s, old, new, n) string
	// 字符串转换转小写
	// strings.ToLower(s) string
	// 字符串转大写
	// strings.ToUpper(s) string
	// 去除开头和结尾的空白字符
	// strings.TrimSpace(s)
	// 去除首尾指定的string
	// strings.Trim(s, "string")
	// 字符串分割 , 一个或多个空白符分割
	// string.Fields(s)
	// 指定分割条件
	// strings.Split(s, seq)  返回slice  常用 for_range 处理
	// 拼接slice 到字符串
	// strings.Join(sl []string, seq string)
	//Strconv()
}

func Strconv() {
	fmt.Println(666666)
	// strconv.Itoa() int --> string
	var a int = 10
	str_a := strconv.Itoa(a)
	fmt.Printf("str_a's type is %T\n", str_a)

	var b string = "hello"
	int_b, _ := strconv.Atoi(b)
	fmt.Printf("int_b's type is %T\n", int_b)
	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, err, s)

}
