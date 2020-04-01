package algorithms

import "testing"

func TestTrie(t *testing.T) {
	root := &trieNode{children: make(map[byte]*trieNode)}
	root.add("catatonic")
	root.add("carlene")
	root.add("carefree")

	if !root.isWord("catatonic") {
		t.Error("catatonic is a word")
	}
	if !root.isWord("carlene") {
		t.Error("catatonic is a word")
	}
	if !root.isWord("carlene") {
		t.Error("catatonic is a word")
	}
	if root.isWord("cat") {
		t.Error("cat is not a word")
	}
	if root.isWord("fashionable") {
		t.Error("fashionable is not a word")
	}
	if root.isWord("carlella") {
		t.Error("catatonic is not a word")
	}
}
