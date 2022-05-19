package main

import (
	"github.com/frui85/gin-vue-blog/model"
	"github.com/frui85/gin-vue-blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
