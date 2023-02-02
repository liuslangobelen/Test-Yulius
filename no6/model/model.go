package model

import (
	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	UserId      uint   `gorm:"column:user_id;type:bigint(20)"`
	Quantity    int    `gorm:"column:quantity"`
	CodeProduct string `gorm:"column:code_product"`
	NameProduct string `gorm:"column:name_product"`
}
