package delete_at_nth

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start2() {
	var head *model.Node

	slice := model.CreateSlice(8)
	fmt.Println(slice)

	for _, v := range slice {
		head = insert2(v, head)
	}
	print2(head)

	deleteNthSl := model.CreateSliceInt(5, 8)
	fmt.Println("deleteNthSl: ", deleteNthSl)

	for _, v := range deleteNthSl {
		head = delete2(v, head)
	}
	print2(head)

}

func insert2(num int, hp *model.Node) *model.Node {
	temp := model.New(num)
	if hp == nil {
		return temp
	}
	temp1 := hp
	for temp1.Link != nil {
		temp1 = temp1.Link
	}
	temp1.Link = temp

	return hp
}

func delete2(place int, hp *model.Node) *model.Node {
	if place == 1 {
		return hp.Link
	}

	temp := hp
	for i := 0; i < place-2; i++ {
		temp = temp.Link
	}
	temp.Link = temp.Link.Link

	return hp

}

func print2(hp *model.Node) {

	for hp != nil {
		fmt.Println("-----", hp)
		hp = hp.Link
	}

}
