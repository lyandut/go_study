package main

import (
	"errors"
	"fmt"
)

func returnMultiple() {
	div := func(x, y int) (int, error) { // 多返回值列表必须使用括号
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}

	test := func() (int, error) {
		return div(5, 0) // 多返回值用作 return 结果
	}

	log := func(x int, err error) {
		fmt.Println(x, err)
	}

	log(test()) // 多返回值用作实参
}

func returnRenamedDiv(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
		return
	}
	z = x / y
	return // 隐式返回，相当于 "return z, err"
}

func returnRenamedAdd(x, y int) (z int) { // 命名返回值也必须使用括号
	z = x - y
	{
		z := x + y // 新定义的同名局部变量，同名遮蔽
		return z   // 显示 return
	}
	return
}

func main() {
	returnMultiple()
	fmt.Println(returnRenamedDiv(5, 0))
	fmt.Println(returnRenamedAdd(1, 2))
}

/*
	PS: 必须对全部返回值命名，否则编译器会搞不清状况
		如果返回值类型能明确表明其含义，就尽量不要对其命名
*/
