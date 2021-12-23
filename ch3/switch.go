package main

func switchInit() {
	switch x := 5; x { // switch 支持初始化语句
	default: // 编译器确保不会先执行 default 块
		x += 100
		println(x)
	case 5, 6: // 多个匹配条件命中其一即可（OR)
		x += 50
		println(x)
	}
}

func switchFallthrough() {
	switch x := 5; x {
	default:
		println(x)
	case 5:
		x += 10
		println(x)

		fallthrough // 继续执行下一 case，但不再匹配条件表达式
	case 6:
		x += 20
		println(x)

		// fallthrough // 如果在此继续 fallthrough，不会执行 default，完全按照源码顺序执行
		// 错误：不能在 'switch' 语句的 final case 中使用 'fallthrough'
	}
}

func switchBreak() {
	switch x := 5; x {
	case 5:
		x += 10
		println(x)

		if x >= 15 {
			break // 终止，不再执行后续语句
		}

		fallthrough // 必须是 case 块的最后一条语句，可使用 break 语句阻止
	case 6:
		x += 20
		println(x)
	}
}

func switch2If() {
	switch x := 5; { // 相当于 ``switch x:=5; true {...}``
	case x > 5:
		println("a")
	case x > 0 && x <= 5: // 不能写成 ``case x > 0, x <= 5``，因为多条件是 OR 关系
		println("b")
	default:
		println("z")
	}
}

func main() {
	switchInit()
	switchFallthrough()
	switchBreak()
	switch2If()
}
