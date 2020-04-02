package algorithms

func binarySearchIterative(array []int, x int) bool {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + ((right - left) / 2)

		if array[mid] == x {
			return true
		}

		if x < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

// Time Complexity: O(log(n))
func binarySearchRecursive(array []int, x int) bool {
	var search func(array []int, x, left, right int) bool

	search = func(array []int, x, left, right int) bool {
		if left > right {
			return false
		}

		// mid := (right + left) / 2
		// ^ right + left could be potentially bigger than max int

		// this way we ensure that mid wont exceed the limit
		mid := left + ((right - left) / 2)
		item := array[mid]
		if item == x {
			return true
		}
		if x < array[mid] {
			return search(array, x, left, mid-1)
		}
		return search(array, x, mid+1, right)
	}
	return search(array, x, 0, len(array)-1)
}
