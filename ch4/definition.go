package main

// 函数属于第一类对象（first-class object）
// 第一类对象指可在运行期创建，可用作函数参数或返回值，可存入变量的实体。最常见的用法就是匿名函数
func hello() {
	println("hello world!")
}

func exec(f func()) {
	f()
}

// FormatFunc 定义函数类型，方便阅读和代码维护
type FormatFunc func(string, ...interface{}) (string, error)

func format(f FormatFunc, s string, a ...interface{}) (string, error) {
	return f(s, a...)
}

// 从函数返回局部变量指针是安全的，编译器会通过逃逸分析（escape analysis）来决定是否在堆上分配内存。
func escape() *int {
	a := 0x100
	return &a
}

func main() {
	f := hello
	exec(f)

	var a *int = escape()
	println(a, *a)
}

/*
   $ go build -gcflags "-l -m" definition.go    // 禁用函数内联，输出优化信息
		moved to heap: a
   $ go tool objdump -s "main\.main" definition // 反汇编确认
		CALL main.escape(SB)
   $ go build -gcflags "-m" definition.go       // 默认优化方式，允许内联
		inlining call to escape
		moved to heap: a 						// Go 1.17.5，与书中不同，还是分配在堆上
   $ go tool objdump -s "main\.main" definition // 反汇编确认
		CALL runtime.printpointer(SB)
*/
