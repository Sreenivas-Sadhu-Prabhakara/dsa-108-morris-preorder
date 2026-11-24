package main

import (
	"reflect"
	"testing"
)

func TestMorrisPreorder(t *testing.T) {
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	want := []int{4, 2, 1, 3, 6, 5, 7}
	if got := MorrisPreorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("MorrisPreorder = %v, want %v", got, want)
	}
	// Threads removed, tree intact.
	if root.Left.Right.Right != nil || root.Left.Left.Right != nil {
		t.Error("temporary threads were not removed")
	}

	if got := MorrisPreorder(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("nil tree: %v", got)
	}
	chain := tn(3, tn(2, tn(1, nil, nil), nil), nil)
	if got := MorrisPreorder(chain); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("left chain: %v", got)
	}
}
