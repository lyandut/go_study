package main

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
		for i := 0; i < 2; i++ {
			x := i                 // x 每次循环都重新定义，避免引用同一环境变量
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
	closureSync()
}
