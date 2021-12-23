package main

import "fmt"

func operatorShift() {
	a := 1.0 << 3 // 常量表达式（包括常量展开）
	fmt.Printf("%T, %v\n", a, a)

	var s uint = 3
	//b := 1.0 << s                // invalid operation: 1 << s (shift of type float64)
	//fmt.Printf("%T, %v\n", b, b) // 编译器推断 b 为浮点数类型

	var c int32 = 1.0 << s // 自动将 1.0 转换为 int32 类型
	fmt.Printf("%T, %v\n", c, c)
}

func operatorBitClear() {
	const (
		read byte = 1 << iota
		write
		exec
		freeze
	)

	a := read | write | freeze
	b := read | freeze | exec
	c := a &^ b // 相当于 a ^ read ^ freeze，但不包括 exec，根据b清除a的标记位
	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
}

func main() {
	operatorShift()
	operatorBitClear()
}
