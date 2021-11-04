package insert_at_nth

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start2() {

	var head *model.Node

	var length, index, value int

	fmt.Print("How many anodes should be in a linked list: ")
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		fmt.Print("Enter a \"num\" field index of a current node: ")
		fmt.Scan(&index)
		fmt.Print("Enter a \"num\" field value of a current node: ")
		fmt.Scan(&value)

		head = insert2(index, value, head)
		print2(head)
	}

}

func insert2(i, v int, hp *model.Node) *model.Node {
	newNode := model.New(v)
	if hp == nil {
		return newNode
	} else if i == 1 {
		return model.NewWithLink(v, hp)
	}
	temp := hp
	for n := 0; n < i-2; n++ {
		temp = temp.Link
	}
	newNode.Link = temp.Link
	temp.Link = newNode

	return hp
}

func print2(hp *model.Node) {

	fmt.Print("The linked list consists of following values: ")
	for hp != nil {
		fmt.Printf("%d, ", hp.Num)
		hp = hp.Link
	}
	fmt.Println()

}
