package main

import (
	"fmt"
)

// 声明一个结构体 四个字段 name age gender job
type person struct {
	name   string
	age    int
	gender bool
	job    string
}

func main() {
	// 结构体使用
	var (
		P  person
		p2 person
		p3 person
	) // P 是一个person类型的变量
	// 属性赋值 方法1
	P.name = "xiaoqiang"
	P.age = 18
	P.gender = true
	//P.job = "lawyer" // 可以不用全部赋值
	fmt.Println(P)
	// 属性赋值 方法2 必须全部赋值
	p2 = person{"laowang", 18, false, "stu"} // 需要全部赋值
	fmt.Println(p2)
	// 第三种赋值方法 需要全部赋值
	p3 = person{name: "zf", age: 80, gender: true, job: "teacher"}
	fmt.Println(p3)
	// 使用new函数分配指针  make(内置类型分配空间), new(自定义类型分配空间)
	p4 := new(person)
	fmt.Println(p4)
	p, age := Older(P, p3)
	fmt.Println("---------")
	fmt.Println(P)
	fmt.Println(p3)
	fmt.Printf("年龄大的是:%s, 年龄差:%d\n", p.name, age)
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age

}
