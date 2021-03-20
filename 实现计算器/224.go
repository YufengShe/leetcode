package main
//
//import (
//	"strings"
//)
//
//func calculate(s string) int {
//
//	s = strings.Replace(s, " ", "", -1)
//	stack := []int{}
//
//	sign := 1
//	rst := 0
//	num := 0
//
//	for i:=0; i<len(s); i++ {
//		if s[i] == '-' {
//			rst  = rst + sign*num
//			num = 0
//			sign = -1
//		} else if s[i] == '+' {
//			rst  = rst + sign*num
//			num = 0
//			sign = 1
//		}else if s[i] == '(' {
//			stack = append(stack, rst)
//			stack = append(stack, sign)
//			rst = 0
//			sign = 1
//		} else if s[i]>='0' && s[i]<='9'{
//			num = num*10 + int(s[i]-'0')
//		} else if s[i] == ')' {
//			rst = rst + num*sign
//			rst = rst + stack[len(stack)-1]*stack[len(stack)-2]
//			stack = stack[:len(stack)-2]
//
//			num = 0
//			sign = 1
//
//		}
//	}
//
//	rst = rst +num*sign
//
//	return rst
//
//}
//
