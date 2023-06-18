package structures

import "errors"

type GraphStructsNode struct {
	id       int
	adjacent []int
	graph    *GraphStructs // for lookup purposes
}

type GraphStructs struct {
	lookup map[int]*GraphStructsNode
}

func newGraph() *GraphStructs {
	return &GraphStructs{lookup: make(map[int]*GraphStructsNode)}
}

func (g *GraphStructs) node(id int) *GraphStructsNode {
	return g.lookup[id]
}

func (g *GraphStructs) addNode(id int) {
	node := &GraphStructsNode{id: id, graph: g}
	g.lookup[id] = node
}

func (g *GraphStructs) addNodes(ids ...int) {
	for _, id := range ids {
		g.addNode(id)
	}
}

func (g *GraphStructs) addEdge(source, destination int) error {
	s := g.node(source)
	if s == nil {
		return errors.New("source node not found")
	}

	d := g.node(destination)
	if d == nil {
		return errors.New("destination node not found")
	}

	s.adjacent = append(s.adjacent, d.id)
	return nil
}

// Depth-first-search
// Time Complexity: O(V+E) where V is number of vertices in the graph and E is number of edges in the graph.
func (g *GraphStructs) hasPathDFSRecursive(source, destination int) bool {
	s := g.node(source)
	if s == nil {
		return false
	}

	d := g.node(destination)
	if d == nil {
		return false
	}

	visited := make(map[int]struct{})
	return searchDFS(g, s.id, d.id, visited)
}

func searchDFS(g *GraphStructs, source, destination int, visited map[int]struct{}) bool {
	if _, ok := visited[source]; ok {
		return false
	}
	visited[source] = struct{}{} // mark as visited

	if source == destination { // found
		return true
	}

	for _, childId := range g.node(source).adjacent {
		if searchDFS(g, childId, destination, visited) {
			return true
		}
	}

	return false
}

func (g *GraphStructs) hasPathDFSStack(source, destination int) bool {
	s := g.node(source)
	if s == nil {
		return false
	}

	d := g.node(destination)
	if d == nil {
		return false
	}

	visited := make(map[int]struct{})
	stack := []int{s.id}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		var current int
		current, stack = stack[lastIdx], stack[:lastIdx] // get current node and pop the stack

		if current == destination {
			return true
		}

		if _, ok := visited[current]; !ok { // mark as seen
			visited[current] = struct{}{}
		}

		for _, child := range g.lookup[current].adjacent {
			if _, ok := visited[child]; !ok {
				stack = append(stack, child)
			}
		}
	}

	return false
}

// Breadth-first search
// Time Complexity: O(V+E) where V is number of vertices in the graph and E is number of edges in the graph.
func (g *GraphStructs) hasPathBFS(source, destination int) bool {
	s := g.node(source)
	if s == nil {
		return false
	}

	d := g.node(destination)
	if d == nil {
		return false
	}

	queue := []int{source}
	visited := map[int]struct{}{source: {}}

	for len(queue) > 0 {
		var current int
		current, queue = queue[0], queue[1:] // dequeue

		if current == destination {
			return true
		}

		for _, child := range g.lookup[current].adjacent {
			if _, ok := visited[child]; !ok {
				queue = append(queue, child)
				visited[child] = struct{}{}
			}
		}
	}

	return false
}
