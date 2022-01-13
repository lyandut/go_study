package main

import "fmt"

func forInit() {
	count := func() int {
		print("count.")
		return 3
	}

	for i, c := 0, count(); i < c; i++ { // 初始化语句的 count 函数仅执行一次
		println("a", i)
	}

	c := 0
	for c < count() { // 条件表达式中的 count 重复执行 4 次
		println("b", c)
		c++
	}
}

func forRange() {
	data := [3]string{"a", "b", "c"}

	for i := range data { // 只返回 1st value
		println(i, data[i])
	}

	for _, s := range data { // 忽略 1st value
		println(s)
	}

	for range data { // 仅迭代，不返回。可用来执行清空 channel 等操作
	}

	for i, s := range data {
		println(i, s, &i, &s) // 定义的局部变量会重复使用
	}
}

func forArraySlice() {
	data := [3]int{10, 20, 30}
	for i, x := range data { // 从 data 复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

	data = [3]int{10, 20, 30}
	for i, x := range data[:] { // 仅复制 slice，不包括底层 array
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}
}

func forFunc() {
	data := func() []int {
		println("origin data.")
		return []int{10, 20, 30}
	}

	for i, x := range data() { // range 目标表达式是函数调用，仅被执行一次
		println(i, x)
	}
}

func main() {
	forInit()
	forRange()
	forArraySlice()
	forFunc()
}
