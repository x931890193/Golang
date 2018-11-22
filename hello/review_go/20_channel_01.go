/*
Create on: 2018-10-04 下午1:55
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
// goroutine  微线程 协程 python coroutine
package main

import "fmt"

// 建立信道 channel 信道是什么, in short 他就是goroutine之间相互通信的东西
// 用来实现goroutine之间的发信息和接受信息, 做goroutine之间的内存共享
// 以下是建立信道的两种方式  无缓存的信道,  随进随出, 进一个出一个, 多于一个会阻塞, 无论写或者读
// 无缓冲的信道在取消息和存消息的时候都会挂起当前的goroutine
// 从无缓冲信道取数据，必须要有数据流进来才可以，否则当前线阻塞
// 数据流入无缓冲信道, 如果没有其他goroutine来拿走这个数据，那么当前线阻塞

var ch1 chan string = make(chan string) // 此处的 make() 用于内建类型的内存分配
// ch1 := make(chan int) // 自动推导类型 需要在{} 中使用
var cache_ch1 chan int = make(chan int, 3) // 声明一个有缓冲的信道 , 第二个参数指明缓冲数据大小
//  buffered channel.
func main() {
	ch2 := make(chan string)
	go Demo1()                   // 开启 goroutine
	fmt.Println(len(ch1), <-ch1) // 从管道中读取数据, 无缓冲的管道, len永远是0
	go func(msg string) {
		ch2 <- msg // 放入管道
	}("ping!")
	fmt.Println(<-ch2) // 读取数据

	Demo2() // 正常调用函数
	fmt.Println("-------------------")
	buffered_chan := make(chan string, 10)
	go Demo3(buffered_chan) // 开启goroutine 参数为 channel
	for v := range buffered_chan {
		// 是range不等到信道关闭是不会结束读取的。也就是如果缓冲信道干涸了，那么range就会阻塞当前goroutine, 所以死锁.
		if len(buffered_chan) == 0 { // 判断channel的长度  如果长度为0,跳出循环
			break
		}
		fmt.Println(v)
	}
	fmt.Println("------------------")
	buffered_chan2 := make(chan int, 20)
	go Demo4(buffered_chan2)
	for v := range buffered_chan2 {
		fmt.Println(v)
	}

}

func Demo1() {
	// 无缓冲的信道 函数demo
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	messages := "i am done"
	ch1 <- messages // 把messages 放入管道

}

func Demo2() {
	// 有缓冲的信道  函数demo
	cache_ch1 <- 10
	cache_ch1 <- 11
	cache_ch1 <- 12
	// 管道缓冲为 3 再添加数据 报死锁
	// fatal error: all goroutines are asleep - deadlock!
	//cache_ch1 <- 13

	fmt.Println(<-cache_ch1)
	fmt.Println(<-cache_ch1)
	fmt.Println(<-cache_ch1)
	// 同理 缓冲长度为3  再次读取 报死锁
	//fmt.Println(<- cache_ch1)

}

func Demo3(ch chan string) {
	// 信道数据读取和信道关闭
	for i := 0; i < 10; i++ {
		ch <- "hello world"
	}

}

func Demo4(ch chan int) {
	for i := 1; i < 21; i++ {
		ch <- i
	}
	// 被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。
	close(ch) // 显示的关闭信道
}
