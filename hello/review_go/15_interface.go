package main

// interface 是一组method的签名, 通过interface来定义对象的一组行为
// interface类型定义了一组方法, 如果某个对象实现了某个接口的所有方法,
// 则此对象实现了此接口
// 你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
// Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。
// 一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。
import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human //匿名字段
	company string
	money   float32
}

// human 实现了sayhi方法
//func (h, *Human)Sayhi() {
//	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
//}

//human 实现sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("la la la", lyrics)

}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main(){



}