package model

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}

// 查询分类是否存在
func CheckCate(name string) int {
	cate := new(Category)
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS //200
}

// 新增分类
func CreateCate(data *Category) int {
	//传入结构体，返回code状态码
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	if pageSize == -1 {
		err = db.Limit(pageSize).Find(&cate).Error
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	}
	//err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	//fmt.Println(gorm.ErrRecordNotFound)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}
	return cate
}

// todo 查询分类下的所有文章

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
