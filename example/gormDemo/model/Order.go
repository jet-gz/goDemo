package model

import "gorm.io/gorm"

//  一对多
// 订单
type Order struct {
	gorm.Model
	OrderNo    string      `gorm:"type:varchar(20);not null;comment:'这是一个说明'"`
	Commoditys []Commodity `gorm:"foreignKey:OrderID"`
}

// 商品
type Commodity struct {
	gorm.Model
	OrderID uint   // 这个就是外键
	Name    string `gorm:"type:varchar(20);comment:说明"`
}
