package problems

// https://leetcode.com/problems/flood-fill/
//
// An image is represented by a 2-D array of integers,
// each integer representing the pixel value of the image (from 0 to 65535).
//
// Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill,
// and a pixel value newColor, "flood fill" the image.
//
// To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally
// to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally
// to those pixels (also with the same color as the starting pixel), and so on.
// Replace the color of all of the aforementioned pixels with the newColor.
//
// At the end, return the modified image.
//
// Example 1:
// Input:
// image = [[1,1,1],[1,1,0],[1,0,1]]
// sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation:
//  From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
//  by a path of the same color as the starting pixel are colored with the new color.
//  Note the bottom corner is not colored 2, because it is not 4-directionally connected
//  to the starting pixel.
//
// Note:
//  - The length of image and image[0] will be in the range [1, 50].
//  - The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
//  - The value of each color in image[i][j] and newColor will be an integer in [0, 65535].
//

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	rows := len(image)
	cols := len(image[0])

	// coordinates delta in 4 directions
	deltaRows := []int{-1, 0, 1, 0}
	deltaCols := []int{0, -1, 0, 1}

	// in the worst case we enqueue all the points from the image
	queue := make([]int, 0, rows*cols)

	// create an unique index (sequential number) for the starting point
	startingIdx := sr*cols + sc
	// enqueue starting point
	queue = append(queue, startingIdx)
	// store initial color
	color := image[sr][sc]
	// change color
	image[sr][sc] = newColor

	// start traversing the image from the starting point
	for len(queue) > 0 {
		var idx int
		idx, queue = queue[0], queue[1:]

		// getting back row and col from sequential number
		r, c := idx/cols, idx%cols

		for step := 0; step < 4; step++ {
			newR, newC := r+deltaRows[step], c+deltaCols[step]
			newIdx := newR*cols + newC

			// if we are outside of boundaries and the color is not the one we wold like to change
			if newR < 0 || newC < 0 || newR >= rows || newC >= cols || image[newR][newC] != color {
				continue
			}

			// change color
			image[newR][newC] = newColor

			// enqueue the next starting point
			queue = append(queue, newIdx)
		}
	}

	return image
}
