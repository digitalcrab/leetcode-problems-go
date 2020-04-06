package problems

// https://leetcode.com/problems/number-of-islands/
//
// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
//
// Example 1:
// Input:
//   11110
//   11010
//   11000
//   00000
// Output: 1
//
// Example 2:
// Input:
//   11000
//   11000
//   00100
//   00011
// Output: 3

func numIslands(grid [][]byte) int {
	return numIslandsBreadthFirstSearch(grid)
}

func numIslandsBreadthFirstSearch(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	// prepare steps
	diffRows := [4]int{1, 0, -1, 0}
	diffCols := [4]int{0, 1, 0, -1}

	rows := len(grid)
	cols := len(grid[0])
	var islands int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// lets say we found this first node and going to use it as tree root
			if grid[r][c] == '1' {
				islands++
				// mark as visited
				grid[r][c] = '0'

				// create id
				currentId := r*cols + c

				// create queue of neighbours
				queue := make([]int, 0, rows*cols)
				// add current cell into the queue
				queue = append(queue, currentId)

				for len(queue) > 0 {
					var idx int
					// dequeue item and resize the queue the
					idx, queue = queue[0], queue[1:]

					// get the neighbour cell
					neighbourR := idx / cols
					neighbourC := idx % cols

					// move step by step in 4 directions
					for step := 0; step < 4; step++ {
						newR := neighbourR + diffRows[step]
						newC := neighbourC + diffCols[step]

						if newR >= 0 && newC >= 0 && newR < rows && newC < cols && grid[newR][newC] == '1' {
							// mark as visited
							grid[newR][newC] = '0'
							// schedule into the queue
							queue = append(queue, newR*cols+newC)
						}
					}
				}
			}
		}
	}

	return islands
}

//
func numIslandsDepthFirstSearch(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	var islands int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// lets say we found this first node and going to use it as tree root
			if grid[r][c] == '1' {
				islands++
				numIslandsDFS(grid, r, c)
			}
		}
	}

	return islands
}

func numIslandsDFS(grid [][]byte, r, c int) {
	// out of boundaries of not a part of island (or visited before)
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}

	// mark as visited
	grid[r][c] = '0'

	numIslandsDFS(grid, r+1, c)
	numIslandsDFS(grid, r-1, c)
	numIslandsDFS(grid, r, c+1)
	numIslandsDFS(grid, r, c-1)
}
