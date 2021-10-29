package delete_at_nth

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {

	var head *model.Node
	var place int

	insert(3, &head)
	insert(43, &head)
	insert(58, &head)
	insert(90, &head)
	print(head)
	fmt.Print("Enter a node index number you would like do delete: ")
	fmt.Scan(&place)
	delete(place, &head)
	print(head)

}

func insert(num int, hp **model.Node) {
	temp := model.New(num)
	if *hp == nil {
		*hp = temp
		return
	}

	temp.Link = *hp
	*hp = temp

}

func print(h *model.Node) {
	fmt.Print("List: ")
	for h != nil {
		fmt.Printf("%d, ", h.Num)
		h = h.Link
	}
	fmt.Println()
}

func delete(p int, hp **model.Node) {

	if p == 1 {
		*hp = (*hp).Link
		return
	}

	temp := *hp
	for i := 0; i < p-2; i++ {
		temp = temp.Link
	}
	temp1 := temp.Link
	temp.Link = temp1.Link
}
