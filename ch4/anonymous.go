package main

// 直接执行
func anonymousInit() {
	func(s string) {
		println(s)
	}("hello world!")
}

// 赋值给变量
func anonymousVariable() {
	add := func(x, y int) int {
		return x + y
	}

	println(add(1, 2))
}

// 作为参数
func anonymousParameter() {
	test := func(f func()) {
		f()
	}

	test(func() {
		println("hello world!")
	})
}

// 作为返回值
func anonymousReturn() {
	test := func() func(int, int) int {
		return func(x, y int) int {
			return x + y
		}
	}

	add := test()
	println(add(1, 2))
}

// 作为结构体字段
func anonymousStruct() {
	type calc struct { // 定义结构体类型
		mul func(x, y int) int // 函数类型字段
	}
	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}

	println(x.mul(2, 3))
}

// 经通道传递
func anonymousChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(x, y int) int {
		return x + y
	}

	println((<-c)(1, 2))
}

func main() {
	anonymousInit()
	anonymousVariable()
	anonymousParameter()
	anonymousReturn()
	anonymousStruct()
	anonymousChannel()
}
