package list_test

import (
	"testing"

	"go_datastructures/list"
)

func TestLinkedList_Add(t *testing.T) {
	tests := []struct {
		name       string
		itemsToAdd []string
	}{
		{
			name:       "single item",
			itemsToAdd: []string{"abc"},
		},
		{
			name:       "multiple items",
			itemsToAdd: []string{"a", "b", "c", "d", "e", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &list.LinkedList[string]{}
			for _, itm := range tt.itemsToAdd {
				l.Add(itm)
			}
			if len(tt.itemsToAdd) != l.Count() {
				t.Fatalf("expected listh to have %d items, but had %d", len(tt.itemsToAdd), l.Count())
			}
		})
	}
}

func TestLinkedList_Iterator(t *testing.T) {
	tests := []struct {
		name  string
		items []string
	}{
		{
			name:  "single item",
			items: []string{"abc"},
		},
		{
			name:  "multiple items",
			items: []string{"a", "b", "c", "d", "e", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.NewLinkedList(tt.items...)
			iter := l.Iterator()
			i := 0
			for {
				ok, itm := iter.Next()
				if !ok {
					break
				}
				if tt.items[i] != itm {
					t.Fatalf("expected element on index %d to be %s, but got %s", i, tt.items[i], itm)
				}
				i++
			}
			if i != len(tt.items) {
				t.Fatalf("expected iterator to yield %d items, but got %d", len(tt.items), i)
			}
		})
	}
}
