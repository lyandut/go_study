package main

import (
	"fmt"
	"log"
	"os"
)

func deferClosure() {
	x, y := 1, 2

	defer func(x int) {
		println("defer x, y =", x, y) // y 为闭包引用
	}(x) // 注册时复制调用函数

	x += 100 // 对 x 的修改不会影响延迟调用
	y += 200
	println(x, y)
}

func deferFILO() {
	defer println("a")
	defer println("b")
}

func deferReturn() {
	test := func() (z int) {
		defer func() {
			println("defer:", z)
			z += 100 // 修改命名返回值
		}()

		return 100 // 实际执行次序：z = 100, call defer, ret
	}

	println("test:", test())
}

func deferLoop() {
	// 日志处理算法
	do := func(n int) {
		path := fmt.Sprintf("./log/%d.txt", n)

		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return
		}

		// 该延迟调用在此匿名函数结束时执行，而非 main
		defer f.Close()

		//... do something ...
	}

	for i := 0; i < 10000; i++ {
		do(i)
	}
}

func main() {
	deferClosure()
	deferFILO()
	deferReturn()
	deferLoop()
}
