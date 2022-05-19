package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(20);not null"`
	Password string `json:"password" gorm:"type:varchar(20);not null"`
	Role     int    `json:"role" gorm:"type:int"` //角色，0为管理员，1为普通用户
}
