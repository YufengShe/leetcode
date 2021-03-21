package main

import "fmt"

func main() {

	a := make([]int, 0)

	b := a

	for i:=0; i<30; i++ {
		b = append(b,1)
	}

	fmt.Println(a, b)



}

