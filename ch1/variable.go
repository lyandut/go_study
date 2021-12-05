package main

func main() {
	var x int32             // 默认初始化为零值
	var s = "Hello, World!" // 支持类型推断
	y := 100                // short variable declaration, 只能用在函数内部
	println(x, s, y)

	x = 200 // 赋值，区别于定义
	println(x, s, y)
}
