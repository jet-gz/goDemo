package model

import (
	"fmt"
	"ginValidator/util"

	"gopkg.in/go-playground/validator.v9"
)

type IValidator interface {
	Validator()
}

type InputLogin struct {
	IValidator
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Password string `json:"password" validate:"required,max=16,min=6" label:"密码"`
	Code     int    `json:"code" validate:"gte=1000,lte=9999" label:"验证码"`
	C1       string `json:"c1" validate:"checkName" label:"测试字段"`
}

func Register() {

	util.RegisterCustom(checkName, "checkName", "{0}字段名称不能叫张三!")

}

func checkName(fl validator.FieldLevel) bool {

	val := fl.Field().String()
	//如果是张三就会触发验证
	return val != "张三"
}

func (input *InputLogin) Validate() {
	fmt.Println(input)
	err := util.GetError(input)
	// err := util.FormValidator.Struct(input)
	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		//翻译输出
	// 		m := err.Translate(util.Trans)
	// 		fmt.Println(m)
	// 	}

	// }
	if err != nil {
		fmt.Println("没有")
		return
	}

	for k, v := range err {
		fmt.Println(k, v)
	}

}
