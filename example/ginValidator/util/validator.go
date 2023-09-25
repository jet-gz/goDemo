package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var FormValidator *validator.Validate
var Trans ut.Translator

func Register() {
	//中文翻译器
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	Trans, _ = uni.GetTranslator("zh")
	FormValidator = validator.New()
	//注册自定义验证函数
	//v.RegisterValidation("checkName", checkName)
	//注册翻译器
	zhTranslations.RegisterDefaultTranslations(FormValidator, Trans)
	FormValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return fmt.Sprint(field.Name, ",", label)
	})
}

// 注册自定义规则
func RegisterCustom(fn validator.Func, name string, msg string) {

	// 注册自定义验证器
	FormValidator.RegisterValidation(name, fn)
	if msg != "" {
		// 自定义错误信息
		FormValidator.RegisterTranslation(name, Trans, func(ut ut.Translator) error {
			return ut.Add(name, msg, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(name, fe.Field())
			return t
		})
	}

}

func GetError(input interface{}) map[string]string {
	var errs map[string]string
	err := FormValidator.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			m := err.Translate(Trans)
			fmt.Println("111222", m)
			n := strings.Index(m, ",")
			fmt.Println("111", m[:n], m[n:])
			errs[m[:n]] = m[n:]
		}
		return errs
	}
	return nil

}

// func (input *InputLogin) Validate() {
// 	err := util.FormValidator.Struct(input)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			//翻译输出
// 			m := err.Translate(util.Trans)
// 			fmt.Println(m)
// 		}

// 	}

// }
