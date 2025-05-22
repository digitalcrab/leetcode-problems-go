package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	data   []int
	search map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:   make([]int, 0),
		search: make(map[int]int),
	}
}

// Insert
// Time Complexity: O(n)
func (this *RandomizedSet) Insert(val int) bool {
	if _, found := this.search[val]; found { // O(1)
		return false
	}

	// append to the data slice
	this.data = append(this.data, val) // this probably gonna be copy paste of elements O(n) ??
	// save the index (gonna be the last one)
	this.search[val] = len(this.data) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// edge case
	if len(this.data) == 0 {
		return false
	}

	idx, found := this.search[val]
	if !found {
		return false
	}

	// now it's important to swap it with the last element, so we keep track of indexes
	last := this.data[len(this.data)-1]
	// move the last element to the position of the val
	this.data[idx] = last
	// update search index
	this.search[last] = idx
	// cut the last element
	this.data = this.data[:len(this.data)-1]
	// remove from the map
	delete(this.search, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.data))
	return this.data[idx]
}

func main() {
	set := Constructor()
	fmt.Println(set.Remove(0))
	fmt.Println(set.Remove(0))
	fmt.Println(set.Insert(0))
	fmt.Println(set.GetRandom())
	fmt.Println(set.Remove(0))
	fmt.Println(set.Insert(0))
}
