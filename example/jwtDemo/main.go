package main

import (
	"fmt"
	"jwtDemo/utils"
)

func main() {

	fy := utils.CreateUserFactory()
	ut, _ := fy.GenerateToken(101, "jet", "18000000", 100)
	fmt.Println(ut)
	claims, _ := fy.ParseToken(ut)
	fmt.Println("解析结果", claims)

	//ut = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDEsInVzZXJfbmFtZSI6ImpldCIsInBob25lIjoiMTgwMDAwMDAiLCJleHAiOjE2OTI4NDU3NjQsIm5iZiI6MTY5Mjg0NTY1NH0.0ef9aQEuYSvZ0PZg8X2BpkzQkR0hyrQZt10RhnxjbwI"

	isOk := fy.IsEffective(ut)

	if isOk {
		fmt.Println("有效！")
	} else {
		fmt.Println("过期")
	}

}
