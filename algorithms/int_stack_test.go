package algorithms

import (
	"testing"
)

func TestStack(t *testing.T) {
	st := &intStack{}
	st.push(1)
	st.push(2)
	st.push(3)

	if v := st.pop(); v != 3 {
		t.Errorf("unexpected result %d -> %d", v, 3)
	}
	if v := st.pop(); v != 2 {
		t.Errorf("unexpected result %d -> %d", v, 2)
	}
	if v := st.pop(); v != 1 {
		t.Errorf("unexpected result %d -> %d", v, 1)
	}
}
