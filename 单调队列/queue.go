package main

import (
	"fmt"
	"strconv"
)

func maxSlidingWindow(nums []int, k int) []int {

	win := make([]int, 0)
	rst := make([]int, 0)

	for i:=0; i<len(nums); i++ {
		if i<k-1 {
			for len(win)>0 && nums[i]>win[len(win)-1] {
				win = win[:len(nums)-1]  //pop_tail
			}
			win = append(win, nums[i])
		} else {
			for len(win)>0 && nums[i]>win[len(win)-1] {
				win = win[:len(nums)-1]
			}
			win = append(win, nums[i])

			rst = append(rst, win[0])

			if win[0] == nums[i-k+1] {
				win = win[1:]
			}
		}


	}

	return rst
}

func main() {

	s := "-123"
	num, err := strconv.Atoi(s)
	fmt.Println(num, err)
}
