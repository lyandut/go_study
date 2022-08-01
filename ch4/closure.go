package main

import (
	"sync"
	"time"
)

// 闭包：函数和其引用环境的组合体，本质上返回的是一个 funcval 结构体
func closureInit() {
	// test 返回的匿名函数会引用上下文环境变量 x
	test := func(x int) func() {
		println(&x)
		return func() {
			println(&x, x)
		}
	}

	f := test(0x100)
	f()
}

// “延迟求值”特性
func closureDelay() {
	test := func() []func() {
		var s []func()
		for i := 0; i < 2; i++ { // for 循环复用局部变量 i
			x := i                 // 解决方法：x 每次循环都重新定义，避免引用同一环境变量 i
			s = append(s, func() { // 将多个匿名函数添加到列表
				println(&x, x)
			})
		}
		return s // 返回匿名函数列表
	}

	for _, f := range test() { // 迭代执行所有匿名函数
		f()
	}
}

// “延迟求值”导致并发错误
func closureDelayError() {
	for i := 0; i < 2; i++ {
		//i := i // fix
		go func() {
			println(i)
		}()
	}
	time.Sleep(time.Second)
}

func closureDelayFix() {
	var wg sync.WaitGroup
	wg.Add(2) // 协程计数器
	for i := 0; i < 2; i++ {
		go func(j int) {
			defer wg.Done() // 计数器-1
			println(j)
		}(i)
	}
	wg.Wait() // 阻塞直到计数器为0
}

func closureSync() {
	test := func(x int) (func(), func()) { // 返回两个匿名函数
		return func() {
				println(x)
				x += 10 // 修改环境变量
			}, func() {
				println(x) // 显示环境变量
			}
	}

	a, b := test(100)
	a()
	b()
}

func main() {
	closureInit()
	closureDelay()
	closureDelayError()
	closureDelayFix()
	closureSync()
}
