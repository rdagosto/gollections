package gollections

import (
	"testing"
)

func TestAllArray(t *testing.T) {
	col := New("a", "b", "c")
	items := col.All()
	expected := []interface{}{"a", "b", "c"}

	for i, v := range items {
		if v != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}

func TestAllArrayStruct(t *testing.T) {
	type test struct {
		id   int
		name string
	}
	col := New(test{id: 1, name: "one"}, test{id: 2, name: "two"}, test{id: 3, name: "three"})
	items := col.All()
	expected := []interface{}{test{id: 1, name: "one"}, test{id: 2, name: "two"}, test{id: 3, name: "three"}}

	for i, v := range items {
		if v != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}
