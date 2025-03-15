package gollections

import (
	"testing"
)

func TestFilterArray(t *testing.T) {
	col := New(1, 2, 3, 4, 5)
	filtered := col.Filter(func(item interface{}) bool {
		return item.(int) > 2
	})
	expected := []interface{}{3, 4, 5}

	if len(filtered.All()) != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), len(filtered.All()))
	}

	for i, v := range filtered.All() {
		if v != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}

func TestFilterStruct(t *testing.T) {
	type test struct {
		id    int
		value int
	}
	col := New(test{id: 1, value: 10}, test{id: 2, value: 15}, test{id: 3, value: 20})

	filtered := col.Filter(func(item interface{}) bool {
		return item.(test).value >= 15
	})

	if len(filtered.All()) != 2 {
		t.Errorf("Expected 2 items, got %d", len(filtered.All()))
	}
}
