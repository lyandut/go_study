package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1. 展开 & 显式类型转换
	const a = 100            // 无类型声明的产量
	const b byte = a         // 直接展开a，相当于 const b byte = 100
	const aa int = 100       // 显式指定常量类型，编译器会做强类型检查
	const bb byte = byte(aa) // 显式类型转换

	// errors
	// const n = byte(-100)  // constant -100 overflows byte
	// println(&a, a)        // cannot take the address of a

	// 2. 常量值可以是某些编译器能计算出结果的表达式，如 unsafe.Sizeof、len、cap 等
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello world!")
	)

	// 3. 在常量组中如不指定类型和初始值，则与上一行非空常量右值（表达式文本）相同
	const (
		x uint16 = 120
		y
		s = "abc"
		z
	)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)

	// 4. 枚举，借助 iota 标识符实现一组自增常量值
	type mem uint // 建议用自定义类型实现用途明确的枚举类型
	const (
		_      = iota             // 0
		KB mem = 1 << (10 * iota) // 1 << (10 * 1)
		MB                        // 1 << (10 * 2)
		GB                        // 1 << (10 * 3)
	)
	println(KB, MB, GB)
}
