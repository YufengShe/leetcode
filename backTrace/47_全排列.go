package main

import "fmt"
var rst [][]int = [][]int{}
var chose []int = []int{}
var visited []int
var num []int
func permuteUnique(nums []int) [][]int {
	num = nums
	visited = make([]int, len(nums))
	sort(nums)

	//var backTrace func()
	backTrace()
	return rst
}

func sort(arr []int) {

	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}


func backTrace() {
	if len(chose) == len(num) {
		temp := make([]int, len(num))
		copy(temp, chose)
		rst = append(rst, temp)
	}

	for i:=0; i<len(num); i++ {
		if visited[i] == 1 {
			continue
		}

		//if i>0 && visited[i-1]==0 && num[i] == num[i-1] {
		//	continue
		//}
		//choose
		chose = append(chose, num[i])
		visited[i] = 1

		//digui
		backTrace()

		//recover choose
		chose = chose[:len(chose)-1]
		visited[i] = 0
		if i+1<len(num) && num[i+1]==num[i] {
			i++
		}

	}

}
func main() {
	nums := []int{1,1,2}
	fmt.Println(permuteUnique(nums))
}
