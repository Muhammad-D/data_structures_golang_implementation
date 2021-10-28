package model

type Node struct {
	Num  int
	Link *Node
}

func New(n int) *Node {
	return &Node{
		Num: n,
	}
}
