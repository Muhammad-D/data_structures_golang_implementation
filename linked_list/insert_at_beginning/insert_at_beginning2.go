package insert_at_beginning

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start2() {

	slice := model.CreateSlice(10)
	var head *model.Node

	fmt.Println(slice)

	for _, v := range slice {
		head = insert2(v, head)
	}

	print2(head)

}

func insert2(num int, hp *model.Node) *model.Node {
	return model.NewWithLink(num, hp)
}

func print2(hp *model.Node) {
	for hp != nil {
		fmt.Println("-----", hp)
		hp = hp.Link
	}
}
