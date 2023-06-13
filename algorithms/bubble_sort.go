package algorithms

func BubbleSortNaive(nums []int) {
	l := len(nums)
	// repeat as many times as we have elements, with the hope that on every iteration we get
	// closer to the sorted array
	for range nums {
		for i := 1; i < l; i++ {
			if left, right := nums[i-1], nums[i]; left > right {
				nums[i-1], nums[i] = right, left
			}
		}
	}
}

func BubbleSortOptimized(nums []int) {
	l := len(nums)
	// repeat as many times as we have elements minus the times we've done that
	// since on every iteration we move 1 element to the end (biggest)
	for j := range nums {
		for i := 1; i < l-j; i++ {
			if left, right := nums[i-1], nums[i]; left > right {
				nums[i-1], nums[i] = right, left
			}
		}
	}
}
