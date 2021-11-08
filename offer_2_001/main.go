package main

import (
	"fmt"
	"math"
)

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	if a > 0 {
		a = ^a + 1 //取反
	}
	if b > 0 {
		b = ^b + 1 //取反
	}
	if sign == 1 {
		return divideCore(a, b)
	} else {
		return ^divideCore(a, b) + 1 //取反
	}
}

//除数被除数转换成负数后的除法
func divideCore(a, b int) int {
	rst := 0
	for a <= b {
		cnt := 1
		temp := b
		for temp >= math.MinInt32<<2 && a <= temp<<1 {
			cnt = cnt << 1
			temp = temp << 1
		}
		a = a - temp
		rst += cnt
	}
	return rst
}

func main() {
	fmt.Println(divide(15, 2))
}
