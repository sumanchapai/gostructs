package node

import (
	"testing"
)

// TestTaillessNode creates a node with no tail and tests if it's head and tail
// values are correct
func TestTaillessNode(t *testing.T) {
	headValue := 5
	var tailValue *Node[int]
	n := Node[int]{Value: headValue}
	if n.Value != headValue {
		t.Fatalf("Invalid head value in node")
	}
	if n.Next != tailValue {
		t.Fatalf("Invalid tail value in node")
	}
}

// TestNode creates a node with no tail and tests if it's head and tail
// values are correct
func TestNode(t *testing.T) {
	firstValue := 1
	secondValue := 2
	var tailValue *Node[int]
	n := Node[int]{Value: firstValue}
	n.Next = &Node[int]{Value: secondValue}
	if n.Value != firstValue {
		t.Fatalf("Invalid first value in node")
	}
	if n.Next.Value != secondValue {
		t.Fatalf("Invalid second value in node")
	}
	if n.Next.Next != tailValue {
		t.Fatalf("Invalid tail value in node")
	}
}
