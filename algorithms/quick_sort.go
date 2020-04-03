package algorithms

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	// pick randomly
	pivot := nums[left+(right-left)/2]
	// split array into parts
	idx := quickSortPartition(nums, left, right, pivot)

	// sort for every partition
	quickSort(nums, left, idx-1)
	quickSort(nums, idx, right)
}

func quickSortPartition(nums []int, left, right, pivot int) int {
	// moving to the middle point (or somewhere they meet each other)
	for left <= right {
		// keep moving from left and right simultaneously
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}

		// if we found element out of order we swat it
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	// return new partition point
	return left
}
