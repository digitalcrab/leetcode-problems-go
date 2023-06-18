package structures

import "testing"

func printGraph(gr *GraphStructs, print func(format string, args ...interface{})) {
	for _, node := range gr.lookup {
		print("node %d, adjacent:\n", node.id)
		for _, child := range node.adjacent {
			print("\t- %d\n", child)
		}
	}
}

func initialGraph() *GraphStructs {
	gr := newGraph()
	gr.addNodes(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	gr.addEdge(1, 2)
	gr.addEdge(2, 4)
	gr.addEdge(4, 8)
	gr.addEdge(8, 9)
	gr.addEdge(9, 10)
	gr.addEdge(10, 11)
	gr.addEdge(2, 5)
	gr.addEdge(5, 6)
	gr.addEdge(5, 10)
	gr.addEdge(1, 3)
	gr.addEdge(3, 7)
	gr.addEdge(7, 11)

	return gr
}

func TestHasPathDFS(t *testing.T) {
	gr := initialGraph()

	printGraph(gr, t.Logf)

	if res := gr.hasPathDFSRecursive(1, 11); res != true {
		t.Errorf("unexpected result (1 -> 11): %t", res)
	}
	if res := gr.hasPathDFSStack(1, 11); res != true {
		t.Errorf("unexpected result (1 -> 11): %t", res)
	}
	if res := gr.hasPathBFS(1, 11); res != true {
		t.Errorf("unexpected result (1 -> 11): %t", res)
	}

	if res := gr.hasPathDFSRecursive(6, 11); res != false {
		t.Errorf("unexpected result (6 -> 11): %t", res)
	}
	if res := gr.hasPathDFSStack(6, 11); res != false {
		t.Errorf("unexpected result (6 -> 11): %t", res)
	}
	if res := gr.hasPathBFS(6, 11); res != false {
		t.Errorf("unexpected result (6 -> 11): %t", res)
	}
}
