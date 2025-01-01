package main

func isPalindrome(s string) bool {
	// if empty of has only 1 char
	if len(s) < 2 {
		return true
	}

	// let's create 2 pointers that points to the beginning of the string and end
	left, right := 0, len(s)-1

	// let's start the loop over the string from both ends
	for left <= right {
		// skip all chars from the left and right that are not alphanumeric
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		// now we must compare the characters
		// if they are not equal, then it's not a palindrome
		if lower(s[left]) != lower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}
