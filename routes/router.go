package routes

import (
	v1 "github.com/frui85/gin-vue-blog/api/v1"
	"github.com/frui85/gin-vue-blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppModel)
	r := gin.Default()

	rv1 := r.Group("api/v1")
	{
		//User-用户模块的路由接口
		rv1.POST("user/add", v1.AddUser)
		rv1.GET("users", v1.GetUsers)
		rv1.PUT("user/:id", v1.EditUser)
		rv1.DELETE("user/:id", v1.DeleteUser)

		//Category-分类模块的路由接口

		//Article-文章模块的路由接口

		//test hello
		rv1.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(":" + utils.HttpPort)
}
