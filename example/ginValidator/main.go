package main

import (
	"fmt"
	"ginValidator/model"
	"ginValidator/util"
)

func main() {

	util.Register()
	model.Register()

	input := &model.InputLogin{
		Email:    "jet@dsd.com",
		Password: "12356",
		Code:     444411,
		C1:       "张三",
	}

	input.Validate()

	fmt.Println("ok")

}
