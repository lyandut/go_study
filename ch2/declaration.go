package main

import "os"

// 变量覆盖/幽灵变量：不同作用域，全部是新变量定义
func declarationShadows() {
	x := 100
	println(&x, x)

	{
		x, y := 200, 300
		println(&x, x, y)
	}
}

// 退化赋值：至少有一个新变量被定义，且必须是同一作用域
func declarationDegenerate() {
	x := 100
	println(&x)

	x, y := 200, "abc" // x退化为赋值操作，仅有y是变量定义
	println(&x, x)
	println(y)

	f, err := os.Open("/dev/random")
	buf := make([]byte, 1024)
	n, err := f.Read(buf) // 退化赋值允许重复使用err变量
	println(n, err)
}

func main() {
	declarationShadows()
	declarationDegenerate()
}
