package main

import "fmt"

//type StackElem struct {
//	temp int
//	nextCount int
//}
//
//type Stack []StackElem
//
//func Factory() *Stack {
//	stack := &Stack{}
//	return stack
//}
//
//func (s *Stack)Push(elem StackElem) {
//	*s = append(*s, elem)
//}


func dailyTemperatures(T []int) []int {

	stack := make([]int,0)
	rst := make([]int, len(T))

	for i:=len(T)-1; i>=0; i-- {
		for len(stack)>0 && T[i]>T[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			rst[i]=0
		} else {
			rst[i] = stack[len(stack)-1]-i
		}

		stack = append(stack, i)

	}

	return rst
}

func main() {
	a := []int{73,74,75,71,69,72,76,73}

	fmt.Println(dailyTemperatures(a))
}