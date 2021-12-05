package main

// if
func expressionIf() {
	x := 100
	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println("0")
	}
}

// switch
func expressionSwitch() {
	x := 100
	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}
}

// for
func expressionFor() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}

	x := 0
	for x < 5 { // 相当于 while (x < 5) {...}
		println(x)
		x++
	}

	x = 4
	for { // 相当于 while (true) {...}
		println(x)
		x--
		if x < 0 {
			break
		}
	}

	y := []int{100, 101, 102}
	for i, n := range y { // for...range可以返回索引
		println(i, ":", n)
	}
}

func main() {
	expressionIf()
	expressionSwitch()
	expressionFor()
}
