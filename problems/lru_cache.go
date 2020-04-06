package problems

import "container/list"

// https://leetcode.com/problems/lru-cache/
//
// Design and implement a data structure for Least Recently Used (LRU) cache.
// It should support the following operations: get and put.
//
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present.
//   When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
//
// The cache is initialized with a positive capacity.
//
// Follow up:
//  Could you do both operations in O(1) time complexity?

type entry struct {
	k int
	v int
}

type LRUCache struct {
	cap  int
	data map[int]*list.Element
	keys *list.List
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		data: make(map[int]*list.Element),
		keys: list.New(),
	}
}

func (s *LRUCache) Get(key int) int {
	if val, ok := s.data[key]; ok {
		s.keys.MoveToFront(val)
		return val.Value.(*entry).v
	}
	return -1
}

func (s *LRUCache) Put(key int, value int) {
	if item, ok := s.data[key]; ok {
		s.keys.MoveToFront(item)
		item.Value.(*entry).v = value
		return
	}

	ent := &entry{k: key, v: value}
	s.data[key] = s.keys.PushFront(ent)

	if s.keys.Len() > s.cap {
		item := s.keys.Back()
		s.keys.Remove(item)
		delete(s.data, item.Value.(*entry).k)
	}
}
