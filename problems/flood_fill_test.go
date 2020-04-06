package problems

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	cases := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		expected [][]int
	}{
		{
			image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:       1,
			sc:       1,
			newColor: 2,
			expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image:    [][]int{{0, 0, 0}, {0, 1, 1}},
			sr:       1,
			sc:       1,
			newColor: 1,
			expected: [][]int{{0, 0, 0}, {0, 1, 1}},
		},
		{
			image:    [][]int{{0, 0, 1}, {0, 1, 1}},
			sr:       1,
			sc:       1,
			newColor: 1,
			expected: [][]int{{0, 0, 1}, {0, 1, 1}},
		},
	}

	for i, c := range cases {
		// copy the image
		image := make([][]int, len(c.image))
		copy(image, c.image)

		if res := floodFill(image, c.sr, c.sc, c.newColor); !reflect.DeepEqual(res, c.expected) {
			t.Errorf("#%d, unexpected result %v of %v (%d %d %d) -> %v", i, res, c.image, c.sr, c.sc, c.newColor, c.expected)
		}
	}
}
