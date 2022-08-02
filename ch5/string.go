package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func stringInit() {
	s := "雨痕\x61\142\u0041"
	fmt.Printf("%s\n", s)
	fmt.Printf("% x, len: %d\n", s, len(s))
}

func stringNil() {
	var s string
	println(s == "")
	//println(s == nil) // cannot compare s == nil (mismatched types string and untyped nil)
}

func stringRaw() {
	s := `line\r\n,
	line 2`
	println(s)
}

func stringOperator() {
	s := "ab" + // 跨行时，加法操作符必须在上一行结尾
		"cd"
	println(s == "abcd")
	println(s > "abcd")
}

func stringSlice() {
	s := "abc"
	println(s[1])
	//println(&s[1]) // cannot take address of s[1] (value of type byte)

	s = "abcdefg"
	s1 := s[:3]
	s2 := s[1:4] // [start, end)
	s3 := s[2:]
	println(s1, s2, s3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

func stringFor() {
	s := "雨痕"

	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%d: [%c]\n", i, s[i])
	}

	for i, c := range s { // rune: Unicode字符
		fmt.Printf("%d: [%c]\n", i, c)
	}
}

func printDataPointer(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func stringConvert() {
	s := "hello world!"
	printDataPointer("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)
	printDataPointer("string to []byte, bs: %x\n", &bs)
	printDataPointer("[]byte to string, s2: %x\n", &s2)

	rs := []rune(s)
	s3 := string(rs)
	printDataPointer("string to []rune, rs: %x\n", &rs)
	printDataPointer("[]rune to string, s3: %x\n", &s3)
}

func stringConvertUnsafe() {
	toString := func(bs []byte) string {
		return *(*string)(unsafe.Pointer(&bs))
	}

	bs := []byte("hello world!")
	s := toString(bs)
	printDataPointer("bs: %x\n", &bs)
	printDataPointer("s : %x\n", &s)
}

func main() {
	stringInit()
	stringNil()
	stringRaw()
	stringOperator()
	stringSlice()
	stringFor()
	stringConvert()
	stringConvertUnsafe()
}
