package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	//gorm.Model  //相当与父类
	ID   int16
	Name string
	Age  int16
}

// 自动以 表名
func (Student) TableName() string {
	return "user_student"
}

// 添加之前
func (u *Student) BeforeCreate(tx *gorm.DB) (err error) {

	fmt.Println("Student Before.........")

	return
}
