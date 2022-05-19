package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignKey:ID;references:Cid"`
	//Category Category `gorm:"foreignKey:ID"`  //如果不设置影响字段，则会自动和 Article ID建立外键关联
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(20);not null"`
	Cid     int    `json:"cid" gorm:"type:bigint;not null;index"` //必须创建index索引，才能设置外键关联，否则 references:Cid" 会报 Error 1215: Cannot add foreign key constraint 错误
	Desc    string `json:"desc" gorm:"type:varchar(200)"`
	Content string `json:"content" gorm:"type:longtext"`
	Img     string `json:"img" gorm:"type:varchar(100)"`
}
