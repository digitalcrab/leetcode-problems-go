package structures

import "testing"

func TestHashTable_Set(t *testing.T) {
	ht := NewHashTable(10)
	ht.Set("abc", "value1")
	ht.Set("bac", "value2") // This will likely collide with `abc` because of `simpleHashFunc` implementation
	ht.Set("cab", "value3")
	ht.Set("cba", "value4")
	ht.Set("key1", "value1")

	val, found := ht.Get("abc")
	if !found || val != "value1" {
		t.Errorf("Expected value1, got %v (found = %t)", val, found)
	}

	val, found = ht.Get("bac")
	if !found || val != "value2" {
		t.Errorf("Expected value2, got %v (found = %t)", val, found)
	}

	val, found = ht.Get("cab")
	if !found || val != "value3" {
		t.Errorf("Expected value3, got %v (found = %t)", val, found)
	}

	val, found = ht.Get("cba")
	if !found || val != "value4" {
		t.Errorf("Expected value4, got %v (found = %t)", val, found)
	}

	val, found = ht.Get("key1")
	if !found || val != "value1" {
		t.Errorf("Expected value1, got %v (found = %t)", val, found)
	}
}

func TestHashTable_UpdateViaSet(t *testing.T) {
	ht := NewHashTable(10)
	ht.Set("key1", "value1")
	ht.Set("key1", "new_value1") // Update key1

	val, found := ht.Get("key1")
	if !found || val != "new_value1" {
		t.Errorf("Expected new_value1, got %v (found = %t)", val, found)
	}
}

func TestHashTable_Delete(t *testing.T) {
	ht := NewHashTable(10)
	ht.Set("key1", "value1")

	val, found := ht.Get("key1")
	if !found || val != "value1" {
		t.Errorf("Expected value1, got %v (found = %t)", val, found)
	}

	ht.Delete("key1")
	ht.Delete("key2")
	ht.Delete("key3")

	val, found = ht.Get("key1")
	if found || val != nil {
		t.Errorf("Expected nil, got %v (found = %t)", val, found)
	}
}
