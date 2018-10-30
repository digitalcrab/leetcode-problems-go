package problems

import "math"

// Given a 32-bit signed integer, reverse digits of an integer.
//
// Time Complexity: O(log(x)). There are roughly log10(x) digits in x, We reduce a number of iterations by 10
// Space Complexity: O(1)
//
func ReverseInteger(x int) int {
	res := 0

	// loop over example 123
	for x != 0 {
		// 1) 123 % 10 = 3; 2) 12 % 10 = 2; 3) 1 % 10 = 1
		mod := x % 10

		// 1) 123 / 10 = 12; 2) 12 / 10 = 1; 3) 1 / 10 = 0
		x = x / 10

		// before we multiply by 10 we have to check the current value by dividing max int by 10

		// 2147483647 / 10 = 214748364 and only 7 left
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && mod > 7) {
			return 0
		}

		// -2147483648 / 10 = -214748364 and only -8 left
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && mod < -8) {
			return 0
		}

		// 1) 0 * 10 + 3 = 3; 2) 3 * 10 + 2 = 32; 3) 32 * 10 + 1 = 321
		res = res*10 + mod
	}

	return res
}
