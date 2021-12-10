package main

import (
	"fmt"
	"reflect"
)

func unnamedStruct() {
	var a struct { // 匿名结构类型
		x int    `x` // struct tag 也属于类型组成部分，
		s string `s` // 而不仅仅是元数据的描述
	}

	var b struct {
		x int
		s string
	}

	fmt.Printf("a type:%T\n", a)
	fmt.Printf("b type:%T\n", b)
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
}

func unnamedFunction() {
	var a func(int, string) // 函数的参数顺序也属于签名组成部分
	var b func(string, int)

	fmt.Printf("a type:%T\n", a)
	fmt.Printf("b type:%T\n", b)
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
}

func unnamedConversion() {
	type data [2]int
	var d data = [2]int{1, 2} // 基础类型相同，右值为未命名类型
	fmt.Println(d)

	a := make(chan int, 2)
	var b chan<- int = a // 双向通道转换为单向通道，其中b为未命名类型
	b <- 2
}

func main() {
	unnamedStruct()
	unnamedFunction()
	unnamedConversion()
}
