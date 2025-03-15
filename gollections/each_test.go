package gollections

import (
	"testing"
)

func TestEachArray(t *testing.T) {
	col := New(1, 2, 3)
	sum := 0
	col.Each(func(item interface{}) {
		sum += item.(int)
	})
	expected := 6
	if sum != expected {
		t.Errorf("Expected sum %d, got %d", expected, sum)
	}
}

func TestEachStruct(t *testing.T) {
	type test struct {
		id    int
		value int
	}
	col := New(test{id: 1, value: 10}, test{id: 2, value: 15}, test{id: 3, value: 20})
	sum := 0
	col.Each(func(item interface{}) {
		sum += item.(test).value
	})
	expected := 45
	if sum != expected {
		t.Errorf("Expected sum %d, got %d", expected, sum)
	}
}
