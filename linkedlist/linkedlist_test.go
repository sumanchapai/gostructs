package linkedlist

import (
	"fmt"
	"testing"
)

// TestEmptyLinkedList creates a linked list with empty slice
func TestEmptyLinkedList(t *testing.T) {
	testSlice := []int{}
	lst := LinkedListFromSlice(testSlice)
	nilLst := LinkedList[int]{}
	fmt.Println(nilLst, lst)
	if lst != nilLst {
		t.Fatalf("Not found empty linked list")
	}
}

// TestLinkedList creates a linked list with some integers
func TestLinkedList(t *testing.T) {
	testSlice := []int{2, 3, 5, 7, 11}
	lst := LinkedListFromSlice(testSlice)
	current := lst.Head
	var i int
	for i = 0; current != nil; i++ {
		expected := testSlice[i]
		found := current.Head
		if found != expected {
			t.Fatalf("Expected %v at position %v. Found %v", expected, i, found)
		}
		current = current.Tail
	}
	if i != len(testSlice) {
		t.Fatalf("Expected size of list to be %v. Found %v", len(testSlice), i)
	}
}

// TestGetSliceFromLinkedList test construction of slice from
// linked list
func TestGetSliceFromLinkedList(t *testing.T) {
	// Note that this method relies on LinkedListFromSlice
	// to build it a valid link list
	testSlice := []int{2, 3, 5, 7, 11}
	lst := LinkedListFromSlice(testSlice)
	expected := testSlice
	found := SliceFromLinkedList(lst)
	if len(expected) != len(found) {
		t.Fatalf("Expected %v at Found %v", expected, found)
	}
	for i := range expected {
		if expected[i] != found[i] {
			t.Fatalf("Expected %v at Found %v", expected, found)
		}
	}
}
