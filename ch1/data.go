package main

import "fmt"

// 切片（slice）==> 动态数组
func dataSlice() {
	x := make([]int, 0, 5) // 创建size=0, capacity=5的切片
	for i := 0; i < 8; i++ {
		x = append(x, i) // 追加数据。当超出capacity限制时，自动分配更大的存储空间
	}
	fmt.Println(x)
	fmt.Println(len(x), cap(x)) // 扩容机制: len<1024, cap*2

	y := [3]int{1} // 创建长度为3的数组，默认初值为0
	y[1] = 2
	fmt.Println(y)
}

func dataMap() {
	m := make(map[string]int) // 创建字典对象
	m["a"] = 1                // 添加或设置
	m["b"] = 2
	x, ok := m["b"] // 使用ok-idiom获取值，可知道 key/value 是否存在
	fmt.Println(x, ok)
	delete(m, "a") // 删除
	fmt.Println(m)
}

func dataStruct() {
	type user struct {
		name string
		age  byte
	}

	type manager struct {
		user  // 匿名嵌入其他类型
		title string
	}

	var m manager
	m.name = "Tom"  // 直接访问匿名字段成员
	m.user.age = 29 // 使用匿名类型访问成员
	m.title = "CTO"
	fmt.Println(m)
}

func main() {
	dataSlice()
	dataMap()
	dataStruct()
}
