package implementation

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {

	var head *model.Node
	Insert(3, &head)
	Insert(5, &head)
	Insert(45, &head)
	Insert(265, &head)
	Insert(14, &head)

	Print(head)

}

func Insert(num int, hp **model.Node) {
	temp := model.New(num)
	temp1 := *hp
	if temp1 != nil {
		for temp1.Link != nil {
			temp1 = temp1.Link
		}
		temp1.Link = temp
		return
	}
	*hp = temp
}

func Print(hp *model.Node) {

	for hp != nil {
		fmt.Println(">>>", *hp)
		hp = hp.Link
	}

}
