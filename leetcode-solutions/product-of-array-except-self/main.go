package main

import (
	"fmt"
	"math"
)

func productExceptSelf(nums []int) (out []int) {
	if len(nums) < 2 {
		panic("unexpected") // based on the constraints
	}

	// create the same size array for the output, default to all 0
	out = make([]int, len(nums))

	// store id of a single element with a value 0
	zeroId := -1
	// store the product of all elements
	productWithOutZero := 1

	for i, n := range nums { // Time: O(n)
		if n == 0 {
			// if more than one zero then everything is zero as the result
			// makes no sense to calculate
			if zeroId != -1 {
				return out
			}
			zeroId = i
		} else {
			productWithOutZero *= n // calculate product without zeros
		}
	}

	// if there is only one 0 in the array everything will be zero except itself
	if zeroId != -1 {
		out[zeroId] = productWithOutZero
		return
	}

	// here is the normal case
	for i, n := range nums { // Time: O(n)
		// we need to divide to the current (to exclude it from the product)
		out[i] = divide(productWithOutZero, n)
	}

	return
}

func divide(dividend, divisor int) int {
	// the easy way ;)
	// return dividend / divisor

	// special case if both are the same
	if dividend == divisor {
		return 1
	}

	// keep the sign of the result (it's minus only if one of the parts is minus)
	minus := (dividend < 0 && divisor >= 0) || (divisor < 0 && dividend >= 0)

	// from now on we operate only with absolute values
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))

	var answer int

	// keep going while we have something to divide
	for dividendAbs >= divisorAbs { // Time: O(M/N), where M is bits length of dividend and N, bits length of divisor
		// find out how many times we can divide `dividend`
		var cnt int
		// with `<<` we multiply on 2 ** (cnt + 1)
		// so ```divisorAbs << (cnt + 1)```  is ```divisor * (2 ** (cnt + 1))```
		for dividendAbs >= (divisorAbs << (cnt + 1)) {
			/* Time Complexity: O(log(N)), where n is divisor,
			   log because we reduce the problem with every iteration */
			cnt++
		}

		// now we know how much we can take from it

		// let's add to the answer first, that would be how maybe times 2 ** cnt, we were able to do thing
		answer += 1 << cnt

		// reduce dividend to the amount of divisor * (2 ** cnt)
		dividendAbs -= divisorAbs << cnt
	}

	// do not forget the sign
	if minus {
		return -answer
	}

	return answer
}

func main() {
	answer := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Printf("The answer is: %v\n", answer) // [24,12,8,6]
}
