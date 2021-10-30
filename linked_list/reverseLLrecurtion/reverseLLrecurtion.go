package reversellrecurtion

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
	fmt.Print("A reverse linked list: ")
	printInReverse(head)
	fmt.Println()

}

func insert(n int, h *model.Node) *model.Node {
	temp := model.New(n)
	if h == nil {
		return temp
	}
	temp.Link = h
	return temp
}

func printInReverse(ph *model.Node) {

	//base case
	if ph == nil {
		return
	}
	//recursion case
	printInReverse(ph.Link)
	fmt.Printf("%d, ", ph.Num)

}
