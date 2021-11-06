package arrimpl

import (
	"fmt"

	"github.com/Muhammad-D/data_structures_golang_implementation/stack_dt/model"
)

func Start() {

	myArrStack := model.NewNode()

	mySlice := model.CreateSlice(9)
	fmt.Printf("Slice:---------%v\n", mySlice)

	for _, v := range mySlice {
		myArrStack.Push(v)
	}
	myArrStack.PrintStack()
	myArrStack.Top()

	fmt.Println("----------------------------------------------------------------------")

	for i := 0; i < 5; i++ {
		myArrStack.Pop()
	}
	myArrStack.PrintStack()
	myArrStack.Top()
	fmt.Println(myArrStack.IsEmpty())
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < 4; i++ {
		myArrStack.Pop()
	}
	myArrStack.PrintStack()
	myArrStack.Top()
	fmt.Println(myArrStack.IsEmpty())

}
