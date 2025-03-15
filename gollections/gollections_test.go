package gollections

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	col := New(1, 2, 3)
	if len(col.All()) != 3 {
		t.Errorf("Expected 3 items, got %d", len(col.All()))
	}
}
