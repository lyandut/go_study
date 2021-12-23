package main

import "fmt"

type data struct {
	x int
	s string
}

func main() {
	var a data = data{1, "abc"}
	// var a data = {1, "abc"}  // 语法错误: unexpect { （缺类型标签）

	b := data{
		1,
		"abc",
	}

	c := []int{
		1,
		2}

	d := []int{1, 2,
		3, 4,
		5,
	}

	fmt.Println(a, b, c, d)
}
