package gollections

import (
	"testing"
)

func TestMapArray(t *testing.T) {
	col := New(1, 2, 3)
	mapped := col.Map(func(item interface{}) interface{} {
		return item.(int) * 2
	})
	expected := []interface{}{2, 4, 6}

	if len(mapped.All()) != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), len(mapped.All()))
	}

	for i, v := range mapped.All() {
		if v != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}

func TestMapStruct(t *testing.T) {
	type test struct {
		id    int
		value int
	}

	col := New(test{id: 1, value: 10}, test{id: 2, value: 15}, test{id: 3, value: 20})

	mapped := col.Map(func(item interface{}) interface{} {
		return item.(test).value * 2
	})

	expected := []interface{}{20, 30, 40}

	if len(mapped.All()) != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), len(mapped.All()))
	}

	for i, v := range mapped.All() {
		if v.(int) != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], v)
		}
	}
}
