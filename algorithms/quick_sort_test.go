package algorithms

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 4, 39, 7, 7, 11, 15, 9, 57, 0, 34}

	QuickSort(arr)

	if !reflect.DeepEqual(arr, []int{0, 2, 4, 7, 7, 9, 11, 15, 34, 39, 57}) {
		t.Errorf("unexpected result %v", arr)
	}
}
