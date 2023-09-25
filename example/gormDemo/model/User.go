package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	Name  string
	Email string
}

type User struct {
	ID     int64
	Author Author `gorm:"embedded"` //`gorm:"embedded;embeddedPrefix:author_"`
	Age    int16
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("删除钩子...........")
	return
}
