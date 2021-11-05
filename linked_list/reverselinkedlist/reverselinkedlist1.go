package reverselinkedlist

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start1() {
	var head *model.Node
	slice := model.CreateSlice(6)
	fmt.Printf("Slice:--------- %v\n", slice)
	for _, v := range slice {
		head = insert1(v, head)
	}
	print1(head)
	head = reverseRecurtion(head)
	fmt.Println("-----------------AFTER REVERSING-----------------")
	print1(head)
}

func insert1(num int, hp *model.Node) *model.Node {
	return model.NewWithLink(num, hp)
}

func print1(hp *model.Node) {

	fmt.Print("Linked list: [")
	for hp != nil {
		fmt.Printf("%d ", hp.Num)
		hp = hp.Link
	}
	fmt.Printf("]\n")
}

func reverseRecurtion(current *model.Node) *model.Node {

	if current.Link == nil {
		return current
	}
	head := reverseRecurtion(current.Link)
	current.Link.Link = current
	current.Link = nil
	return head
}
