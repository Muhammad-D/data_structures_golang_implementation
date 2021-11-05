package doublyll

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

//Start FUNC...
func Start() {

	var headPointer *model.DoublyNode
	slice := model.CreateSlice(5)
	fmt.Println(slice)

	for _, v := range slice {
		headPointer = insertAtBeg(v, headPointer)
	}
	print(headPointer)
	reversePrint(headPointer)

	fmt.Print("\n\n-----------------------------------------------------------------------------\n\n\n")

	var myHeadPointer *model.DoublyNode
	mySlice := model.CreateSlice(9)
	fmt.Println("-----------", mySlice)
	for _, v := range mySlice {
		myHeadPointer = insertAtTail(v, myHeadPointer)
	}
	print(myHeadPointer)

}

//inserts a doubly linked list node at the beginning FUNC...
func insertAtBeg(num int, hp *model.DoublyNode) *model.DoublyNode {
	if hp == nil {
		return model.NewDoublyNode(num, nil, nil)
	}
	hp.Prev = model.NewDoublyNode(num, hp, nil)
	return hp.Prev
}

func insertAtTail(num int, hp *model.DoublyNode) *model.DoublyNode {

	if hp == nil {
		return model.NewDoublyNode(num, nil, nil)
	}

	temp := hp
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = model.NewDoublyNode(num, nil, temp)
	return hp
}

//prints a doubly linked list FUNC...
func print(hp *model.DoublyNode) {

	fmt.Print("Linked list: ")

	for hp != nil {
		fmt.Printf("%d ", hp.Num)
		hp = hp.Next
	}

	fmt.Println()

}

//prints a doubly linked list in reverse FUNC...
func reversePrint(hp *model.DoublyNode) {

	for hp.Next != nil {
		hp = hp.Next
	}
	fmt.Print("A reversed linked list: [")
	for hp != nil {
		fmt.Printf("%d ", hp.Num)
		hp = hp.Prev
	}
	fmt.Printf("]\n")

}
