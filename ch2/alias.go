package main

func main() {
	// 别名无需类型转换，可直接赋值
	var a byte = 0x11
	var b uint8 = a
	println(a + b)

	// 64位平台上int和int64结构完全一致，但分属不同类型，须显式转换
	var x int = 100
	var y int64 = int64(x)
	println(x + int(y))
}
