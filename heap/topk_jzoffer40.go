package heap


func GetLeastNumbers(arr []int, k int) []int {
	return getLeastNumbers(arr,k)

}

func getLeastNumbers(arr []int, k int) []int {

	BuildHeap(arr, k) //construct the big top heap with length k
	for i:=k; i<len(arr); i++ {
		if arr[i] < arr[0] {
			arr[0] = arr[i]
			HeapAdjust(arr, 0, k)
		}
	}

	return arr[:k]
}



func BuildHeap(arr []int, length int) {

	for i:=length/2-1; i>=0; i-- {
		HeapAdjust(arr, i, length)
	}
}

func HeapAdjust(arr []int, i int, length int) {

	max := i
	leftChild := 2*i+1
	rightChild := 2*i+2

	if leftChild<length && arr[max] < arr[leftChild] {
		max = leftChild
	}

	if rightChild<length && arr[max] < arr[rightChild] {
		max = rightChild
	}

	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		HeapAdjust(arr, max, length)
	}
}

