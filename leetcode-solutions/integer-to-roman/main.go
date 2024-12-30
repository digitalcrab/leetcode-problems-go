package main

import (
	"math"
	"strings"
)

func intToRoman(num int) (out string) {
	if num > 3999 {
		panic("constraint violation")
	}

	// a small struct to store values of roman number
	type roman struct {
		one, five, ten string
	}

	var (
		// current order (from the end of the number
		order int
		// create roman numbers based on order
		orders = [4]roman{
			0: {"I", "V", "X"},
			1: {"X", "L", "C"},
			2: {"C", "D", "M"},
			3: {"M", "", ""},
		}
		values [4]string
	)

	// let's start the loop until the number if gone
	for num > 0 {
		// n is a single digit we are working with, from the end of the number
		n := num % 10
		num = num / 10
		r := orders[order]

		// store in a values the converted number, idx is from the end
		values[len(orders)-1-order] = romanDigit(n, r.one, r.five, r.ten)
		order++
	}

	return strings.Join(values[:], "")
}

// less efficient because of log and pow use
func intToRomanV2(num int) (out string) {
	if num > 3999 {
		panic("constraint violation")
	}

	// lets start the loop until the number if gone
	for num > 0 {
		// get how many digits in the number
		digits := int(math.Log10(float64(num)))
		// get the pow of 10 to understand what is the number we are dealing with
		p := int(math.Pow10(digits))
		// n is a single digit we are working with
		n := num / p
		// num now is the rest of the number
		num = num % p

		// at this point we know all we need to build the roman number

		if p == 1000 {
			// we assume that more than 3999 can not be
			// M = 1000
			out += romanDigit(n, "M", "", "")
		} else if p == 100 {
			// C = 100, D = 500, M = 1000
			out += romanDigit(n, "C", "D", "M")
		} else if p == 10 {
			// X = 10, L = 50, C = 100
			out += romanDigit(n, "X", "L", "C")
		} else {
			// I = 1, V = 5, X = 10
			out += romanDigit(n, "I", "V", "X")
		}
	}

	return
}

//  1. If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the
//     input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
//  2. If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following
//     symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following
//     subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
//  3. Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot
//     append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
//     Given an integer, convert it to a Roman numeral.
func romanDigit(n int, one, five, ten string) string {
	switch {
	case n == 5:
		return five
	case n == 4:
		return one + five // one minus five
	case n == 9:
		return one + ten // one minus ten
	case n > 5:
		return five + strings.Repeat(one, n-5) // five + one * count
	default:
		return strings.Repeat(one, n) // one * count
	}
}
