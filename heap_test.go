package tree

import (
	"fmt"
	"testing"
)

func Test_NewHeap(t *testing.T) {
	tree := NewHeap([]int{})
	tree.Put(42, 13, 54, 2, 1, 67, 43, 23)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	tree.Put(16)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	tree.Put(43)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	tree.Put(67)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	var last = -1
	for {
		if err := checkBalance(tree); err != nil {
			t.Error(err)
		}
		x, ok := tree.PopMax()
		if !ok {
			break
		}
		if last > 0 && x > last {
			t.Error("got greater than prev")
		}
		last = x
	}
}

func Test_ExistingHeap(t *testing.T) {
	tree := NewHeap([]int{42, 13, 54, 2, 1, 67, 43, 23})
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	tree.Put(16)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	tree.Put(43)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	tree.Put(67)
	if err := checkBalance(tree); err != nil {
		t.Error(err)
	}

	var last = -1
	for {
		if err := checkBalance(tree); err != nil {
			t.Error(err)
		}
		x, ok := tree.PopMax()
		if !ok {
			break
		}
		if last > 0 && x > last {
			t.Error("got greater than prev")
		}
		last = x
	}
}

func checkBalance(tree *Heap[int]) error {
	for i, x := range tree.heap {
		if l := idxOfLeftChild(i); l < len(tree.heap) {
			if tree.heap[l] > x {
				return fmt.Errorf("node #%d has %d which less than it`s child #%d eqal %d\nfull: %v", i, x, l, tree.heap[l], tree.heap)
			}
		}
		if r := idxOfRightChild(i); r < len(tree.heap) {
			if tree.heap[r] > x {
				return fmt.Errorf("node #%d has %d which less than it`s child #%d eqal %d\nfull: %v", i, x, r, tree.heap[r], tree.heap)
			}
		}
	}
	return nil
}
