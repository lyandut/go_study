package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func errorInit() {
	var errDivByZero = fmt.Errorf("division by zero") // fmt.Errorf 与 errors.New 类似

	div := func(x, y int) (int, error) {
		if y == 0 {
			return 0, errDivByZero
		}
		return x / y, nil
	}

	z, err := div(5, 0)
	if err == errDivByZero { // 应通过错误变量，而非文本内容来判定错误类型
		log.Fatalln(err) // 会调用 os.Exit(1) 退出程序，且不会执行 defer 延迟调用
	}
	println(z)
}

type DivError struct { // 自定义错误类型
	x, y int
}

func (DivError) Error() string { // 实现 error 接口方法
	return "division by zero"
}

func errorStruct() {
	div := func(x, y int) (int, error) {
		if y == 0 {
			return 0, DivError{x, y}
		}
		return x / y, nil
	}

	z, err := div(5, 0)
	if err != nil {
		switch e := err.(type) { // 根据类型匹配
		case DivError: // 注意 case 顺序，应将自定义类型放在前面，优先匹配更具体的错误类型
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
		log.Fatalln(err)
	}
	println(z)
}

// panic/recover 类似 try/catch
func panicRecover1() {
	defer func() {
		if err := recover(); err != nil { // 捕获错误
			log.Fatalln(err)
		}
	}()

	panic("i am dead") // 引发错误
	println("exit.")   // 永不会执行
}

// 中断性错误会沿调用堆栈向外传递，要么被外层捕获，要么导致进程崩溃
func panicRecover2() {
	test := func() {
		defer println("test.1")
		defer println("test.2")

		panic("i am dead")
	}

	defer func() {
		log.Println(recover())
	}()

	test()
}

// 连续调用 panic，仅最后一个会被 recover 捕获
func panicRecover3() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal")
			}
		}
	}()

	defer func() {
		panic("you are dead") // 类似重新抛出异常 rethrow
	}() // 可先 recover 捕获，包装后重新抛出

	panic("i am dead")
}

// recover 必须在延迟调用函数中执行才能正常工作
func panicRecover4() {
	catch := func() {
		log.Println("catch:", recover())
	}

	defer catch()                // 捕获
	defer log.Println(recover()) // 失败！
	defer recover()              // 失败！

	panic("i am dead")
}

func panicRecover5() {
	test := func(x, y int) {
		z := 0

		func() { // 利用匿名函数保护 "z = x / y"
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
					z = 0
				}
			}()

			z = x / y
		}()

		println("x / y =", z)
	}

	test(5, 0)
}

func panicRecoverDebug() {
	test := func() {
		panic("i am dead")
	}

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack() // 输出完整调用堆栈信息
		}
	}()

	test()
}

func main() {
	errorInit()
	errorStruct()
	panicRecover1()
	panicRecover2()
	panicRecover3()
	panicRecover4()
	panicRecover5()
	panicRecoverDebug()
}

/*
	建议：除非是不可恢复性、导致系统无法正常工作的错误，否则不建议使用 panic
*/
