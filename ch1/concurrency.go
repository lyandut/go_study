package main

import (
	"fmt"
	"time"
)

func concurrentGoroutine() {
	task := func(id int) {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: %d\n", id, i)
			time.Sleep(time.Second)
		}
	}

	go task(1) // 创建 goroutine
	go task(2)
	time.Sleep(time.Second * 6)
}

// channel 与 goroutine 搭配，实现用通信代替内存共享的CSP模型
func concurrentChannel() {
	consumer := func(data chan int, done chan bool) {
		for x := range data { // 接收数据，直到通道被关闭
			println("recv:", x)
		}
		done <- true // 通知main，消费结束
	}

	producer := func(data chan int) {
		for i := 0; i < 4; i++ {
			data <- i // 发送数据
		}
		close(data) // 生产结束，关闭通道
	}

	done := make(chan bool) // 用于接收消费结束信号
	data := make(chan int)  // 数据管道
	go consumer(data, done)
	go producer(data)
	<-done // 阻塞，直到消费者发回结束信号
}

func main() {
	concurrentGoroutine()
	concurrentChannel()
}
