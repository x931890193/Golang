package main

import (
	"fmt"
)

// 匿名字段的结构体 结构体嵌套?

// 声明一个strcut
type Person struct {
	name   string
	age    int
	weight int
	gender bool
}

type Stu1 struct {
	Person
	class int
}

type Skill []string // []string 类型

// 声明一个带有匿名字段的struct
// 内置类型和自定义类型都是可以作为匿名字段
type Stu struct {
	Person // 匿名字段 包含person的所有字段
	class  int
	Skill  // 自定义类型 作为匿名字段
	int    // 内置类型作为字段
}

func main() {
	// 初始化一个学生 p1
	p1 := Stu1{Person{"zs", 18, 50, true}, 2}
	fmt.Println(p1)
	fmt.Println(p1.Person)
	fmt.Println(p1.Person.age)
	// 初始化 p2
	p2 := new(Stu)
	p2.name = "perter"
	p2.age = 18
	p2.int = 20
	fmt.Println(p2)
}
