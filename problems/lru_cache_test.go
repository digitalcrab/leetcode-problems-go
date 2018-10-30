package problems

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	if v := cache.Get(1); v != 1 {
		t.Errorf("Unexpected value %d", v)
	}

	cache.Put(3, 3) // evicts key 2

	if v := cache.Get(2); v != -1 {
		t.Errorf("Unexpected value %d", v)
	}

	cache.Put(4, 4) // evicts key 1

	if v := cache.Get(1); v != -1 {
		t.Errorf("Unexpected value %d", v)
	}
	if v := cache.Get(3); v != 3 {
		t.Errorf("Unexpected value %d", v)
	}
	if v := cache.Get(4); v != 4 {
		t.Errorf("Unexpected value %d", v)
	}
}

func TestLRUCache2(t *testing.T) {
	cache := NewLRUCache(2)
	if v := cache.Get(2); v != -1 {
		t.Errorf("Unexpected value get(2) = -1 / %d", v)
	}

	cache.Put(2, 6)

	if v := cache.Get(1); v != -1 {
		t.Errorf("Unexpected value get(1) = -1 / %d", v)
	}

	cache.Put(1, 5)
	cache.Put(1, 2)

	if v := cache.Get(1); v != 2 {
		t.Errorf("Unexpected value get(1) = 2 / %d", v)
	}

	if v := cache.Get(2); v != 6 {
		t.Errorf("Unexpected value get(2) = 6 / %d", v)
	}
}
