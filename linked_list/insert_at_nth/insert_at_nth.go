package insert_at_nth

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {

	var head *model.Node
	insert(3, 1, &head)
	insert(5, 2, &head)
	insert(289, 1, &head)
	insert(989, 3, &head)
	insert(7678, 2, &head)
	print(head)

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
