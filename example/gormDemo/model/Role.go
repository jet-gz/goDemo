package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
	//不同步更新permission表
	//Permission  []Permission `json:"permission" gorm:"many2many:role_permissions;association_autoupdate:false;association_autocreate:false"`
	//同步更新permission表
	Permission []Permission `json:"permission" gorm:"many2many:role_permissions"` //多对多
}

type Permission struct {
	gorm.Model
	Title          string         `json:"title"`
	Roles          []Role         `gorm:"many2many:role_permissions"` //多对多
	UserPermission UserPermission `gorm:"references:PermissionID"`    //一对一
}

type UserPermission struct {
	gorm.Model
	PermissionID string
	C1           string
	C2           string
}
