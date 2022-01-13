package main

func gotoInit() {
	for i := 0; i <= 3; i++ {
		println(i)
		if i > 1 {
			goto exit
		}
	}
exit:
	println("exit.")
}

//func gotoError() {
//	tester := func() {
//	test:
//		println("test")
//		println("test exit.")
//	}
//
//start: // 错误：未使用的标签
//	for i := 0; i < 3; i++ {
//	loop:
//		println(i)
//	}
//
//	goto test // 不能跳转到其他函数
//	goto loop // 不能跳转到内层代码块内
//}

func continueBreak() {
outer: // 配合标签，continue 和 break 可在多层嵌套中指定目标层级
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				println()
				continue outer
			}

			if x > 2 {
				break outer
			}

			print(x, ":", y, " ")
		}
	}
}

func main() {
	gotoInit()
	continueBreak()
}
