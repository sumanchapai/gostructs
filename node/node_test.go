package node

import (
	"testing"
)

// TestTaillessNode creates a node with no tail and tests if it's head and tail
// values are correct
func TestTaillessNode(t *testing.T) {
	headValue := 5
	var tailValue *Node[int]
	n := Node[int]{Head: headValue}
	if n.Head != headValue {
		t.Fatalf("Invalid head value in node")
	}
	if n.Tail != tailValue {
		t.Fatalf("Invalid tail value in node")
	}
}

// TestNode creates a node with no tail and tests if it's head and tail
// values are correct
func TestNode(t *testing.T) {
	firstValue := 1
	secondValue := 2
	var tailValue *Node[int]
	n := Node[int]{Head: firstValue}
	n.Tail = &Node[int]{Head: secondValue}
	if n.Head != firstValue {
		t.Fatalf("Invalid first value in node")
	}
	if n.Tail.Head != secondValue {
		t.Fatalf("Invalid second value in node")
	}
	if n.Tail.Tail != tailValue {
		t.Fatalf("Invalid tail value in node")
	}
}
