package linkedlist

import (
	"github.com/sumanchapai/gostructs/node"
)

// Linked list is merely a pointer to a Node
type LinkedList[T any] struct {
	Head *node.Node[T]
}

func LinkedListFromSlice[T any](slice []T) LinkedList[T] {
	var initial LinkedList[T]
	var lastNode *node.Node[T]
	for index, value := range slice {
		newNode := node.Node[T]{Head: value}
		if index == 0 {
			initial.Head = &newNode
			lastNode = &newNode
		} else {
			lastNode.Tail = &newNode
			lastNode = &newNode
		}
	}
	return initial
}

func SliceFromLinkedList[T any](lst LinkedList[T]) []T {
	var slice []T
	// Check against an an empty linked list
	if !lst.IsEmpty() {
		current := lst.Head
		for i := 0; current != nil; i++ {
			slice = append(slice, current.Head)
			current = current.Tail
		}
	}
	return slice
}

// Returns if the linked list is empty
func (l LinkedList[any]) IsEmpty() bool {
	return l.Head == nil
}
