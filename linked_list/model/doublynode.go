package model

//A doubly linked list node STRUCT...
type DoublyNode struct {
	Num  int
	Next *DoublyNode
	Prev *DoublyNode
}

//Create a new doubly linked list node FUNC...
func NewDoublyNode(num int, next, prev *DoublyNode) *DoublyNode {
	return &DoublyNode{
		Num:  num,
		Next: next,
		Prev: prev,
	}
}
