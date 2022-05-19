package model

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(20);not null"`
	Password string `json:"password" gorm:"type:varchar(20);not null"`
	Role     int    `json:"role" gorm:"type:int"` //角色，0为管理员，1为普通用户
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	users := new(User)
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS //200
}

// 新增用户
func CreateUser(data *User) (code int) {
	//传入结构体，返回code状态码
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	if pageSize == -1 {
		err = db.Limit(pageSize).Find(&users).Error
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	}
	//err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	//fmt.Println(gorm.ErrRecordNotFound)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}
	return users
}
