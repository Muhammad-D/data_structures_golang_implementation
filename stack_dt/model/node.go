package model

import "fmt"

type Node struct {
	arr [101]int
	top int
}

func NewNode() *Node {
	return &Node{
		top: -1,
	}
}

func (n *Node) Push(num int) {
	n.top++
	n.arr[n.top] = num
}
func (n *Node) Pop() {
	if n.top == -1 {
		fmt.Println("ERROR\nThe Stack is already empty")
		return
	}
	n.top--
}
func (n *Node) Top() int {
	fmt.Println(">>>>>>>", n.top)
	return n.top
}
func (n *Node) IsEmpty() bool {
	return n.top == -1
}
func (n *Node) PrintStack() {
	if n.top == -1 {
		fmt.Println("ERROR\nThe Stack is empty")
		return
	}

	fmt.Print("Stack:  [")
	for i := 0; i <= n.top; i++ {
		fmt.Printf("%d, ", n.arr[i])
	}
	fmt.Printf("]\nStack Top: %d\n", n.top)

}
