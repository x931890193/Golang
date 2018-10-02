package main

import (
    "fmt"
)
var HelpText string=`Go程序设计的一些规则

Go之所以会那么简洁，是因为它有一些默认的行为：

    大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
    大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
`

func main(){
    fmt.Print(HelpText)
    fmt.Println("-----------------")
    iamarray()
    fmt.Println("-----------------")
    iamdynamic_array()
    fmt.Println("-----------------")
}


func iamarray(){
    var arr [20]int // 声明一个int类型的数组, 长度为20
    arr[0] = 1
    arr[2] = 2
    // 数组不能改变长度, 当作为参数传入函数的时候 传递的是副本
    fmt.Println(arr, len(arr)) // len()函数 获取长度
    arr2 := [2]int{1,2} // 声明长度为2的int数组, 并且赋初值
    fmt.Println(arr2)
    arr3 := [3]string{"hello", "world", "nice to meet you"}
    fmt.Println(arr3)
    arr4 := [2][4]int{[4]int{1,2,3,4},[4]int{2,3,4,5}} // 二维数组
    fmt.Println(arr4)
    arr5 := [3][2]string{{"aa", "bb"},{"cc","dd"},{"ee","ff"}}
    fmt.Println(arr5)
}

func iamdynamic_array(){
    // slice
    // 声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。
    arr := []string{"hello"} // 声明一个动态数组 不指定长度, 字符串 元素
    fmt.Println(arr)
    arr1 := []byte{'a','b','c','d'}
    fmt.Println(string(arr1))
    var i, j []byte
    i = arr1[2:3] // arr1[:n] == arr1[0:n] 切片操作
    j = arr1[0:1] // arr[n:] == arr1[n:len(arr1)]
    fmt.Println(string(i), string(j))


}


