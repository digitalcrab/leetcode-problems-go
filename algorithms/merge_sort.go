package algorithms

// Time Complexity:  O(n log(n)) - always
// Space Complexity: O(n) - we have to copy an array
func MergeSort(nums []int) {
	mergeSort(nums, make([]int, len(nums)), 0, len(nums)-1)
}

func mergeSort(nums, temp []int, left, right int) {
	if left >= right {
		return
	}

	idx := left + (right-left)/2

	mergeSort(nums, temp, left, idx)
	mergeSort(nums, temp, idx+1, right)
	mergeHalves(nums, temp, left, right, idx)
}

func mergeHalves(nums, temp []int, leftStart, rightEnd, middle int) {
	leftEnd, rightStart := middle, middle+1
	left, right, index := leftStart, rightStart, leftStart

	for left <= leftEnd && right <= rightEnd {
		if nums[left] <= nums[right] {
			temp[index] = nums[left]
			left++
		} else {
			temp[index] = nums[right]
			right++
		}

		index++
	}

	// copy leftover from the left side
	for i := left; i < leftEnd+1; i++ {
		temp[index] = nums[i]
		index++
	}
	for i := right; i < rightEnd+1; i++ {
		temp[index] = nums[i]
		index++
	}

	// copy everything back
	for i := leftStart; i < rightEnd+1; i++ {
		nums[i] = temp[i]
	}
}
