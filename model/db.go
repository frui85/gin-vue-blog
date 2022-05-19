package model

import (
	"fmt"
	"github.com/frui85/gin-vue-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	if utils.Db == "mysql" {
		dsn := utils.DbUser + ":" +
			utils.DbPassword + "@tcp(" +
			utils.DbHost + ":" +
			utils.DbPort + ")/" +
			utils.DbName +
			"?charset=utf8mb4&parseTime=True&loc=Local"
		fmt.Println(utils.Db)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   utils.DbPrefix, //设置表前缀
				SingularTable: true,           //禁用复数形式
			},
		})
	}

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数！", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	dbcp, err := db.DB()
	if err != nil {
		fmt.Println("数据库连接池方法获取失败，错误：", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbcp.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbcp.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbcp.SetConnMaxLifetime(10 * time.Second)

}
