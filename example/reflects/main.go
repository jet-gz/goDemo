package main

import (
	"fmt"
	"strings"
)

type IA interface {
	test()
}

type B struct {
	IA
}

func (b *B) test() {
	fmt.Println("bbbbbb")
}

type C struct {
	IA
}

func (c *C) test() {
	fmt.Println("ccccccc")
}

func main() {
	str := "uuu ooo oo  u pp   we"

	n := strings.Index(str, " ")
	fmt.Println(str[:n])
	fmt.Println(str[n:])

}
