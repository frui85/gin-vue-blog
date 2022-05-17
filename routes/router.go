package routes

import (
	"github.com/frui85/gin-vue-blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppModel)
	r := gin.Default()

	rg := r.Group("api/v1")
	{
		rg.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(":" + utils.HttpPort)
}
