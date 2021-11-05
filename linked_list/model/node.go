package model

import (
	"math/rand"
	"time"
)

//A linked list STRUCT...
type Node struct {
	Num  int
	Link *Node
}

//Create a new linked list FUNC...
func New(n int) *Node {
	return &Node{
		Num: n,
	}
}

//Create a new linked list with a "Link" argument FUNC...
func NewWithLink(n int, l *Node) *Node {
	return &Node{
		Num:  n,
		Link: l,
	}
}

//Creat a slice FUNC...
func CreateSlice(length int) []int {

	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)

	for i := range slice {
		slice[i] = rand.Intn(999) - rand.Intn(888)
	}

	return slice
}

//Create a slice with a max value FUNC...
func CreateSliceInt(length int, maxValue int) []int {

	slice := make([]int, length)

	for i := range slice {
		rand.Seed(time.Now().UnixNano())
		slice[i] = rand.Intn(maxValue)
	}
	return slice
}
