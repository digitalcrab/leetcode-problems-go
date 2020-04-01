package algorithms

type trieNode struct {
	children map[byte]*trieNode
	complete bool
}

func (n *trieNode) add(s string) {
	n.assignToNode(s, 0)
}

func (n *trieNode) assignToNode(str string, idx int) {
	if len(str) == idx {
		n.complete = true
		return
	}

	current := str[idx]
	child := n.children[current]

	if child == nil {
		child = &trieNode{children: make(map[byte]*trieNode)}
		n.children[current] = child
	}

	child.assignToNode(str, idx+1)
}

func (n *trieNode) isWord(s string) bool {
	return n.searchInTheNode(s, 0)
}

func (n *trieNode) searchInTheNode(str string, idx int) bool {
	if len(str) == idx {
		return n.complete
	}

	child := n.children[str[idx]]
	if child == nil {
		return false
	}

	return child.searchInTheNode(str, idx+1)
}
