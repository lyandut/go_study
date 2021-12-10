package main

// 引用类型（reference type）特指 slice、map、channel 这三种预定义类型
// new/make 都不是 go 关键字，而是 go 预定义函数

// 引用类型必须使用 make 创建，完成全部内存分配和相关属性初始化
func mkSlice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkMap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

// new 按指定类型长度分配零值内存，返回指针，但不关心类型内部构造和初始化，是不完整创建
func newInt() int {
	i := new(int)
	*i = 2
	return *i
}

func newSlice() []int {
	s := new([]int)
	// (*s)[0] = 100 // panic: runtime error: index out of range [0] with length 0
	*s = append(*s, 100)
	return *s
}

func newMap() map[string]int {
	m := new(map[string]int)
	(*m)["a"] = 1 // panic: assignment to entry in nil map
	return *m
}

func main() {
	s1 := mkSlice()
	println(s1[0])

	m1 := mkMap()
	println(m1["a"])

	i := newInt()
	println(i)

	s2 := newSlice()
	println(s2[0])

	m2 := newMap()
	println(m2["a"])
}
