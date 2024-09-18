package structures

import "testing"

func TestSinglyLikedList_Display(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		ll   *SinglyLikedList[T]
		want string
	}
	tests := []testCase[int]{
		{
			name: "empty list",
			ll:   NewSinglyLikedList[int](),
			want: "nil",
		},
		{
			name: "one element",
			ll:   NewSinglyLikedList[int](1),
			want: "1 -> nil",
		},
		{
			name: "two elements",
			ll:   NewSinglyLikedList[int](1, 2),
			want: "1 -> 2 -> nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Display(); got != tt.want {
				t.Errorf("Display() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLikedList_InsertAtBeginning(t *testing.T) {
	ll := NewSinglyLikedList[int](1)
	ll.InsertAtBeginning(2)
	ll.InsertAtBeginning(3)
	ll.InsertAtBeginning(4)

	if got := ll.Display(); got != "4 -> 3 -> 2 -> 1 -> nil" {
		t.Errorf("Elements appear in unexpected order")
	}
}

func TestSinglyLikedList_InsertAtEnd(t *testing.T) {
	ll := NewSinglyLikedList[int](4)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.InsertAtEnd(1)

	if got := ll.Display(); got != "4 -> 3 -> 2 -> 1 -> nil" {
		t.Errorf("Elements appear in unexpected order")
	}
}

func TestSinglyLikedList_Search(t *testing.T) {
	ll := NewSinglyLikedList[int](2, 4, 6, 8)

	if got := ll.Search(1); got != false {
		t.Errorf("Element 1 should not be found in the list")
	}
	if got := ll.Search(2); got != true {
		t.Errorf("Element 2 should be found in the list")
	}
}
