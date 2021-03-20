package main

import "fmt"

func main() {

	a := new(int)
	*a = 2

	b := new(int)
	*b = 2

	fmt.Println(a == b)





}

