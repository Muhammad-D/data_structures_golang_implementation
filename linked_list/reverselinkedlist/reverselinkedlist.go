package reverselinkedlist

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {
	var head *model.Node
	head = insert(8, head)
	head = insert(34, head)
	head = insert(98, head)
	head = insert(359, head)
	print(head)
	head = reverse(head)
	print(head)

}

func insert(n int, h *model.Node) *model.Node {
	temp := model.New(n)
	if h == nil {
		return temp
	}
	temp.Link = h
	return temp
}
func print(h *model.Node) {
	fmt.Print("A linked list consists of following values: ")
	for h != nil {
		fmt.Printf("%d, ", h.Num)
		h = h.Link
	}
	fmt.Println()

}
func reverse(current *model.Node) *model.Node {

	var next, prev *model.Node

	if current == nil {
		return nil
	}

	for current != nil {
		next = current.Link
		current.Link = prev
		prev = current
		current = next
	}

	return prev
}
