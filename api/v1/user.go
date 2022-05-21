package v1

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/model"
	"github.com/frui85/gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {}

// 添加用户
func AddUser(c *gin.Context) {
	data := new(model.User)
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUser(c *gin.Context) {}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err)
	}

	pageNum, err := strconv.Atoi(c.Query("pagenum"))
	if err != nil {
		fmt.Println(err)
	}

	// 如果想取消limit或者Offset，就给传值-1，见gorm查询文档说明，https://gorm.io/zh_CN/docs/query.html#Limit-amp-Offset
	if pageSize <= 0 {
		pageSize = -1
	}
	//if pageNum <= 0 {
	//	pageNum = -1
	//}

	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	//判断是否存在用户名，编辑后不能和已有用户名重名
	code := model.CheckUser(data.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
