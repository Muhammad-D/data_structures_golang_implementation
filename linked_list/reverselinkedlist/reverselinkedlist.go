package reverselinkedlist

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {
	var head *model.Node
	insert(8, &head)
	insert(34, &head)
	insert(98, &head)
	insert(359, &head)
	print(head)
	reverse(&head)
	print(head)

}

func insert(n int, h **model.Node) {
	temp := model.New(n)
	if *h == nil {
		*h = temp
		return
	}
	temp.Link = *h
	*h = temp
}

func print(h *model.Node) {
	fmt.Print("A linked list consists of following values: ")
	for h != nil {
		fmt.Printf("%d, ", h.Num)
		h = h.Link
	}
	fmt.Println()

}

func reverse(hp **model.Node) {

	var current, next, prev *model.Node

	current = *hp

	for current != nil {
		next = current.Link
		current.Link = prev
		prev = current
		current = next
	}

	*hp = prev

}
