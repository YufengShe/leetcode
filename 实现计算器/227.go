package main

import (
	"fmt"
)

func calculate(s string) int {

	var compute func() int

	compute = func() int {
		if s == "" {
			return 0
		}

		stack := []int{}
		num := 0
		var sign uint8 = '+'

		for len(s) > 0 {
			x := s[0]
			s  = s[1:]
			if x>='0' && x<='9' {
				num = num*10 + int(x-'0')
			}

			if x == '(' {
				num = compute()
			}

			if x =='+'||x=='-'||x=='*'||x=='/'||len(s)==0 || x ==')'{
				if sign == '+' {
					stack = append(stack, num)
				} else if sign == '-' {
					stack = append(stack, -num)
				} else if sign == '*' {
					stack[len(stack)-1] = stack[len(stack)-1]*num
				} else if sign == '/' {
					stack[len(stack)-1] = stack[len(stack)-1]/num
				}

				num = 0
				sign = x

				if x == ')' {
					break
				}
			}

		}

		return sum(stack)


	}

	return compute()
}

func sum(stack []int) int {
	rst := 0
	for i:=0; i<len(stack); i++ {
		rst += stack[i]
	}
	return rst
}


func main () {
	s := "3+(2/(3 + 2 * 2))"
	fmt.Println(calculate(s))

}
