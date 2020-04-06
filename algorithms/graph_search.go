package algorithms

import "errors"

type graphNode struct {
	id       int
	adjacent []*graphNode
}

func (gn *graphNode) addAdjacent(n *graphNode) {
	gn.adjacent = append(gn.adjacent, n)
}

type graph struct {
	lookup map[int]*graphNode
}

func newGraph() *graph {
	return &graph{lookup: make(map[int]*graphNode)}
}

func (g *graph) node(id int) *graphNode {
	return g.lookup[id]
}

func (g *graph) addNode(id int) {
	node := &graphNode{id: id}
	g.lookup[id] = node
}

func (g *graph) addNodes(ids ...int) {
	for _, id := range ids {
		g.addNode(id)
	}
}

func (g *graph) addEdge(source, destination int) error {
	s := g.node(source)
	if s == nil {
		return errors.New("source node not found")
	}

	d := g.node(destination)
	if d == nil {
		return errors.New("destination node not found")
	}

	s.addAdjacent(d)
	return nil
}

// Depth-first-search
// Time Complexity: O(V+E) where V is number of vertices in the graph and E is number of edges in the graph.
func (g *graph) hasPathDFS(source, destination int) bool {
	s := g.node(source)
	if s == nil {
		return false
	}

	d := g.node(destination)
	if d == nil {
		return false
	}

	visited := make(map[int]struct{})
	var search func(*graphNode, *graphNode, map[int]struct{}) bool

	search = func(s, d *graphNode, v map[int]struct{}) bool {
		if _, ok := v[s.id]; ok {
			return false
		}

		v[s.id] = struct{}{}

		if s.id == d.id {
			return true
		}

		for _, child := range s.adjacent {
			if search(child, d, v) {
				return true
			}
		}

		return false
	}

	return search(s, d, visited)
}

// Breadth-first search
// Time Complexity: O(V+E) where V is number of vertices in the graph and E is number of edges in the graph.
func (g *graph) hasPathBFS(source, destination int) bool {
	s := g.node(source)
	if s == nil {
		return false
	}

	d := g.node(destination)
	if d == nil {
		return false
	}

	nextToVisit := []int{source}
	visited := make(map[int]struct{})

	for len(nextToVisit) > 0 {
		var node *graphNode
		node, nextToVisit = g.node(nextToVisit[0]), nextToVisit[1:]

		if node.id == destination {
			return true
		}
		if _, ok := visited[node.id]; ok {
			return false
		}

		visited[node.id] = struct{}{}

		for _, child := range node.adjacent {
			nextToVisit = append(nextToVisit, child.id)
		}
	}

	return false
}
