package model

import (
	"math/rand"
	"time"
)

type Node struct {
	Num  int
	Link *Node
}

func New(n int) *Node {
	return &Node{
		Num: n,
	}
}
func NewWithLink(n int, l *Node) *Node {
	return &Node{
		Num:  n,
		Link: l,
	}
}

func CreateSlice(length int) []int {

	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)

	for i := range slice {
		slice[i] = rand.Intn(999) - rand.Intn(888)
	}

	return slice
}

func CreateSliceInt(length int, maxValue int) []int {

	slice := make([]int, length)

	for i := range slice {
		rand.Seed(time.Now().UnixNano())
		slice[i] = rand.Intn(maxValue)
	}
	return slice
}
