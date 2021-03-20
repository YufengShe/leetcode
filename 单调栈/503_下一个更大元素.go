package main


func nextGreaterElements(nums []int) []int {

	rst := make([]int, len(nums))
	stack1 := make([]int, 0) //next bigger
	//stack2 := make([]int, 0) //previous bigger

	//next
	n := len(nums)-1
	for i:=2*n-1; i>=0; i-- {

		for len(stack1)>0 && nums[i%n]>=stack1[len(stack1)-1] {
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack1) <=0 {
			rst[i%n] = -1
		} else {
			rst[i%n] = stack1[len(stack1)-1]
		}

		stack1 = append(stack1, nums[i%n])
	}

	return rst

}