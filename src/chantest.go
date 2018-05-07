package main

import (
	"fmt"
	"time"
)

//func main() {
//	// 使用`make(chan 数据类型)`来创建一个Channel
//	// Channel的类型就是它们所传递的数据的类型
//	messages := make(chan int)
//
//	// 使用`channel <-`语法来向一个Channel写入数据
//	// 这里我们从一个新的协程向messages通道写入数据ping
//
//	for sum := 0; sum < 10; sum++ {
//		//s := strconv.Itoa(sum)
//		messages <- sum
//		//go func() { }()
//	}
//
//	// 使用`<-channel`语法来从Channel读取数据
//	// 这里我们从main函数所在的协程来读取刚刚写入
//	// messages通道的数据
//	for i := 0; i < 10; i++ {
//		v := <-messages
//		fmt.Println("receive:", v)
//	}
//	//fmt.Scanln(&msg)
//	//fmt.Println(msg)
//	time.Sleep(1 * time.Second)
//}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}