package model

import (
	"math/rand"
	"time"
)

func CreateSlice(length int) []int {
	slice := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		slice[i] = rand.Intn(999) - rand.Intn(1000)
	}
	return slice
}
