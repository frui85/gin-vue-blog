package v1

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/model"
	"github.com/frui85/gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询分类下所有文章
func GetCateArt(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

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

	data, code := model.GetCateArt(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArtList(c *gin.Context) {
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

	data, code := model.GetArtList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加文章
func AddArt(c *gin.Context) {
	data := new(model.Article)
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArt(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑文章

func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code := model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	code := model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
