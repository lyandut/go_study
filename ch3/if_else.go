package main

import (
	"errors"
	"log"
	"strconv"
)

func ifElseInit() {
	x := 10
	if x = 0; x == 0 { // 修改 x 变量
		println("x == 0")
	}
	if a, b := x+1, x+10; a < b { // 定义一个或多个局部变量（也可以是函数返回值）
		println(a)
	} else {
		println(b) // 局部变量的有效范围包含整个 if/else 块
	}
}

func ifElseCheck1() {
	s := "9"
	n, err := strconv.ParseInt(s, 10, 64) // 使用外部变量
	if err != nil {
		log.Fatalln(err)
	} else if n < 0 || n > 10 {
		log.Fatalln("invalid number")
	}
	println(n)
}

func ifElseCheck2() {
	check := func(s string) error { // 将过于复杂的组合条件重构为函数
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil || n < 0 || n > 10 || n%2 != 0 {
			return errors.New("invalid number")
		}
		return nil
	}

	s := "9"
	if err := check(s); err != nil {
		log.Fatalln(err)
	}
	println("ok")
}

func main() {
	ifElseInit()
	ifElseCheck1()
	ifElseCheck2()
}
