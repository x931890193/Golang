package main

import (
    "fmt"

)

func main() {
    res := i_am_a_map(1,2)
    fmt.Printf("函数返回值:%d\n",res)
    make_new()
}

func i_am_a_map(a,b int) int {
    // 定义一个map
    var m1 map[string]int // nil  不能进行赋值
    // 使用map 创建一个非nil的map
    m1 = make(map[string]int)
    m1["one"] = 1
    // 直接创建map
    m2 := make(map[string]int)
    m2["one"] = 1
    m2["two"] = 2

    // 创建且赋值
    m3 := map[string]float32{"one":1.0, "two":2.0, "three":3.0 }
    fmt.Println("before:","\nm1=",m1,"\nm2=",m2,"\nm3=",m3)
    if status, ok := m3["one"]; ok{
        fmt.Println(ok)
        fmt.Println(status, "is 1")
    }else{
        fmt.Println("no")
    }
    delete(m3,"three")
    fmt.Println("after:", "\nm1=",m1,"\nm2=",m2,"\nm3=",m3)
    return a+b

}

func make_new(){

    fmt.Println("make 用于内建类型(map, slice, channel)的内存分配")
    fmt.Println("new用于各种类型的内存分配, new 返回指针")


}





