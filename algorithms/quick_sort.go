package algorithms

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	partition := quickSortPartition(nums, left, right)

	quickSort(nums, left, partition-1)
	quickSort(nums, partition+1, right)
}

func quickSortPartition(nums []int, left int, right int) int {
	// chose pivot (randomly, in the middle, or at the any end)
	pivot := nums[right]
	i := left

	// loop over chunk
	for j := left; j < right; j++ {
		// if current value is lower than pivot we have to swap the place
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// place pivot on it's position (right it's pivot in this case)
	nums[i], nums[right] = nums[right], nums[i]

	return i
}
