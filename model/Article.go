package model

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	//Category Category `gorm:"foreignKey:ID;references:Cid"`
	//Category Category `gorm:"foreignKey:ID"` //如果不设置影响字段，则会自动和 Article ID建立外键关联
	Category Category `gorm:"foreignKey:Cid"` // Belongs to 而不是 has one
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(20);not null"`
	Cid     int    `json:"cid" gorm:"type:bigint;not null;index"` //必须创建index索引，才能设置外键关联，否则 references:Cid" 会报 Error 1215: Cannot add foreign key constraint 错误
	Desc    string `json:"desc" gorm:"type:varchar(200)"`
	Content string `json:"content" gorm:"type:longtext"`
	Img     string `json:"img" gorm:"type:varchar(100)"`
}

// 新增文章
func CreateArt(data *Article) int {
	//传入结构体，返回code状态码
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// todo 查询单个分类下的所有文章

// todo 查询单个文章

// todo 查询文章列表
func GetArtList(pageSize int, pageNum int) ([]Article, int) {
	var artList []Article
	if pageSize == -1 {
		err = db.Preload("Category").Limit(pageSize).Find(&artList).Error
	} else {
		err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Error
	}
	//err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	//fmt.Println(gorm.ErrRecordNotFound)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}
	return artList, errmsg.SUCCESS
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
