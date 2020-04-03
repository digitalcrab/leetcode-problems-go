package problems

// In a given grid, each cell can have one of three values:
//
// the value 0 representing an empty cell;
// the value 1 representing a fresh orange;
// the value 2 representing a rotten orange.
// Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
//
// Return the minimum number of minutes that must elapse until no cell has a fresh orange.
// If this is impossible, return -1 instead.

// Example 1:
// Input: [[2,1,1],[1,1,0],[0,1,1]]
// Output: 4

// Example 2:
// Input: [[2,1,1],[0,1,1],[1,0,1]]
// Output: -1
// Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten,
// because rotting only happens 4-directionally.

// Example 3:
// Input: [[0,2]]
// Output: 0
// Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.

// Note:
//
// 1 <= grid.length <= 10
// 1 <= grid[0].length <= 10
// grid[i][j] is only 0, 1, or 2.

func OrangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	deltaRows := []int{-1, 0, 1, 0}
	deltaCols := []int{0, -1, 0, 1}

	// in the worst case we enqueue all the items from the grid
	queue := make([]int, 0, rows*cols)
	// we store depth of the cell as idx => depth
	depth := make(map[int]int, rows)

	var answer int

	// find all rotten sports, complexity is O(n) - we go through every item on the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				idx := r*cols + c // sequential number
				queue = append(queue, idx)
				depth[idx] = 0 // initial day
			}
		}
	}

	// rotten the rest and increment days
	// worst case O(n)
	for len(queue) > 0 {
		var idx int
		idx, queue = queue[0], queue[1:] // get first and shrink the queue from the beginning

		// getting back row and col from sequential number
		r, c := idx/cols, idx%cols

		for step := 0; step < 4; step++ {
			newR := r + deltaRows[step]
			newC := c + deltaCols[step]

			// check boundaries: 1) outside of greed; 2) nd has fresh fruit
			if newR < 0 || newC < 0 || newR >= rows || newC >= cols || grid[newR][newC] != 1 {
				continue
			}

			grid[newR][newC] = 2 // rotten
			newIdx := newR*cols + newC

			// enqueue new rotten fruit
			queue = append(queue, newIdx)
			// increase depth
			answer = depth[idx] + 1
			depth[newIdx] = answer
		}
	}

	// check if at least one is still fresh
	// O(n)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return answer
}
