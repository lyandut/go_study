package main

import "fmt"

type X int

// 为X类型定义方法
func (x *X) inc() { // 名称前的参数称作 receiver，作用类似 python self
	*x++
}

// 直接调用匿名字段的方法，实现与继承类似的功能
type mUser struct {
	name string
	age  byte
}

func (u mUser) ToString() string {
	return fmt.Sprintf("%+v", u) // %+v: 输出结构体，会添加字段名
}

type mManager struct {
	mUser
	title string
}

func main() {
	var x X
	x.inc()
	println(x)

	var m mManager
	m.name = "Tom"
	m.age = 29
	println(m.ToString()) // 调用 m_user.ToString()
}
