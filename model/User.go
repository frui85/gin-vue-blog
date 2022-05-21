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
	//处理字符串密码加密
	//data.Password = ScryptPw(data.Password)  //也可以使用gorm钩子方法
	//传入结构体，返回code状态码
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// gorm 钩子方法 BeforeSave
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	//处理字符串密码加密
	u.Password = ScryptPw(u.Password)
	return
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

// 删除用户方法
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户
// --逻辑：1. 一般不会再编辑内修改密码，密码修改需要独立方法重置密码，并且需要输入原始密码进行验证，可以放在“登录login”方法里;
// 2. 更新用户信息需要对用户角色/权限来先验证是否可以更新信息。
// 3. 需要注意 gorm 的Updates 方法支持struct, map 方式，见：https://gorm.io/zh_CN/docs/update.html#%E6%9B%B4%E6%96%B0%E5%A4%9A%E5%88%97 ，两者区别是更新有变化的字段部分，但是struct 更新非零值字段。
// 例如：User的 role 管理员定义值为 0 ,用 struct 会无法更新，所以使用map方式更新。
func EditUser(id int, data *User) int {
	var user User //gorm操作避免使用&User{},而使用&user 显的更简洁
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
