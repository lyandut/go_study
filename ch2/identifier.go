package main

import "strconv"

func main() {
	x, _ := strconv.Atoi("12") // 忽略 Atoi 的 err 返回值
	println(x)
}
