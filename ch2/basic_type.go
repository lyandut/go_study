package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
	fmt.Printf("%#b, %#o, %#x\n", a, b, c)
	fmt.Println(math.MinInt8, math.MaxInt8)

	aa, _ := strconv.ParseInt("1100100", 2, 32)
	bb, _ := strconv.ParseInt("144", 8, 32)
	cc, _ := strconv.ParseInt("64", 16, 32)
	println(aa, bb, cc)

	println("0b" + strconv.FormatInt(aa, 2))
	println("0" + strconv.FormatInt(bb, 8))
	println("0x" + strconv.FormatInt(cc, 16))
}
