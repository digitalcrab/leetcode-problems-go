package main

func generateParenthesis(n int) (out []string) {
	// empty on empty
	if n == 0 {
		return
	}
	// with 1 value there is only one way to have parenthesis
	if n == 1 {
		return []string{"()"}
	}

	// lets describe the step
	type step struct {
		opened, closed int
		str            string
	}

	// start the queue with nothing open and nothing closed
	queue := []step{{opened: 0, closed: 0}}

	// and loop until we have nothing in the stack
	for len(queue) > 0 { // Time: O(n**n)
		// pick the first in the queue
		current := queue[0]
		queue = queue[1:]

		// check if we reached the limit
		if current.opened == n && current.closed == n {
			out = append(out, current.str)
			continue
		}

		// evaluate if we can close parenthesis
		// if opened more than closed and closed in total no more than we can
		if current.opened > current.closed && current.closed < n {
			queue = append(queue, step{
				opened: current.opened,
				closed: current.closed + 1,
				str:    current.str + ")",
			})
		}

		// evaluate if we can open, and we can do it if we haven't opened all of them
		if current.opened < n {
			queue = append(queue, step{
				opened: current.opened + 1,
				closed: current.closed,
				str:    current.str + "(",
			})
		}
	}

	return
}
