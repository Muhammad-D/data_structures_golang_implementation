package reverselinkedlist

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start2() {
	slice := model.CreateSlice(12)
	fmt.Printf("Slice:       %v\n", slice)
	var head *model.Node
	for _, v := range slice {
		head = insert2(v, head)
	}
	print2(head)

	head = reverse2(head, nil)

	print2(head)

}

func insert2(num int, hp *model.Node) *model.Node {
	newNode := model.New(num)
	if hp == nil {
		return newNode
	}
	temp := hp
	for temp.Link != nil {
		temp = temp.Link
	}
	temp.Link = newNode
	return hp
}

func print2(hp *model.Node) {

	fmt.Print("Linked list: [")
	for hp != nil {
		fmt.Printf("%d ", hp.Num)
		hp = hp.Link
	}
	fmt.Printf("]\n")
}

func reverse2(current, prev *model.Node) *model.Node {

	if current == nil {
		return prev
	}

	next := current.Link
	current.Link = prev
	prev = current
	current = next

	return reverse(current, prev)

}
