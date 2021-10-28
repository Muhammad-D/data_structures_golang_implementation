package implementation

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {

	var head *model.Node = model.New(0)
	Insert(3, head)
	Insert(5, head)
	Insert(45, head)
	Insert(265, head)
	Insert(14, head)

	Print(head)

}

func Insert(num int, headPointer *model.Node) {

	for headPointer.Link != nil {
		headPointer = headPointer.Link
	}
	headPointer.Link = model.New(num)

}

func Print(hp *model.Node) {

	for hp != nil {
		fmt.Println(">>>", *hp)
		hp = hp.Link
	}

}
