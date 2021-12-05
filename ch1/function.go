package main

import (
	"errors"
	"fmt"
)

func functionDiv(a, b int) (int, error) { // 定义多个返回值
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func functionAnonymous(x int) func() { // 返回函数类型
	return func() { // 匿名函数
		println(x) // 闭包
	}
}

// 用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用，类似java finally
func functionDefer(a, b int) {
	defer println("dispose1...") // 常用来释放资源、解除锁定，或执行一些清理操作
	defer println("dispose2...") // 可定义多个defer，按FILO顺序执行
	println(a / b)
}

func main() {
	a, b := 10, 0               // 定义多个变量
	c, err := functionDiv(a, b) // 接收多返回值
	fmt.Println(c, err)

	x := 100
	f := functionAnonymous(x)
	f()

	functionDefer(10, 0)
}
