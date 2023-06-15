package structures

import (
	"reflect"
	"testing"
)

// TestTreeDFS focuses on Depth-First-Search
func TestTreeDFS(t *testing.T) {
	tree := NewIntBinaryTree(4, 3, 1, 12, 7)
	tree.Insert(IntVal(8))
	tree.Insert(IntVal(2))

	//			4
	//		   / \
	// 		  3  12
	//		/    /
	//     1    7
	//	 	\	 \
	//		2	  8
	expectedInOrder := []Comparable{
		IntVal(1), IntVal(2), IntVal(3),
		IntVal(4), IntVal(7), IntVal(8),
		IntVal(12),
	}

	var inOrderBuf []Comparable
	tree.Top.DFSInOrder(&inOrderBuf)

	if !reflect.DeepEqual(expectedInOrder, inOrderBuf) {
		t.Errorf("unexpected in order result %v -> %v", inOrderBuf, expectedInOrder)
	}

	//			4
	//		   / \
	// 		  3  12
	//		/    /
	//     1    7
	//	 	\	 \
	//		2	  8
	expectedPreOrder := []Comparable{
		IntVal(4), IntVal(3), IntVal(1),
		IntVal(2), IntVal(12), IntVal(7),
		IntVal(8),
	}

	var preOrderBuf []Comparable
	tree.Top.DFSPreOrder(&preOrderBuf)

	if !reflect.DeepEqual(expectedPreOrder, preOrderBuf) {
		t.Errorf("unexpected pre order result %v -> %v", preOrderBuf, expectedPreOrder)
	}

	//			4
	//		   / \
	// 		  3  12
	//		/    /
	//     1    7
	//	 	\	 \
	//		2	  8
	expectedPostOrder := []Comparable{
		IntVal(2), IntVal(1), IntVal(3),
		IntVal(8), IntVal(7), IntVal(12),
		IntVal(4),
	}

	var postOrderBuf []Comparable
	tree.Top.DFSPostOrder(&postOrderBuf)

	if !reflect.DeepEqual(expectedPostOrder, postOrderBuf) {
		t.Errorf("unexpected post order result %v -> %v", postOrderBuf, expectedPostOrder)
	}
}

// TestTreeBFS focuses on Breadth-First-Search
func TestTreeBFS(t *testing.T) {
	tree := NewIntBinaryTree(4, 3, 1, 12, 7)
	tree.Insert(IntVal(8))
	tree.Insert(IntVal(2))

	//			4
	//		   / \
	// 		  3  12
	//		/    /
	//     1    7
	//	 	\	 \
	//		2	  8
	expected := []Comparable{
		IntVal(4), IntVal(3), IntVal(12),
		IntVal(1), IntVal(7), IntVal(2),
		IntVal(8),
	}

	var buf []Comparable
	tree.Top.BFS(&buf)

	if !reflect.DeepEqual(expected, buf) {
		t.Errorf("unexpected bfs result %v -> %v", buf, expected)
	}
}
