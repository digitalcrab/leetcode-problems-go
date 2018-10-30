package problems

import "container/list"

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
