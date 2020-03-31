package problems

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 22)
	if !reflect.DeepEqual(res, []int{1, 3}) {
		t.Errorf("unexpected result %v", res)
	}
}
