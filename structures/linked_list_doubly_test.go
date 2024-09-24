package structures

import "testing"

func TestDoublyLikedList_Display(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		ll      *DoublyLikedList[T]
		reverse bool
		want    string
	}
	tests := []testCase[int]{
		{
			name: "empty list",
			ll:   NewDoublyLikedList[int](),
			want: "nil",
		},
		{
			name: "one element",
			ll:   NewDoublyLikedList[int](1),
			want: "1 -> nil",
		},
		{
			name: "two elements",
			ll:   NewDoublyLikedList[int](1, 2),
			want: "1 -> 2 -> nil",
		},
		{
			name: "more elements",
			ll:   NewDoublyLikedList[int](1, 2, 3, 4, 5),
			want: "1 -> 2 -> 3 -> 4 -> 5 -> nil",
		},
		{
			name:    "more elements in reverse",
			ll:      NewDoublyLikedList[int](1, 2, 3, 4, 5),
			reverse: true,
			want:    "5 -> 4 -> 3 -> 2 -> 1 -> nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Display(tt.reverse); got != tt.want {
				t.Errorf("Display() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLikedList_InsertAtBeginning(t *testing.T) {
	ll := NewDoublyLikedList[int](1)
	ll.InsertAtBeginning(2)
	ll.InsertAtBeginning(3)
	ll.InsertAtBeginning(4)

	if got := ll.Display(false); got != "4 -> 3 -> 2 -> 1 -> nil" {
		t.Errorf("Elements appear in unexpected order")
	}
	if got1, got2 := ll.Head.Value, ll.Tail.Value; got1 != 4 || got2 != 1 {
		t.Errorf("Head is expected to have a value of 4 and tail value of 1")
	}
}

func TestDoublyLikedList_InsertAtEnd(t *testing.T) {
	ll := NewDoublyLikedList[int](4)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(1)

	if got := ll.Display(false); got != "4 -> 3 -> 2 -> 1 -> nil" {
		t.Errorf("Elements appear in unexpected order: %s", got)
	}
	if got1, got2 := ll.Head.Value, ll.Tail.Value; got1 != 4 || got2 != 1 {
		t.Errorf("Head is expected to have a value of 4 and tail value of 1")
	}
}
