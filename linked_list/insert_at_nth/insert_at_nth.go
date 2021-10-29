package insert_at_nth

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {

	var head *model.Node

	var n, num, place int

	fmt.Print("How many nodes should in a linked list: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter a value of a new node: ")
		fmt.Scan(&num)
		fmt.Print("Enter a position of the new node: ")
		fmt.Scan(&place)
		insert(num, place, &head)
		print(head)
	}
}

func insert(num, place int, hp **model.Node) {

	temp := model.New(num)
	temp1 := *hp

	if temp1 == nil {

		*hp = temp

		return
	}

	if place == 1 {
		temp.Link = temp1
		*hp = temp
		return
	}

	for i := 0; i < place-2; i++ {
		temp1 = temp1.Link
	}
	temp.Link = temp1.Link
	temp1.Link = temp

}

func print(hp *model.Node) {

	fmt.Print("A linked list consists of following numbers: ")
	for hp != nil {
		fmt.Printf(">> %d << ", hp.Num)
		hp = hp.Link
	}
	fmt.Println()
}
