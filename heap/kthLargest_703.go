package heap

type KthLargest struct {
	SHeap []int //small top heap
	k int
}


func Constructor(k int, nums []int) KthLargest {

	heap := make([]int, k)
	for i:=0; i<k; i++ {
		heap[i] = -1<<31 //min
	}

	for i:=0; i<len(nums); i++ {
		if nums[i]>heap[0] {
			heap[0]=nums[i]
			SHeapAdjust(heap,i,k)
		}
	}

	return KthLargest{
		SHeap:heap,
		k:k,
	}

}


func SHeapAdjust(nums []int, i int, length int) {
	min := i
	left := 2*i+1
	right := 2*i+2

	if left<length && nums[left]<nums[min] {
		min = left
	}

	if right<length && nums[right]<nums[min] {
		min = right
	}

	if min != i {
		nums[min], nums[i] = nums[i], nums[min]
		SHeapAdjust(nums, min, length)
	}
}

func (this *KthLargest) Add(val int) int {
	if val > this.SHeap[0] {
		this.SHeap[0] = val
		SHeapAdjust(this.SHeap, 0, this.k)
	}

	return this.SHeap[0]
}