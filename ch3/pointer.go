package main

func pointerAddress() {
	x := 10
	var p *int = &x
	*p += 20
	println(p, *p)

	// m := map[string]int{"a": 1}
	// println(&m["a"]) // Cannot take the address of 'm["a"]'
}

func pointerStruct() {
	a := struct {
		x int
	}{}
	a.x = 100
	p := &a
	p.x += 10 // 相当于 p->x += 10
	println(p.x)

	// zero-size 对象的地址是否相等和具体实现的版本有关，
	// 不过肯定不等于 nil
	var b, c struct{}
	println(&b, &c)              // 0xc000032768 0xc000032768
	println(&b == &c, &b == nil) // false false
}

func main() {
	pointerAddress()
	pointerStruct()
}
