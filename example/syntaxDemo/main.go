package main

import (
	"math/rand"
)

func getRandom() *int {
	tmp := rand.Intn(100)
	return &tmp
}
func main() {

	n := getRandom()
	println(*n)
}
