package problems

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(res, []int{0, 1}) {
		t.Errorf("unexpected result %v", res)
	}
}
