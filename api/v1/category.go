package v1

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/model"
	"github.com/frui85/gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询分类是否存在
func CateExist(c *gin.Context) {

}

// 添加分类
func AddCate(c *gin.Context) {
	data := new(model.Category)
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCate(c *gin.Context) {
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

	data := model.GetCate(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	//判断是否存在用户名，编辑后不能和已有用户名重名
	code := model.CheckCate(data.Name)
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	code := model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
