package structures

// HashTable represents a hast map, with buckets and hash function
// Disclaimer: This is not anyhow close to production ready code and meant only for education.
type HashTable struct {
	buckets []*hashNode
	size    int
	hashFn  func(string) int
}

// hashNode represent a single node in the hash table bucket
// next is used to chain nodes in case of hash coli
type hashNode struct {
	key   string
	value any
	next  *hashNode
}

// simpleHashFunc is super artificial hash func based on int representation of char ;)
func simpleHashFunc(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash
}

// NewHashTable creates a new hash map with a given size - number of buckets to split the data
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*hashNode, size),
		size:    size,
		hashFn:  simpleHashFunc,
	}
}

func (ht *HashTable) bucketIdx(key string) int {
	// calculate a hash first
	hash := ht.hashFn(key)
	// find the bucket index
	idx := hash % ht.size
	return idx
}

// Set adds value to the hash map
// Time Complexity:
// O(1) average-case if there are few collisions (probably not with `simpleHashFunc` implementation)
// O(n)worst-case scenario where `n` is number of all elements in the hash map that fall into the same bucket
func (ht *HashTable) Set(key string, value any) {
	idx := ht.bucketIdx(key)
	// create a new hash map node
	node := &hashNode{key: key, value: value}

	// bucket is empty we add a first node
	if ht.buckets[idx] == nil {
		ht.buckets[idx] = node
		return
	}

	// start with the first element
	current := ht.buckets[idx]

	// keep the loop running till we run out of elements
	for current != nil {
		// if the item has the same key, we update it and return
		if current.key == key {
			current.value = value
			return
		}

		// we found the last element, so we give his next pointer to a node
		if current.next == nil {
			current.next = node
			return
		}

		// otherwise just move the pointer
		current = current.next
	}
}

// Get returns a value from a hash map by the key
// Time Complexity: similar to Set
func (ht *HashTable) Get(key string) (any, bool) {
	idx := ht.bucketIdx(key)
	// bucket if not empty
	if ht.buckets[idx] != nil {
		// loop over all items
		for current := ht.buckets[idx]; current != nil; current = current.next {
			// exactly this key
			if current.key == key {
				return current.value, true
			}
		}
	}
	// in case of nothing found, lets return just nil
	return nil, false
}

// Delete removes a value from the hash map by the key
// Time Complexity: similar to Set
func (ht *HashTable) Delete(key string) {
	idx := ht.bucketIdx(key)
	// bucket if empty
	if ht.buckets[idx] == nil {
		return
	}

	current := ht.buckets[idx]

	// remove first element if necessary
	if current.key == key {
		ht.buckets[idx] = current.next
		return
	}

	// keep track of prev and current elements
	prev := current
	current = current.next

	// loop over all items
	for current != nil {
		// find matching element
		if current.key == key {
			// we must switch links for prev.next element point to the current.next
			// essentially jump over the current
			prev.next = current.next
			// nothing to do more
			return
		}

		// move pointers of a current to a new one
		prev, current = current, current.next
	}
}
