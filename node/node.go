package node

type Node[T any] struct {
	Head T
	Tail *Node[T]
}
