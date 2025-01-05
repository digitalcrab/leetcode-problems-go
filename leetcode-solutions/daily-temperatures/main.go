package main

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))

	// store the index of the elem
	var stack []int

	for i, v := range temperatures {
		// if there is something in the stack that is smaller than the current temperature?
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			// element that is smaller than current
			smallerIdx := stack[len(stack)-1]
			// remove this elem from the stack
			stack = stack[:len(stack)-1]
			// update the result for smaller idx (that is the distance between current and smaller)
			answer[smallerIdx] = i - smallerIdx
		}

		// add current index to the stack for future analyses
		stack = append(stack, i)
	}

	return answer
}
