package main

import (
	"fmt"
	"log"
	"time"
)

func parameterTest(x, y int, s string, _ bool) *int { // 参数列表中，相邻的同类型参数可合并
	return nil
}

func parameterPassByValue() {
	test := func(x *int) {
		fmt.Printf("pointer: %p, target: %v\n", &x, x) // 输出形参 x 的地址
	}

	a := 0x100
	p := &a
	fmt.Printf("pointer: %p, target: %v\n", &p, p) // 输出实参 p 的地址
	test(p)
}

// 示例：指针参数导致实参变量被分配到堆上
func parameterEscape(p *int) {
	go func() { // 延长 p 生命周期
		println(p)
	}()
}

// 使用二级指针实现传出参数（out）
func parameterOut() {
	test := func(p **int) {
		x := 100
		*p = &x
	}

	var p *int
	test(&p)
	println(*p)
}

// 利用复合结构类型，变相实现可选参数和命名实参功能
func parameterStruct() {
	type serverOption struct {
		address string
		port    int
		path    string
		timeout time.Duration
		log     *log.Logger
	}

	newOption := func() *serverOption {
		return &serverOption{ // 默认参数
			address: "0.0.0.0",
			port:    8080,
			path:    "/var/test",
			timeout: time.Second * 5,
			log:     nil,
		}
	}

	server := func(option *serverOption) {
		fmt.Printf("%+v\n", *option)
	}

	opt := newOption()
	opt.port = 8085 // 命名参数设置
	server(opt)
}

// 变参本质上就是一个切片。只能接收一个到多个同类型参数，且必须放在列表尾部
func parameterVariable() {
	test1 := func(a ...int) {
		fmt.Printf("%T, %v\n", a, a) // 显示类型和值
	}

	test2 := func(a ...int) {
		for i := range a {
			a[i] += 100
		}
	}

	test1(1, 2, 3, 4)

	a := [3]int{10, 20, 30}
	test1(a[:]...) // 数组 -> 转换为 slice -> 展开
	test2(a[:]...)
	fmt.Println(a)
}

func main() {
	parameterTest(1, 2, "abc", true)

	parameterPassByValue()

	x := 100
	p := &x
	parameterEscape(p)

	parameterOut()

	parameterStruct()

	parameterVariable()
}

/*
   $ go build -gcflags "-m" parameter.go       // 输出编译器优化策略
		moved to heap: x
   $ go tool objdump -s "main\.main" parameter
		CALL runtime.newobject(SB)    		   // 在堆上为 x 分配内存
		CALL main.parameterEscape(SB)
*/
