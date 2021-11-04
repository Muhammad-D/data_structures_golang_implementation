package implementation

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start2() {

	var head *model.Node

	head = insert(23, head)
	head = insert(529, head)
	head = insert(731, head)
	head = insert(481, head)
	head = insert(1993, head)
	head = insert(815, head)

	print(head)

}

func insert(num int, hp *model.Node) *model.Node {
	temp := model.New(num)
	hpCopy := hp

	if hp == nil {
		return temp
	}

	for hpCopy.Link != nil {
		hpCopy = hpCopy.Link
	}
	hpCopy.Link = temp
	return hp
}

func print(hp *model.Node) {

	for hp != nil {
		fmt.Println(">>>", hp)
		hp = hp.Link
	}

}
