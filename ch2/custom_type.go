package main

import "fmt"

func customFlags() {
	type flags byte
	const (
		read flags = 1 << iota
		write
		exec
	)
	f := read | exec
	fmt.Printf("%b\n", f) // 输出二进制标记位
}

func customUserEvent() {
	type ( // 组
		user struct { // 结构体
			name string
			age  uint8
		}

		event func(string) bool // 函数类型
	)

	u := user{"Tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}
	f("abc")
}

func main() {
	customFlags()
	customUserEvent()
}
