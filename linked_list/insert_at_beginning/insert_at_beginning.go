package insert_at_beginning

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/linked_list/model"
)

func Start() {

	var head *model.Node

	var n, x int

	fmt.Print("How many nodes there should be? ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter value ")
		fmt.Scanln(&x)
		head = insert(x, head)
		print(head)
	}

}

func insert(n int, hp *model.Node) *model.Node {
	temp := model.New(n)
	temp.Link = hp
	return temp
}

func print(hp *model.Node) {
	fmt.Printf("List consists of nex values: ")
	for hp != nil {
		fmt.Printf("%d, ", hp.Num)
		hp = hp.Link
	}
	fmt.Println()
}
